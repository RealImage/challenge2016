package main

import (
	"fmt"
	"github.com/RealImage/challenge2016/service"
	"log"
	"net/http"
	"time"
)

func InitServer() {
	//initialise the app, config and load the dataset into memory
	app := service.NewApp()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("started server on port:", app.Config.Port)

	err := server.ListenAndServe()
	log.Fatalln(err)
}
