package main

import (
	"fmt"
	"github.com/RealImage/challenge2016/service"
	"net/http"
	"time"
)

func InitServer() {
	app := service.NewApp()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Println("started server on port : ", app.Config.Port)

	err := server.ListenAndServe()
	app.Logger.Fatal(err)
}
