package routes

import (
	"qube-cinemas-challenge/controller"

	"github.com/gin-gonic/gin"
)


func Routes(c *gin.Engine){
	c.GET("/locations", controller.GetLocations)
}