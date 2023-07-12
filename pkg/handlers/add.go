package handlers

import (
	"context"
	"net/http"

	"chng2016/pkg/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// DistributorAddRequest ...
type DistributorAddRequest struct {
	DistributorID string                         `json:"distributorID" binding:"required"`
	Permission    *models.DistributorPermissions `json:"permission" binding:"required"`
}

// SubDistributorAddRequest ...
type SubDistributorAddRequest struct {
	DistributorID    string
	SubDistributorID string                         `json:"subDistributorID" binding:"required"`
	Permission       *models.DistributorPermissions `json:"permission" binding:"required"`
}

// AddDistributor ...
func (d *DistributorHandler) AddDistributor(c *gin.Context) {
	req := DistributorAddRequest{}

	// validating input arguments
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": d.validator.CustomValidationError(&req, err),
		})
		return
	}

	// validate permissions codes
	if err := d.validatePermissionCodes(req.Permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	distributor := d.dataStore.GetCache(req.DistributorID)
	if distributor == nil {
		distributor = &models.Distributor{
			ID:          req.DistributorID,
			Permissions: req.Permission,
			TrieRoot:    &models.TrieNode{Children: make(map[string]*models.TrieNode)},
		}
		d.dataStore.SetCache(req.DistributorID, distributor)
	} else {
		distributor.Permissions = req.Permission
		distributor.TrieRoot = &models.TrieNode{Children: make(map[string]*models.TrieNode)}
	}

	if err := d.addToTrie(distributor.TrieRoot, req.Permission.Include, req.Permission.Exclude); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"distributorID": req.DistributorID,
	})
}

func (d *DistributorHandler) validatePermissionCodes(permission *models.DistributorPermissions) error {
	eg, _ := errgroup.WithContext(context.Background())
	for _, ir := range permission.Include {
		includeRegion := ir
		eg.Go(func() error {
			_, err := d.util.GetRegionDetails(includeRegion)
			if err != nil {
				return err
			}
			return nil
		})
	}

	for _, er := range permission.Exclude {
		excludeRegion := er
		eg.Go(func() error {
			_, err := d.util.GetRegionDetails(excludeRegion)
			if err != nil {
				return err
			}
			return nil
		})
	}

	return eg.Wait()
}

func (d *DistributorHandler) addToTrie(root *models.TrieNode, include []string, exclude []string) error {
	eg, _ := errgroup.WithContext(context.Background())

	for _, region := range include {
		r := region
		eg.Go(func() error {
			return d.dataStore.AddRegionToTrie(root, r, true)
		})
	}

	for _, region := range exclude {
		r := region
		eg.Go(func() error {
			return d.dataStore.AddRegionToTrie(root, r, false)
		})
	}

	return eg.Wait()
}

func (d *DistributorHandler) AddSubDistributor(c *gin.Context) {
	req := SubDistributorAddRequest{}
	req.DistributorID = c.Param("distributorID")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": d.validator.CustomValidationError(&req, err),
		})
		return
	}

	// validate permissions codes
	if err := d.validatePermissionCodes(req.Permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminDistributor := d.dataStore.GetCache(req.DistributorID)
	if adminDistributor == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "admin distributor details not found",
		})
		return
	}

	distributor := d.dataStore.GetCache(req.SubDistributorID)
	if distributor == nil {
		distributor = &models.Distributor{
			ID:          req.DistributorID,
			Permissions: req.Permission,
			TrieRoot:    &models.TrieNode{Children: make(map[string]*models.TrieNode)},
		}
		d.dataStore.SetCache(req.SubDistributorID, distributor)
	} else {
		distributor.Permissions = req.Permission
		distributor.TrieRoot = &models.TrieNode{Children: make(map[string]*models.TrieNode)}
	}

	// check parent permission for the following action

	eg, _ := errgroup.WithContext(context.Background())

	for _, incReg := range distributor.Permissions.Include {
		includeRegion := incReg
		eg.Go(func() error {
			region, err := d.util.GetRegionDetails(includeRegion)
			if err != nil {
				return err
			}

			if !d.dataStore.HasPermission(adminDistributor.TrieRoot, adminDistributor.Permissions.Include, adminDistributor.Permissions.Exclude, region.CountryCode, region.StateCode, region.CityCode) {
				return ErrPermissionDenied
			}

			return nil
		})
	}

	for _, excReg := range distributor.Permissions.Exclude {
		excludeRegion := excReg
		eg.Go(func() error {
			region, err := d.util.GetRegionDetails(excludeRegion)
			if err != nil {
				return err
			}
			if !d.dataStore.HasPermission(adminDistributor.TrieRoot, adminDistributor.Permissions.Include, adminDistributor.Permissions.Exclude, region.CountryCode, region.StateCode, region.CityCode) {
				return ErrPermissionDenied
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Permission.Exclude = append(req.Permission.Exclude, adminDistributor.Permissions.Exclude...)
	if err := d.addToTrie(distributor.TrieRoot, req.Permission.Include, req.Permission.Exclude); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"subDistributorID": req.SubDistributorID,
	})
}
