package main

import (
	"RealImage/distributor"
	"RealImage/models"
	"RealImage/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read CSV file with location data.
	locationFilePath := "cities.csv"

	err := utils.ReadLocations(locationFilePath)
	if err != nil {
		fmt.Println("Error reading locations:", err)
		return
	}

	// Parse the distributorInfo.json to retrieve the list of distributors
	var distributorList []distributor.Distributor
	if distributorInfo, err := ioutil.ReadFile("input/distributorInfo.json"); err != nil {
		fmt.Println(err)
		return
	} else {
		distributorList = make([]distributor.Distributor, 0)
		if err = json.Unmarshal(distributorInfo, &distributorList); err != nil {
			fmt.Println("Error while marshalling JSON ", err)
			return
		}
	}

	distributor.CreateDistributors(distributorList)

	// Parse the enquiry.json to check a distributor for distribution permissions
	enquiry := models.EnquiryForm{}
	if inquiryData, err := ioutil.ReadFile("input/enquiry.json"); err != nil {
		fmt.Println(err)
		return
	} else {
		if err = json.Unmarshal(inquiryData, &enquiry); err != nil {
			fmt.Println("Error while unmarshalling JSON ", err)
			return
		}
	}

	// Get distributor from name
	enquiredDistributer := distributor.GetDistributor(enquiry.DistributorName)
	enquiredLocation := enquiry.Location

	// Check for authorization
	if isAuthorized := enquiredDistributer.CheckPermissions(enquiredLocation); isAuthorized {
		fmt.Printf("%s is authorized for distribution in %v", enquiry.DistributorName, enquiredLocation)
	} else {
		fmt.Printf("%s is not authorized for distribution in %v", enquiry.DistributorName, enquiredLocation)
	}
}
