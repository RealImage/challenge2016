package main

import (

	"github.com/gin-gonic/gin"
)

// Create a new distributor
func CreateDistributor(regions []Region) gin.HandlerFunc {
	return func(c *gin.Context) {
		var distributor Distributor
		if err := c.ShouldBindJSON(&distributor); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if _, exists := distributors[distributor.Name]; exists {
			c.JSON(400, gin.H{"error": "Distributor with the same name already exists"})
			return
		}

		// Check hierarchical permissions
		if distributor.Parent != "" {
			parentDistributor, parentExists := distributors[distributor.Parent]
			if !parentExists {
				c.JSON(400, gin.H{"error": "Parent distributor does not exist"})
				return
			}
			if !IsSubsetOf(distributor.Permissions, parentDistributor.Permissions) {
				c.JSON(400, gin.H{"error": "Invalid permissions for the child distributor"})
				return
			}
		}

		// Validate distributor permissions with region data
		for _, region := range regions {
			for _, included := range distributor.Permissions.Include {
				if region.CityName == included {
					for _, excluded := range distributor.Permissions.Exclude {
						if region.CityName == excluded {
							c.JSON(400, gin.H{"error": "Invalid permissions for region: " + region.CityName})
							return
						}
					}
				}
			}
		}

		// Store the distributor
		distributors[distributor.Name] = distributor
		c.JSON(201, gin.H{"message": "Distributor created successfully"})
	}
}

// Retrieve all distributors and their regions
func GetAllDistributors(c *gin.Context) {
	distributorData := make(map[string]Distributor)
	for distributorName, distributor := range distributors {
		distributorData[distributorName] = distributor
	}
	c.JSON(200, distributorData)
}


func CheckDistributorPermission(c *gin.Context) {
	distributorName := c.Param("distributor")
	region := c.Param("region")
	if CheckPermission(distributorName, region) {
		c.JSON(200, gin.H{"message": "Yes"})
	} else {
		c.JSON(403, gin.H{"message": "No"})
	}
}