package main

import (
	"challenge2016/cli"
	"log"
)

func main() {
	cl := cli.PermissionsCli{}
	err := cl.Run()
	if err != nil {
		log.Fatalf("Finished with error: %s", err)
	}
}
