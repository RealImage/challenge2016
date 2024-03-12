package main

import (
	"context"
	"distributor/constants"
	"distributor/controller"
	util "distributor/utils"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env values%s", err)
	} else {
		fmt.Println("Env values loaded")
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		port = constants.DEFAULT_PORT
	}

	err = util.LoadLocations("cities.csv")
	if err != nil {
		fmt.Printf("Getting error in loading Csv file%s", err)
	}
	ctx := context.Background()
	router := controller.NewRouter(ctx)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Printf("Failed to serve%s", err)
	}
}
