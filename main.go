package main

import (
	"github.com/binkkatal/challenge2016/distribution"
)

func init() {
	distribution.UniversalAreaList = distribution.RetrieveAreas("./cities.csv")
}

func main() {

	distribution.GetInput()

}
