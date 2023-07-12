package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler ...
func (h *DistributorHandler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"serverAlive": true,
	})
}
