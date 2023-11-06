package main

import (
		"github.com/gin-gonic/gin"
		"challenge2016/routes")

func main(){
	router:=gin.Default() //Created a routeconst
	routes.InitialiseRoutes(router)
	router.Run()
}




