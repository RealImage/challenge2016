package routes

import ("github.com/gin-gonic/gin"
		"challenge2016/controllers"
	)

func InitialiseRoutes(router *gin.Engine){
	router.GET("/", controllers.Home)
	router.GET("/check-permissions", controllers.CheckPermissions)
	router.POST("/add-distributor", controllers.AddDistributor)
	router.NoRoute(func(c *gin.Context){
		c.JSON(404, gin.H{"error":"Not Found"})
	})
}

