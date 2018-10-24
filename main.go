package main

import (
	"fmt"
	"log"

	"github.com/souvikhaldar/challenge2016/file"
)

func main() {
	var path string
	if _, e := fmt.Scanf("%s", &path); e != nil {
		log.Panic(e)
	}
	if e := file.Readfile(path); e != nil {
		log.Panic(e)
	}
}
