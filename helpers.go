package main


// Check if a distributor has permission for a region
func CheckPermission(distributorName string, region string) bool {
	for distributorName != "" {
		distributor, exists := distributors[distributorName]
		if !exists {
			return false
		}
		permissions := distributor.Permissions

		for _, excluded := range permissions.Exclude {
			if region == excluded {
				return false
			}
		}

		for _, included := range permissions.Include {
			if region == included {
				return true
			}
		}

		distributorName = distributor.Parent
	}
	return false
}

// Check if the given permissions are a subset of another set of permissions
func IsSubsetOf(subset Permission, superset Permission) bool {
	for _, excluded := range subset.Exclude {
		if !contains(superset.Exclude, excluded) {
			return false
		}
	}
	for _, included := range subset.Include {
		if !contains(superset.Include, included) {
			return false
		}
	}
	return true
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}