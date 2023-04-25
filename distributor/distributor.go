package distributor

import "strings"

type Distributor struct {
	included             map[string]bool
	excluded             map[string]bool
	subDistributors      []*Distributor
	subDistributorsCount int
}

// NewDistributor Constructor for Distributor class
func NewDistributor(included, excluded []string, subDistributors []*Distributor,
	subDistributorsCount int) *Distributor {
	distributor := &Distributor{
		included:             make(map[string]bool),
		excluded:             make(map[string]bool),
		subDistributors:      subDistributors,
		subDistributorsCount: subDistributorsCount,
	}

	// Add included regions
	for _, region := range included {
		distributor.included[region] = true
	}

	// Add excluded regions
	for _, region := range excluded {
		distributor.excluded[region] = true
	}

	return distributor
}

func (d *Distributor) GetIncludedRegion() map[string]bool {
	return d.included
}

func (d *Distributor) GetExcludedRegion() map[string]bool {
	return d.excluded
}

// AddIncludedRegion Define methods to add or remove regions from the permissions
func (d *Distributor) AddIncludedRegion(region string) {
	d.included[region] = true
}

func (d *Distributor) RemoveIncludedRegion(region string) {
	delete(d.included, region)
}

func (d *Distributor) AddExcludedRegion(region string) {
	d.excluded[region] = true
}

func (d *Distributor) RemoveExcludedRegion(region string) {
	delete(d.excluded, region)
}

func (d *Distributor) AddSubDistributor(included, excluded []string, subDistributors []*Distributor,
	subDistributorsCount int) *Distributor {

	newSubDistributor := NewDistributor(included, excluded, subDistributors, subDistributorsCount)
	d.subDistributors = append(d.subDistributors, newSubDistributor)
	return newSubDistributor
}

func (d *Distributor) GetSubDistributorsCount() int {
	return d.subDistributorsCount
}

// HasPermission Define a method to check if a given region is included or excluded
func (d *Distributor) HasPermission(region string) bool {
	if d.excluded[region] {
		return false
	}
	if d.included[region] {
		// check if any of the distributor's sub-distributors have excluded the region
		for subRegion, _ := range d.excluded {
			if strings.HasPrefix(subRegion, region) {
				return false
			}
		}

		return true
	}
	return false
}
