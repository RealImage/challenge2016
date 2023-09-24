package estimation

type Distributor struct {
	ParentDistributor      string                  `json:"parentDistributor,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	Include                []IncludeExcludeDetails `json:"include,omitempty"`
	Exclude                []IncludeExcludeDetails `json:"exclude,omitempty"`
	CheckDistributionInput []IncludeExcludeDetails `json:"checkDistributionInput,omitempty"`
	AddPermissions         []IncludeExcludeDetails `json:"addPermissions,omitempty"`
}

type IncludeExcludeDetails struct {
	City         string `json:"city,omitempty"`
	ProvinceName string `json:"provinceName,omitempty"`
	CountryName  string `json:"countryName,omitempty"`
}

func CheckDistributionInclude(input Distributor) bool {
	// csvData := SetDetails()
	for _, i := range input.Include {
		for _, j := range input.CheckDistributionInput {
			if checkRegions(i.City, j.City) {
				return true
			} else if checkRegions(i.CountryName, j.CountryName) {
				return true
			} else if checkRegions(i.ProvinceName, j.ProvinceName) {
				return true
			}
		}

	}
	return false
}
func CheckDistributionExclude(input Distributor) bool {
	// csvData := SetDetails()
	for _, i := range input.Exclude {
		for _, j := range input.CheckDistributionInput {
			if checkRegions(i.City, j.City) {
				return true
			} else if checkRegions(i.ProvinceName, j.ProvinceName) {
				return true
			} else if checkRegions(i.CountryName, j.CountryName) {
				return true
			}
		}

	}
	return false
}

func AddDistributor(input Distributor, PermissionInput Distributor, dist string) (result Distributor) {
	if input.Name == dist {

		for _, j := range input.AddPermissions {

			k := IncludeExcludeDetails{}
			k.City = j.City
			k.CountryName = j.CountryName
			k.ProvinceName = j.ProvinceName
			PermissionInput.Include = append(PermissionInput.Include, k)
		}
		PermissionInput.Exclude = input.Exclude
		PermissionInput.CheckDistributionInput = input.CheckDistributionInput

	}
	return PermissionInput
}

func CheckDistribution(input Distributor) (string, string) {
	var OriginalInput Distributor
	var PermissionInputDist2 Distributor
	var PermissionInputDist3 Distributor
	var permissionResultForDistributor2 Distributor
	var permissionResultForDistributor3 Distributor
	if input.ParentDistributor == "Distributor1" || input.Name == "Distributor1" {
		OriginalInput = input
	}
	if input.ParentDistributor == "Distributor2" || input.Name == "Distributor2" {
		permissionResultForDistributor2 = AddDistributor(input, PermissionInputDist2, input.Name)
		permissionResultForDistributor2.Name = input.Name
		permissionResultForDistributor2.AddPermissions = input.AddPermissions
	}
	if input.Name == "Distributor3" {
		permissionResultForDistributor3 = AddDistributor(permissionResultForDistributor2, PermissionInputDist3, "Distributor3")
	}

	var include bool
	var exclude bool
	if input.Name == "Distributor1" {
		include = CheckDistributionInclude(OriginalInput)
		exclude = CheckDistributionExclude(OriginalInput)
	} else if input.Name == "Distributor2" {
		include = CheckDistributionInclude(permissionResultForDistributor2)
		exclude = CheckDistributionExclude(permissionResultForDistributor2)
	} else if input.Name == "Distributor3" {
		include = CheckDistributionInclude(permissionResultForDistributor3)
		exclude = CheckDistributionExclude(permissionResultForDistributor3)
	}

	if exclude {
		return "", "false"
	} else if include {
		return "true", ""
	}
	return "", "false"
}

func checkRegions(city string, state string) bool {
	if city == "" {
		return false
	} else if state == "" {
		return false
	}
	if city == state {
		return true
	} else {
		return false
	}
}
