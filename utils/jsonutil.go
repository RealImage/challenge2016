package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ResponseJson struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

func WriteResponseJson(w http.ResponseWriter, status int, data any) {
	response := ResponseJson{
		Status: strconv.Itoa(status),
		Result: data,
	}

	js, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		response.Result = err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
