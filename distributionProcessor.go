package main

import "errors"

type distributionProcessor int

func (d *distributionProcessor) validateDistribution(newLocation *location, loginUsername string) error {

	distributionUserName := newLocation.Username
	loginUser := getUserFromUsers(loginUsername)
	loginUserRole := loginUser.Role
	distributionUser := loginUser
	distributionUserRole := loginUserRole

	if distributionUserName == "" {
		distributionUserName = loginUsername
	} else {
		_, ok := credentialsObject.getFromCredentialMap(distributionUserName)
		if !ok {
			return errors.New(invalidCredentials)
		}
		distributionUser = getUserFromUsers(distributionUserName)
		distributionUserRole = distributionUser.Role
	}

	if loginUserRole == adminRole && distributionUserRole == adminRole {

		return errors.New(adminUsersCannotDistribute)
	}

	if loginUserRole == distributorRole && distributionUserRole == adminRole {

		return errors.New(notAuthorized)
	}

	parent := getUserFromUsersHelper(loginUser, distributionUserName)
	if parent == nil {
		return errors.New(notAuthorized)
	}

	if !isValidLocation(distributionUser, *newLocation) {
		return errors.New(notAuthorizedToDistribute)
	}

	return nil

}
