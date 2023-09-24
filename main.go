package main

import (
	"distributor/estimation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	result := estimation.SetDetails()
	fmt.Print(len(result))
	router := gin.Default()
	router.POST("/getPermission", finalResult())
	router.Run(":8080")
}

func finalResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input estimation.Distributor
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		if input.ParentDistributor == "" {
			input.Name = "Distributor1"
		}
		var result string
		result1, result2 := estimation.CheckDistribution(input)
		if result1 != "" {
			result = result1
		} else {
			result = result2
		}
		ctx.JSON(http.StatusOK, result)

	}
}
