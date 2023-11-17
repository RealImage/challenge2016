package main

import (
	datacsv "qube-cinemas-challenge/data-csv"
	"qube-cinemas-challenge/routes"

	"github.com/gin-gonic/gin"
)

func init(){
	datacsv.DataFetch()
}

func main(){
	r:= gin.Default()
	
	routes.Routes(r)

	r.Run()
}