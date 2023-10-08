package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// for setting the response
func (l *LocationData) SetAPIResponse(c *gin.Context, err string, data interface{}) {
	response := make(map[string]interface{})
	Status := http.StatusOK
	if err != "" {
		Status = http.StatusBadRequest
	}
	response["error"] = err
	response["data"] = data
	c.JSON(Status, response)
}
