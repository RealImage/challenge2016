package main

import (
	"log"
)

func initAdminCredsAndHashes() (err error) {

	adminUser := user{
		Name: "admin",
		Role: adminRole,
	}

	defaultAdminHash, err = getPasswordHash("admin")
	if err != nil {
		log.Fatal(err)
	}

	defaultDistributorHash, err = getPasswordHash("distributor")
	if err != nil {
		log.Fatal(err)
	}

	adminCredential := credential{Username: adminUser.Name, EncryptedPassword: defaultAdminHash}

	err = credentialsObject.putIntoCredentialMap(adminCredential)
	if err != nil {
		return err
	}
	users = append(users, &adminUser)

	return nil

}

func (c *credentials) putIntoCredentialMap(creds credential) error {

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.credentialMap[creds.Username] = creds
	return nil

}

func (c *credentials) getFromCredentialMap(authToken string) (credential, bool) {

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if creds, ok := c.credentialMap[authToken]; ok {
		return creds, true
	}

	return credential{}, false
}

func (c *credentials) deleteFromCredentialMap(creds credential) {

	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.credentialMap, creds.Username)

}
