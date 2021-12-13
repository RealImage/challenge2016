package service

import (
	"challenge2016/models"
	database "challenge2016/utils"
	"fmt"
	"strings"
)

//Function to get all the data and their permissions
func CheckAllDistributorData() (output []models.Distributors) {
	var distributors []models.Distributors
	sqlDb := database.DB
	sqlDb.Raw("select * from distributors").Find(&distributors)
	return distributors
}

// function to insert the data
func InsertDistributor(input models.Distributors) (output string, message string) {
	sqlDB := database.DB
	var subDisSeniority string
	var permissions models.Permissions
	if input.Subdistributor != "" {
		fmt.Println("running sub dis query")
		sqlDB.Raw("select seniority from distributors where distributorname = (?)", input.Subdistributor).Scan(&subDisSeniority)
		if subDisSeniority == "" {
			return "", "Such subdistributor does not exist"
		}

	}
	fmt.Println("permissions", permissions)
	// To add the subdistributor included permission to distributor

	if input.Seniority == "High" {
		input.Seniority = "3"
	} else if input.Seniority == "Medium" {
		input.Seniority = "2"
	} else if input.Seniority == "Low" {
		input.Seniority = "1"
	}
	fmt.Println("Distrib", subDisSeniority)
	if input.Subdistributor != "" && input.Seniority >= subDisSeniority {
		return "", "Cannot assign the Distributor as subDistributor"
	} else {
		sqlDB.Raw("select included,excluded from distributors where distributorname = (?)", input.Subdistributor).Scan(&permissions)
		splitPermissionIncluded := strings.Split(permissions.Included, ",")
		if len(splitPermissionIncluded) > 0 {
			fmt.Println("Inside splitPermissionIncluded")
			for j := range splitPermissionIncluded {
				if !strings.Contains(input.Included, splitPermissionIncluded[j]) {
					input.Included = input.Included + "," + splitPermissionIncluded[j]
				}
			}
		}

		splitPermissionexcluded := strings.Split(permissions.Excluded, ",")
		if len(splitPermissionexcluded) > 0 {
			fmt.Println("Inside splitPermissionexcluded")
			for j := range splitPermissionexcluded {
				if !strings.Contains(input.Excluded, splitPermissionexcluded[j]) {
					input.Excluded = input.Excluded + "," + splitPermissionexcluded[j]
				}
			}
		}
	}
	structureInput := models.Distributors{Id: input.Id, Distributorname: input.Distributorname, Included: input.Included, Excluded: input.Excluded, Subdistributor: input.Subdistributor, Seniority: input.Seniority}
	val := sqlDB.Select("distributorname", "included", "excluded", "subdistributor", "seniority").Create(&structureInput)
	fmt.Println("input", structureInput)
	fmt.Println("value", val.Error)
	if val.Error != nil {
		err := fmt.Sprintf("%s", val.Error)
		return "", err
	}
	return "Insertion done", ""
}

func Checkdistributorpermissions(input models.Data) string {
	fmt.Println("checkdistributorpermissions")
	sqlDB := database.DB
	permission := models.Permissions{}
	var flag = 0
	val := sqlDB.Raw("select excluded,included from distributors where distributorname = (?)", input.Distributorname).Scan(&permission)
	if val.Error != nil {
		fmt.Println("error", val.Error)
		err := fmt.Sprintf("%s", val.Error)
		return err
	}
	fmt.Println("val", val)
	fmt.Println("input", input, permission.Excluded)
	x := strings.Contains("BANGALORE-KARNATAKA-INDIA", "KARNATAKA-INDIA")
	fmt.Println("x", x)
	splitExcluded := strings.Split(permission.Excluded, ",")
	for i := range splitExcluded {
		if strings.Contains(splitExcluded[i], input.Permission) || strings.Contains(input.Permission, splitExcluded[i]) {
			flag++
		}
	}
	if flag > 0 {
		return "no"
	}
	splitIncluded := strings.Split(permission.Included, ",")
	for i := range splitIncluded {
		if strings.Contains(splitIncluded[i], input.Permission) || strings.Contains(input.Permission, splitIncluded[i]) {
			flag++
		}
	}
	if flag > 0 {
		return "yes"
	}
	fmt.Println("permission", permission)
	return "no"
}
