package distributor

import (
	"RealImage/models"
	"RealImage/utils"
	"fmt"
)

// Distributor represents a distributor and its permissions.
type Distributor struct {
	Name       string             `json:"Name"`
	Permission *models.Permission `json:"Permission"`
	ParentName string             `json:"ParentName,omitempty"`
	Parent     *Distributor       `json:"Parent,omitempty"`
}

// Data structure to store all the distributors Map[distributorName]->Distributor
var distributors map[string]*Distributor

func init() {
	distributors = make(map[string]*Distributor)
}

func (d *Distributor) SetName(name string) {
	d.Name = name
}

func (d *Distributor) SetParent(parentDistributor *Distributor) {
	d.Parent = parentDistributor
}

func (d *Distributor) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// Add distributor to map and establish parent-child relationship
func CreateDistributors(distributorList []Distributor) {
	for i := 0; i < len(distributorList); i++ {
		if distributorList[i].ParentName != "" {
			distributorList[i].Parent = GetDistributor(distributorList[i].ParentName)
		}
		distributors[distributorList[i].Name] = &distributorList[i]
	}
}

// Get distributor by name
func GetDistributor(name string) *Distributor {
	return distributors[name]
}

// Set include permissions for distributor
func (d *Distributor) SetIncludePermissions(codes []models.Location) {
	if d.Permission == nil {
		d.Permission = &models.Permission{}
	}
	d.Permission.Include = codes
	fmt.Println("Include permissions set for ", d.Name)
}

// Set exclude permissions for distributor
func (d *Distributor) SetExcludePermissions(codes []models.Location) {
	if d.Permission == nil {
		d.Permission = &models.Permission{}
	}
	d.Permission.Exclude = codes
}

// Check if the distributor is not authorized to distribute in a location
func (d *Distributor) checkExcludePermissions(location models.Location) bool {
	emptyLocation := models.Location{}
	// Check if the region is explicitly excluded.
	for _, exclude := range d.Permission.Exclude {
		enclosingRegion := utils.GetEnclosingRegion(location)
		for {
			if exclude == location || exclude == enclosingRegion {
				return true
			} else if enclosingRegion == emptyLocation {
				break
			} else {
				enclosingRegion = utils.GetEnclosingRegion(enclosingRegion)
			}
		}
	}

	if d.Parent != nil {
		// Check if there is any entry for the enclosing region in the permissions of the distributor.
		if d.Parent.checkExcludePermissions(location) {
			return true
		}

	}

	return false
}

// Check if the distributor is authorized to distribute in a location
func (d *Distributor) checkIncludePermissions(location models.Location) bool {
	emptyLocation := models.Location{}
	for _, include := range d.Permission.Include {
		enclosingRegion := utils.GetEnclosingRegion(location)
		for {
			if include == location || include == enclosingRegion {
				return true
			} else if enclosingRegion == emptyLocation {
				break
			} else {
				enclosingRegion = utils.GetEnclosingRegion(enclosingRegion)
			}
		}
	}
	return false
}

// Check include and exclude permissions combined
func (d *Distributor) CheckPermissions(location models.Location) bool {

	isExcluded := d.checkExcludePermissions(location)

	isIncluded := d.checkIncludePermissions(location)

	if isExcluded {
		return false
	}
	if isIncluded {
		return true
	}
	return false
}
