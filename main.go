package main

import (
	"challengeQube/internal/handlers"
	"challengeQube/internal/jobs"
	"log"
	"net/http"
)

func main() {
	err := jobs.ParseCsvData()
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("listening")
	http.ListenAndServe(":8080", handlers.GetRouter())
}
