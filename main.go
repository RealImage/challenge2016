/******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/
package main

import (
	rest "./restservice"
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Unable to run application due to %+v", r)
			os.Exit(1)
		}
	}()
	rest.StartService()
}
