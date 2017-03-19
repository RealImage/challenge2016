package main

import (
	"flag"
	"log"
	"os"

	"./distribution"
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Unable to run application due to %+v", r)
			os.Exit(1)
		}
	}()

	distributor1 := &distribution.Distributor{}
	distributor1.Initialize("DISTRIBUTOR1", nil)
	distributor1.Include("IN")
	distributor1.Exclude("KA-IN")
	distributor1.Exclude("CENAI-TN-IN")

	distributor2 := &distribution.Distributor{}
	distributor2.Initialize("DISTRIBUTOR2", distributor1)
	log.Println(distributor2.Permissions())
	distributor2.Include("IN")
	distributor2.Exclude("TN-IN")
	log.Println(distributor2.Permissions())

	distributor3 := &distribution.Distributor{}
	distributor3.Initialize("DISTRIBUTOR3", distributor2)
	log.Println(distributor3.Permissions())
	err := distributor3.Include("HBALI-KA-IN")
	log.Println(err)
	log.Println(distributor3.Permissions())
}
