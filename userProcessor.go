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

	creds := credential{}
	creds.Username = newUser.Name
	creds.EncryptedPassword = newHash
	credentialsObject.putIntoCredentialMap(creds)

	u.appendUsers(parentUser, newUser)

	return nil
}

func (u *userProcessor) appendUsers(loginUser, newUser *user) {

	if loginUser.Role == adminRole && newUser.Role == adminRole {

		users = append(users, newUser)
		return
	}

	loginUser.Children = append(loginUser.Children, newUser)

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
