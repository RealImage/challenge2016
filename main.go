package main

import (
	"flag"
	"log"
	"os"

	"./application"
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Unable to run application due to %+v", r)
			os.Exit(1)
		}
	}()

	application.RunApplication()
}
