package main

import "errors"

type distributionProcessor int

func (d *distributionProcessor) validateDistribution(newLocation *location, loginUsername string) error {

	inUser := newLocation.Username
	loginUser := getUserFromUsers(loginUsername)
	loginUserRole := loginUser.Role
	distributionUser := loginUser
	distributionUserRole := loginUserRole
	distributionUserName := loginUsername

	if inUser != "" {
		distributionUserName = inUser
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

	if distributionUserName != loginUsername {
		child := getUserFromUsersHelper(loginUser, distributionUserName)
		if child == nil {
			return errors.New(notAuthorized)
		}
	}

	for {

		if !isValidLocation(distributionUser, *newLocation) {
			return errors.New(notAuthorizedToDistribute)
		}

		parent := getUserFromUsers(distributionUser.Parent)
		if parent.Role == adminRole {
			break
		}

		distributionUser = parent

	}

	return nil

}
