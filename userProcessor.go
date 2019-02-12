package main

import "errors"

type userProcessor int

func (u *userProcessor) getUsers(creds *credential) *user {

	return getUserFromUsers(creds.Username)

}

func (u *userProcessor) createUser(loginUsername string, newUser *user) error {

	_, isUserExist := credentialsObject.getFromCredentialMap(newUser.Name)
	if isUserExist {
		return errors.New(nameAlreadyExists)
	}

	loginUser := getUserFromUsers(loginUsername)

	newHash, err := u.getRoleBasedValidationAndFields(loginUser, newUser)
	if err != nil {
		return err
	}

	parentUser := u.getNewUserParent(loginUser, newUser)
	if parentUser == nil {
		return errors.New(invalidParent)
	}

	err = u.validateIncludeAndExcludeLocations(parentUser, newUser)
	if err != nil {
		return err
	}

	creds := credential{}
	creds.Username = newUser.Name
	creds.EncryptedPassword = newHash
	credentialsObject.putIntoCredentialMap(creds)

	u.appendUsers(parentUser, newUser)

	return nil
}

func (u *userProcessor) appendUsers(parentUser, newUser *user) {

	if parentUser.Role == adminRole && newUser.Role == adminRole {

		users = append(users, newUser)
		return
	}

	parentUser.Children = append(parentUser.Children, newUser)

}

func (u *userProcessor) getRoleBasedValidationAndFields(loginUser, newUser *user) (newHash []byte, err error) {

	newUserRole := newUser.Role
	loginUserRole := loginUser.Role

	if loginUserRole == distributorRole && newUserRole == adminRole {
		err = errors.New(notAuthorizedToCreateAdmin)
		return
	}

	newHash = defaultDistributorHash

	if newUserRole == adminRole {
		if newUser.Parent != "" {
			newUser.Parent = ""
		}
		if len(newUser.Includes) != 0 {
			newUser.Includes = []location{}
		}

		if len(newUser.Excludes) != 0 {
			newUser.Excludes = []location{}
		}
		newHash = defaultAdminHash
	}

	if len(newUser.Children) != 0 {
		newUser.Children = []*user{}
	}

	return

}

func (u *userProcessor) getNewUserParent(loginUser, newUser *user) *user {

	newUserParent := newUser.Parent
	if newUserParent == "" {
		return loginUser
	}

	loginUserName := loginUser.Name
	if newUserParent == loginUserName {
		return loginUser
	}

	_, ok := credentialsObject.getFromCredentialMap(newUserParent)
	if !ok {
		return nil
	}

	parent := getUserFromUsersHelper(loginUser, newUserParent)
	if parent != nil {
		return parent
	}

	return nil

}

func (u *userProcessor) validateIncludeAndExcludeLocations(parentUser, newUser *user) (err error) {

	if len(newUser.Includes) == 0 {
		return nil
	}

	checkUser := parentUser

	if parentUser.Role == adminRole {
		checkUser = nil
	}

	err = u.validateLocations(newUser.Includes)
	if err != nil {
		return
	}

	err = u.validateLocations(newUser.Excludes)
	if err != nil {
		return
	}

	for _, loc := range newUser.Includes {
		if !isValidLocation(checkUser, loc) {
			err = errors.New(invalidIncludeLocation + loc.String())
			return
		}
	}

	for _, loc := range newUser.Excludes {
		if !isValidLocation(checkUser, loc) {
			err = errors.New(invalidExcludeLocation + loc.String())
			return
		}

	}

	return
}

func (u *userProcessor) validateLocations(inLoc []location) error {

	for _, loc := range inLoc {
		err := validateLocation(loc)
		if err != nil {
			return err
		}
	}

	return nil

}
