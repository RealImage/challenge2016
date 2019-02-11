package main

type userProcessor int

func (u *userProcessor) getUsers(creds credential) *user {

	return getUserFromUsers(creds.Username)

}
