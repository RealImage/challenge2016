package main

import (
	uuid "github.com/satori/go.uuid"
)

func (s *session) putIntoSessionMap(creds credential) (string, error) {

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	authToken := uuid.String()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sessionMap[authToken] = creds
	return authToken, nil

}

func (s *session) getFromSessionMap(authToken string) (credential, bool) {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if creds, ok := s.sessionMap[authToken]; ok {
		return creds, true
	}

	return credential{}, false
}
