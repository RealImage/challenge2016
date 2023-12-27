package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/challenge2016/model"
	"github.com/challenge2016/store"
	"github.com/challenge2016/util"
	"github.com/gin-gonic/gin"
)

var storage store.LocalStorage

func bootstrap() {
	now := time.Now()
	currentDate := now.Format("2006-01-02")
	file := fmt.Sprintf("%s.json", currentDate)
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		// File does not exist, create it
		ref, err := os.Create(file)
		if err != nil {
			log.Fatal("unable to create file")
		}
		storage = *store.NewLocalStorage(ref)
	}
	// File already exists, open it
	ref, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("unable to open file")
	}
	storage = *store.NewLocalStorage(ref)
}
func main() {
	bootstrap()
	router := gin.Default()
	router.POST("/distribution", func(ctx *gin.Context) {
		var req model.DistributionReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage.LoadData()
		permissions := make(map[string]model.Permissions)
		permissions[req.Name] = req.Permissions
		distribution := storage.Data
		distribution[req.Name] = req.Permissions
		storage.Data = distribution
		storage.SaveData()
		ctx.JSON(http.StatusOK, gin.H{"messaage": "distribution data stored successfully"})
	})
	router.GET("/check/:distributor/:region", func(c *gin.Context) {
		distributorKey := c.Param("distributor")
		region := strings.Split(c.Param("region"), "-")
		storage.LoadData()
		if strings.Contains(distributorKey, "<") {
			distributors := strings.Split(distributorKey, "<")
			i := len(distributors) - 1
			distributorPermissions, ok := storage.GetPermissions(distributors[i])
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "Distributor key not found"})
				return
			}
			var permissions model.Permissions
			for {
				i = i - 1
				if i < 0 {
					break
				}
				permissions, ok := distributorPermissions.ChildPermissions[distributors[i]]
				if ok {
					if permissions.ChildPermissions != nil {
						i = i - 1
					}
				}
				fmt.Println(permissions)
			}
			result := util.IsAuthorized(permissions, region)
			c.JSON(http.StatusOK, gin.H{"authorized": result})
		} else {
			distributorPermissions, ok := storage.GetPermissions(distributorKey)
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "Distributor key not found"})
				return
			}
			result := util.IsAuthorized(distributorPermissions, region)
			c.JSON(http.StatusOK, gin.H{"authorized": result})
		}
	})

	fmt.Println("Server is running on :8080")
	router.Run(":8080")
}
