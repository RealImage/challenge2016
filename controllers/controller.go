package controllers

import ("github.com/gin-gonic/gin"
		//"challenge2016/"
		)

func AllDistrubter(c *gin.Context){
	
	c.String(200, "Your list is here")

	return
}