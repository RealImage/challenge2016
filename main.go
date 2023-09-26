package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type DistributorPermission struct {
	distributor_name string
	isIncluded       bool
	isExcluded       bool
	city             string
	state            string
	country          string
}

func main() {
	var input_distributor_name string
	var input_region string
	authorizations, err := LoadPermissionData("permission_data.csv")
	if err != nil {
		fmt.Println("Not able to load permission data, Please check the CSV File")
		os.Exit(0)
	}
	test_datas := make(map[string][2]string)
	test_datas["test_data1"] = [2]string{"DISTRIBUTOR1", "CHICAGO-ILLINOIS-UNITEDSTATES"}
	test_datas["test_data2"] = [2]string{"DISTRIBUTOR1", "DELHI-INDIA"}
	test_datas["test_data3"] = [2]string{"DISTRIBUTOR1", "CHENNAI-TAMILNADU-INDIA"}
	test_datas["test_data4"] = [2]string{"DISTRIBUTOR1", "BANGALORE-KARNATAKA-INDIA"}
	test_datas["test_data5"] = [2]string{"DISTRIBUTOR1", "MYSORE-KARNATAKA-INDIA"}
	test_datas["test_data6"] = [2]string{"DISTRIBUTOR2", "TAMILNADU-INDIA"}
	test_datas["test_data7"] = [2]string{"DISTRIBUTOR2", "CHENNAI-TAMILNADU-INDIA"}
	test_datas["test_data8"] = [2]string{"DISTRIBUTOR3", "HUBLI-KARNATAKA-INDIA"}
	for _, v := range test_datas {
		input_distributor_name = v[0]
		input_region = v[1]
		fmt.Println("Distributor Name:", input_distributor_name)
		fmt.Println("Region:", input_region)
		result := HasPermission(input_distributor_name, input_region, authorizations)

		if result {
			fmt.Println("Permission: YES")
		} else {
			fmt.Println("Permission: NO")
		}
	}
}

func LoadPermissionData(file_path string) ([]DistributorPermission, error) {
	distributor_permissions := []DistributorPermission{}
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		var dp DistributorPermission
		if record[0] == "Distributor" {
			continue
		}
		city, state, country := split_region(record[2])
		dp.distributor_name = record[0]
		dp.isIncluded = record[1] == "INCLUDE"
		dp.isExcluded = record[1] == "EXCLUDE"
		dp.city = city
		dp.state = state
		dp.country = country
		distributor_permissions = append(distributor_permissions, dp)
	}
	return distributor_permissions, nil
}

func HasPermission(input_distributor_name string, input_region string, auths []DistributorPermission) bool {
	city, state, country := split_region(input_region)
	isHavingPermission := false
	if country != "" {
		isCountryHavePermission := validate_country_permission(input_distributor_name, country, auths)
		if state != "" && isCountryHavePermission {
			isStateHavePermission := validate_state_permission(input_distributor_name, isCountryHavePermission, country, state, auths)
			if city != "" && isStateHavePermission {
				isCityHavePermission := validate_city_permission(input_distributor_name, isStateHavePermission, country, state, city, auths)
				return isCityHavePermission
			} else {
				return isStateHavePermission
			}
		} else {
			return isCountryHavePermission
		}
	} else {
		return isHavingPermission
	}
}

func split_region(region string) (string, string, string) {
	parts := strings.Split(region, "-")
	city, state, country := "", "", ""
	if len(parts) == 3 {
		city, state, country = parts[0], parts[1], parts[2]
	} else if len(parts) == 2 {
		state, country = parts[0], parts[1]
	} else {
		country = parts[0]
	}
	return city, state, country
}

func validate_country_permission(input_distributor_name string, country string, auths []DistributorPermission) bool {
	isPermission := false
	for _, data := range auths {
		if data.distributor_name == input_distributor_name && data.country == country && data.state == "" && data.city == "" {
			if data.isIncluded {
				isPermission = true
			}
			if data.isExcluded {
				isPermission = false
			}
		}
	}
	return isPermission
}

func validate_state_permission(input_distributor_name string, isCountryHavePermission bool, country string, state string, auths []DistributorPermission) bool {
	isPermission := false
	isDataChanged := false
	for _, data := range auths {
		if data.distributor_name == input_distributor_name && data.country == country && data.state == state && data.city == "" {
			if data.isIncluded {
				isPermission = true
				isDataChanged = true
			}
			if data.isExcluded {
				isPermission = false
				isDataChanged = true
			}
		}
	}
	if isDataChanged {
		return isPermission
	} else {
		return isCountryHavePermission
	}
}

func validate_city_permission(input_distributor_name string, isStateHavePermission bool, country string, state string, city string, auths []DistributorPermission) bool {
	isPermission := false
	isDataChanged := false
	for _, data := range auths {
		if data.distributor_name == input_distributor_name && data.country == country && data.state == state && data.city == city {
			if data.isIncluded {
				isPermission = true
				isDataChanged = true
			}
			if data.isExcluded {
				isPermission = false
				isDataChanged = true
			}
		}
	}
	if isDataChanged {
		return isPermission
	} else {
		return isStateHavePermission
	}
}
