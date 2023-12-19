package main

import (
	"bufio"
	"fmt"
	"moviecinema/util"
	"os"
	"strings"
)

type Permission struct {
	PType       string
	Country     string
	Province    string
	City        string
	Distributor string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	authority := util.NewAuthority("abdul")

	cache, err := util.PopulateMapFromCSV("cities.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Print(cache)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add new permission")
		fmt.Println("2. Check permission")
		fmt.Println("3. Exit")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			var input, distributor string
			var permissionType string

			fmt.Print("Enter distributor: ")
			scanner.Scan()
			distributor = scanner.Text()
			dList := strings.Split(distributor, "<")
			// fmt.Println("dlist: ", dList)
			authority.RegisterDistributor(util.Distributor{Name: strings.TrimSpace(dList[0])})

			fmt.Printf("Enter permissions for [%v] (e.g., INCLUDE: HUBLI-KARNATAKA-INDIA EXCLUDE: TAMILNADU-INDIA) ", dList[0])
			scanner.Scan()
			input = scanner.Text()
			// fmt.Println("input: ", input)
			parts := strings.Split(input, ":")
			if len(parts) != 2 {
				fmt.Println("Invalid input format. Please use the format 'INCLUDE: COUNTRY-PROVINCE-CITY EXCLUDE: COUNTRY-PROVINCE-CITY'.")
				continue
			}

			permissionTypeStr := strings.TrimSpace(parts[0])
			input = strings.TrimSpace(parts[1])

			switch strings.ToUpper(permissionTypeStr) {
			case "INCLUDE":
				permissionType = "INCLUDE"
			case "EXCLUDE":
				permissionType = "EXCLUDE"
			default:
				fmt.Println("Invalid permission type. Please enter INCLUDE or EXCLUDE.")
				continue
			}

			parts = strings.Split(input, "-")
			var country, province, city string
			if len(parts) == 1 {
				country = strings.ToUpper(parts[0])

				if authority.IsCompatibeForNewPermissionForCountry(dList, country) {
					d, err := authority.GetDistributorByName(dList[0])
					if err != nil {
						fmt.Println(err)
						continue
					}
					authority.AddNewPermissionForEntireCountry(d, util.Country{Name: country})
				} else {
					fmt.Printf("No permission added for %v", dList[0])
				}

			} else if len(parts) == 2 {
				province, country = strings.ToUpper(parts[0]), strings.ToUpper(parts[1])
				if authority.IsCompatibeForNewPermissionForProvince(dList, province, country, permissionType) {
					d, err := authority.GetDistributorByName(dList[0])
					if err != nil {
						fmt.Println(err)
						continue
					}
					authority.AddNewPermissionForEntireProvince(d, util.Province{Name: province}, util.Country{Name: country}, permissionType)
				} else {
					fmt.Println("No permission added for this distributor")
				}

			} else if len(parts) == 3 {
				city, province, country = strings.ToUpper(parts[0]), strings.ToUpper(parts[1]), strings.ToUpper(parts[2])
				if authority.IsCompatibeForNewPermissionForCity(dList, city, province, country, permissionType) {
					d, err := authority.GetDistributorByName(dList[0])
					if err != nil {
						fmt.Println(err)
						continue
					}
					err = authority.AddNewPermissionForCity(d, util.City{Name: city}, util.Province{Name: province}, util.Country{Name: country}, permissionType)
					if err != nil {
						fmt.Println(err)
						continue
					}
				} else {
					fmt.Println("No permission added for this distributor")
				}
			} else {
				fmt.Println("Invalid String. Please enter this way country, province-country, city-province-country")
				continue
			}
		case 2:
			var input, d string

			fmt.Print("Enter distributor: ")
			fmt.Scan(&d)

			fmt.Print("Enter city or province or country  (e.g., HUBLI-KARNATAKA-INDIA): ")
			scanner.Scan()
			input = scanner.Text()

			parts := strings.Split(input, "-")
			distributor, err := authority.GetDistributorByName(d)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("please enter a valid distributor!!")
				continue
			}
			var country, province, city string
			// if only country is mentioned
			if len(parts) == 1 {
				country = strings.ToUpper(parts[0])
				if !util.IsValidCountry(country, cache) {
					fmt.Printf("Error: [%s] is not a valid country ", country)
					continue
				}
				if distributor.CheckPermissionForEntireCountry(util.Country{Name: country}) {
					fmt.Printf("Distributor: %s is having permission in [%s]", distributor.Name, country)

				} else {
					fmt.Printf("Distributor: %s is not having permission in [%s]", distributor.Name, country)
				}
				continue
				// if province in some country is mentioned
			} else if len(parts) == 2 {
				province, country = strings.ToUpper(parts[0]), strings.ToUpper(parts[1])
				if !util.IsValidProvince(country, province, cache) {
					fmt.Printf("Error: [%s] is not a valid province in [%s]", province, country)
					continue
				}
				if distributor.CheckPermissionForEntireProvince(util.Country{Name: country}, util.Province{Name: province}) {
					fmt.Printf("Distributor: %s is having permission in Entire province[%s]", distributor.Name, province)

				} else {
					fmt.Printf("Distributor: %s is not having permission in Entire province[%s]", distributor.Name, province)
				}
				continue
				// if a city is mentioned with province and country
			} else if len(parts) == 3 {
				city, province, country = strings.ToUpper(parts[0]), strings.ToUpper(parts[1]), strings.ToUpper(parts[2])
				if !util.IsValidCity(country, province, city, cache) {
					fmt.Printf("Error: [%s] is not a valid city in province [%s] in country[%s]", city, province, country)
					continue
				}
				if distributor.CheckPermissionForCity(util.Country{Name: country}, util.Province{Name: province}, util.City{Name: city}) {
					fmt.Printf("Distributor: %s is having permission for this city [%s]", distributor.Name, city)

				} else {
					fmt.Printf("Distributor: %s is not having permission for city [%s]", distributor.Name, city)
				}
				continue
			} else {
				fmt.Println("Invalid String. Please enter this way country, province-country, city-province-country")
				continue
			}

		case 3:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
	/*
		authority := util.NewAuthority("abdul")

		d := util.Distributor{
			Name: "Raju",
		}
		authority.RegisterDistributor(d)

		// fmt.Println(d.CheckPermissionForEntireCountry(util.Country{Name: "India"}))

		// fmt.Println(authority)
		err := authority.AddNewPermissionForEntireCountry(&d, util.Country{Name: "India"})

		if err != nil {
			fmt.Println(err)
			return
		}

		err = authority.AddNewPermissionForEntireCountry(&d, util.Country{Name: "America"})

		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Println(d)

		fmt.Println(d.CheckPermissionForEntireCountry(util.Country{Name: "India"})) // true
		// fmt.Println(d.CheckPermissionForEntireCountry(util.Country{Name: "Bangladesh"})) //--> false
		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Rajasthan"}))                  // true
		fmt.Println(d.CheckPermissionForCity(util.Country{Name: "India"}, util.Province{Name: "Delhi"}, util.City{Name: "Tugalkabad"})) // true

		err = authority.AddNewPermissionForEntireProvince(&d, util.Province{Name: "Rajasthan"}, util.Country{Name: "India"}, "EXCLUDE")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Rajasthan"})) // false

		err = authority.AddNewPermissionForEntireProvince(&d, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "EXCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = authority.AddNewPermissionForEntireProvince(&d, util.Province{Name: "Maxico"}, util.Country{Name: "America"}, "EXCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Delhi"})) // false

		err = authority.AddNewPermissionForEntireProvince(&d, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Delhi"})) // true

		err = authority.AddNewPermissionForEntireProvince(&d, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Printf("main:%v\n", d)

		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Delhi"})) // true

		// test for city

		err = authority.AddNewPermissionForCity(&d, util.City{Name: "Siri"}, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d.CheckPermissionForCity(util.Country{Name: "India"}, util.Province{Name: "Delhi"}, util.City{Name: "Siri"}))

		err = authority.AddNewPermissionForCity(&d, util.City{Name: "Tugalkabad"}, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "EXCLUDE")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d.CheckPermissionForCity(util.Country{Name: "India"}, util.Province{Name: "Delhi"}, util.City{Name: "Tugalkabad"}))
		fmt.Println(d.CheckPermissionForEntireProvince(util.Country{Name: "India"}, util.Province{Name: "Delhi"})) // FALSE

		// fmt.Printf("main:%v\n", d)

		d2 := util.Distributor{
			Name: "Abdul Aleem",
		}
		authority.RegisterDistributor(d2)

		//compatibility check for province
		fmt.Println(authority.CheckCompatibilityForCountry(&d2, &d, util.Country{Name: "India"}))                                           // <nil>
		fmt.Println(authority.CheckCompatibilityForProvince(&d2, &d, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE")) // some error
		fmt.Println(authority.CheckCompatibilityForProvince(&d2, &d, util.Province{Name: "UP"}, util.Country{Name: "India"}, "INCLUDE"))    // <NIL>
		fmt.Println(authority.CheckCompatibilityForProvince(&d2, &d, util.Province{Name: "Rajashthn"}, util.Country{Name: "India"}, "INCLUDE"))
		fmt.Println(authority.CheckCompatibilityForProvince(&d2, &d, util.Province{Name: "Maxico"}, util.Country{Name: "America"}, "INCLUDE"))

		// compatibility check for the cities

		// <nil>
		fmt.Println(authority.CheckCompatibilityForCity(&d2, &d, util.City{Name: "Tugalkabad"}, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE")) // some error
		fmt.Println(authority.CheckCompatibilityForCity(&d2, &d, util.City{Name: "Tugalkabad"}, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "EXCLUDE")) // <nil>
		fmt.Println(authority.CheckCompatibilityForCity(&d2, &d, util.City{Name: "New Delhi"}, util.Province{Name: "Delhi"}, util.Country{Name: "India"}, "INCLUDE"))  // <NIL>
		fmt.Println(authority.CheckCompatibilityForCity(&d2, &d, util.City{Name: "New Delhi"}, util.Province{Name: "UK"}, util.Country{Name: "Nagaland"}, "INCLUDE"))  // <NIL>
	*/

}
