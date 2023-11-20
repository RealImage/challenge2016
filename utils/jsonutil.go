package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ResponseJson struct {
	Status      string      `json:"status"`
	Result      interface{} `json:"result,omitempty"`
	ErrorString string      `json:"error,omitempty"`
}

// WriteResponseJson - a utility function for writing response json
func WriteResponseJson(w http.ResponseWriter, status int, data any, errorString string) {
	response := ResponseJson{
		Status:      strconv.Itoa(status),
		Result:      data,
		ErrorString: errorString,
	}

	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		response.Result = err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
