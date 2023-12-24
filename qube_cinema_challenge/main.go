package main

import (
	"dis1/distributor"
	"dis1/models"
	"dis1/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	locationFilePath := "cities.csv"

	err := utils.ReadLocations(locationFilePath)
	if err != nil {
		fmt.Println("Error parsing cities.csv: ", err)
		return
	}
	//taking distributor information from input json file
	var distributorList []distributor.Distributor
	if distributorInfo, err := ioutil.ReadFile("distributor_info.json"); err != nil {
		fmt.Println("Error while reading distributor data: ", err)
		return
	} else {
		distributorList = make([]distributor.Distributor, 0)
		if err = json.Unmarshal(distributorInfo, &distributorList); err != nil {
			fmt.Println("Error while decoding distributor data ", err)
			return
		}
	}

	distributor.CreateDistributors(distributorList)

	//taking enquiry information from user made json file
	enquiry := models.EnquiryForm{}
	if enquiryData, err := ioutil.ReadFile("enquiry.json"); err != nil {
		fmt.Println("error while reading enquiry data: ", err)
		return
	} else {
		if err = json.Unmarshal(enquiryData, &enquiry); err != nil {
			fmt.Println("Error while decoding enquiry data: ", err)
			return
		}
	}

	inquiredDistributer := distributor.GetDistributor(enquiry.DistributorName)
	inquiredLocation := enquiry.Location
	if authorized := inquiredDistributer.CheckPermissions(inquiredLocation); authorized {
		fmt.Printf("%s is authorized for distribution in %v", enquiry.DistributorName, inquiredLocation)
	} else {
		fmt.Printf("%s is not authorized for distribution in %v", enquiry.DistributorName, inquiredLocation)
	}
}
