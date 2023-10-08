package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// to fetch all available distributors
func (l *LocationData) GetDistributors(c *gin.Context) {
	var distributors interface{}
	if len(l.Distributors) > 0 {
		distributors = l.Distributors
	} else {
		distributors = struct{}{}
	}

	l.SetAPIResponse(c, "", distributors)
}

// to check distributor permission
func (l *LocationData) DistributorPermissionCheck(c *gin.Context) {
	if c.Param("distributor") == "" {
		l.SetAPIResponse(c, "Invalid distributor", "")
		return
	}

	Distributor := c.Param("distributor")
	if _, ok := l.Distributors[Distributor]; !ok {
		l.SetAPIResponse(c, "Invalid distributor", "")
		return
	}

	if c.Param("permission") == "" {
		l.SetAPIResponse(c, "Invalid permission", "")
		return
	}

	PermissionParts := strings.Split(c.Param("permission"), "-")
	if len(PermissionParts) > 3 {
		l.SetAPIResponse(c, "Invalid permission", "")
		return
	}

	var Country, Province, City string
	switch len(PermissionParts) {
	case 3:
		Country, Province, City = PermissionParts[2], PermissionParts[1], PermissionParts[0]
		if Country == "" || Province == "" || City == "" {
			l.SetAPIResponse(c, "Invalid permission", "")
			return
		}
	case 2:
		Country, Province, City = PermissionParts[1], PermissionParts[0], ""
		if Country == "" || Province == "" {
			l.SetAPIResponse(c, "Invalid permission", "")
			return
		}
	case 1:
		Country, Province, City = PermissionParts[0], "", ""
		if Country == "" {
			l.SetAPIResponse(c, "Invalid permission", "")
			return
		}
	}

	ValidLoc := l.ValidateLocation(Country, Province, City)
	if !ValidLoc {
		l.SetAPIResponse(c, "Invalid Region.", "")
		return
	}

	HavePermission := l.PermissionCheckForDistributor(Distributor, Country, Province, City)
	RespString := "NO"
	if HavePermission {
		RespString = "YES"
	}
	l.SetAPIResponse(c, "", RespString)
	return
}

// create/update distributor
func (l *LocationData) CreateDistributor(c *gin.Context) {
	var requestBody DistribRequestBody

	// retreiving request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		l.SetAPIResponse(c, "INVALID DATA", "")
		return
	}

	// Verifying Distributor and Included regions are provided
	if !(requestBody.Distributor != "" && len(requestBody.Include) > 0) {
		l.SetAPIResponse(c, "Invalid Distributor data", "")
		return
	}

	// Checking distributor exists or not
	if _, ok := l.Distributors[requestBody.Distributor]; ok {
		l.SetAPIResponse(c, "Distributor already exists", "")
		return
	}

	// Checking for parent distributor
	if requestBody.ParentDistributor != "" {
		if _, ok := l.Distributors[requestBody.ParentDistributor]; !ok {
			l.SetAPIResponse(c, "Invalid Parent Distributor", "")
			return
		}
		if requestBody.ParentDistributor == requestBody.Distributor {
			l.SetAPIResponse(c, "Parent Distributor and Distributor names are same", "")
			return
		}
	}

	IncludeStruct := make(map[string]struct{})
	if len(requestBody.Include) > 0 {
		// checking the provided regions. If there is any invalid region process failed
		for _, Inc := range requestBody.Include {
			IncParts := strings.Split(Inc, "-")

			if len(IncParts) > 3 {
				l.SetAPIResponse(c, "Invalid Region", "")
				return
			}

			var Country, Province, City string

			switch len(IncParts) {
			case 3:
				Country, Province, City = IncParts[2], IncParts[1], IncParts[0]
				if Country == "" || Province == "" || City == "" {
					l.SetAPIResponse(c, "Invalid permission", "")
					return
				}
			case 2:
				Country, Province, City = IncParts[1], IncParts[0], ""
				if Country == "" || Province == "" {
					l.SetAPIResponse(c, "Invalid permission", "")
					return
				}
			case 1:
				Country, Province, City = IncParts[0], "", ""
				if Country == "" {
					l.SetAPIResponse(c, "Invalid permission", "")
					return
				}
			}

			// validate locations
			ValidLoc := l.ValidateLocation(Country, Province, City)
			if !ValidLoc {
				l.SetAPIResponse(c, "Invalid Region.", "")
				return
			}

			if requestBody.ParentDistributor != "" {
				// checking for parent's authorization in regions
				valid := l.ParentDistributorValidate(requestBody.ParentDistributor, Country, Province, City)
				if !valid {
					l.SetAPIResponse(c, "Invalid Region. Parent not authorized for the region.", "")
					return
				}
			}

			// for checking the existence of country in IncludeStruct
			validEntry := l.ExistenceCheck(IncludeStruct, Inc, City, Province, Country)
			if !validEntry {
				l.SetAPIResponse(c, "Repeated use of region or Duplicate Region entries in included list.", "")
				return
			}

			IncludeStruct[Inc] = struct{}{}
		}
	}

	ExcludeStruct := make(map[string]struct{})
	if len(requestBody.Exclude) > 0 {
		for _, Exc := range requestBody.Exclude {
			ExcParts := strings.Split(Exc, "-")

			if len(ExcParts) > 3 {
				l.SetAPIResponse(c, "Invalid Region", "")
				return
			}

			var Country, Province, City string

			switch len(ExcParts) {
			case 3:
				Country, Province, City = ExcParts[2], ExcParts[1], ExcParts[0]
			case 2:
				Country, Province, City = ExcParts[1], ExcParts[0], ""
			case 1:
				Country, Province, City = ExcParts[0], "", ""
			}

			ValidLoc := l.ValidateLocation(Country, Province, City)
			if !ValidLoc {
				l.SetAPIResponse(c, "Invalid Region.", "")
				return
			}

			ValidExc := l.ValidateExcludedRegions(IncludeStruct, Country, Province, City)
			if !ValidExc {
				l.SetAPIResponse(c, "Invalid Regions in exclude.", "")
				return
			}

			// for checking the existence of country in ExcludeStruct
			validEntry := l.ExistenceCheck(ExcludeStruct, Exc, City, Province, Country)
			if !validEntry {
				l.SetAPIResponse(c, "Repeated use of region or Duplicate Region entries in excluded list.", "")
				return
			}

			ExcludeStruct[Exc] = struct{}{}
		}
	}

	l.Distributors[requestBody.Distributor] = DistributorData{
		ParentDistributor: requestBody.ParentDistributor,
		Included:          IncludeStruct,
		Excluded:          ExcludeStruct,
	}

	if requestBody.ParentDistributor != "" {
		l.DistributorParent[requestBody.Distributor] = requestBody.ParentDistributor
	}

	l.SetAPIResponse(c, "", "")
	return
}
