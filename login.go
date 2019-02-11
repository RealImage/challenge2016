package main

import (
	"net/http"
)

type loginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, req *http.Request) {
	respondJSON(w, 200, loginStruct{"hiuser", "hipassword"})
}
