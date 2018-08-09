package distributor

// Checking distributor is already available or not
func checkDistributorExists(distributor Distributor) (isAvailable bool, errMsg string) {
	_, isDistributorAvailable := distributors[distributor.Name]
	if !isDistributorAvailable {
		isAvailable = false
	} else {
		isAvailable = true
		errMsg = "Unable to create distributor, " + distributor.Name + " already exist"
	}
	return
}

// Validating distributor's parent presence
func validateDistributorParent(distributor Distributor) (isValid bool, errMsg string) {
	_, isParentDistributorAvailable := distributors[distributor.ParentDistributor]
	if distributor.ParentDistributor == "none" || distributor.ParentDistributor == "" || isParentDistributorAvailable {
		isValid = true
	} else {
		isValid = false
		errMsg = distributor.Name + "'s parent " + distributor.ParentDistributor + " doesn't exist"
	}
	return
}
