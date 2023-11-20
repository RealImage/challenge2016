package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	repo "github.com/nikhilsiwach28/Cinema-Distribution-System/inMemory"
	"github.com/nikhilsiwach28/Cinema-Distribution-System/models"
)

// DISTRIBUTOR SERVICE
type DistributorService interface {
	CheckAuthorization(distributorID uuid.UUID, region string) bool
	RegisterNewDistributor(distributor models.CreateDistributorRequest) (models.Distributor, error)
}

type distributorService struct {
	repo repo.InMemory
}

func NewDistributorService(inMemory repo.InMemory) *distributorService {
	return &distributorService{repo: inMemory}
}

func (d *distributorService) CheckAuthorization(id uuid.UUID, region string) bool {
	parts := strings.Split(region, "-")

	switch len(parts) {
	case 1: // Only country present
		return d.repo.CheckIfCountryAuth(id, parts[0])
	case 2: // State and country present
		return d.repo.CheckIfStateAuth(id, parts[0], parts[1])
	case 3: // City, state, and country present
		return d.repo.CheckIfCityAuth(id, parts[0], parts[1], parts[2])
	default: // Invalid format
		return false
	}
}

func (d distributorService) RegisterNewDistributor(disReq models.CreateDistributorRequest) (models.Distributor, error) {
	dist := models.CreateRequestToDistributor(disReq)
	if ok := d.repo.Add(*dist); ok != true {
		return *dist, errors.New("failed to add distributor")
	} else {
		return *dist, nil
	}

}

// SPLIT DISTRIBUTOR SERVICE

type SplitDistributorService interface {
	SplitDistribution(models.CreateSplitDistributorRequest) (models.Distributor, error)
}

type splitDistributorService struct {
	repo repo.InMemory
}

func NewSplitDistributorService(inMemory repo.InMemory) *splitDistributorService {
	return &splitDistributorService{repo: inMemory}
}

func (sd splitDistributorService) SplitDistribution(splitDistReq models.CreateSplitDistributorRequest) (models.Distributor, error) {
	dist := models.SplitRequestToDistributor(splitDistReq)
	sd.repo.CheckIfIncludesAuthByParent(dist, splitDistReq.ParentId)
	sd.repo.CheckIfExcludesAuthByParent(dist, splitDistReq.ParentId)
	if ok := sd.repo.Add(*dist); ok != true {
		return *dist, errors.New("failed to add distributor")
	} else {
		return *dist, nil
	}
}
