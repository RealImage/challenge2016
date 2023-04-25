package handler

import (
	"RealImageSolution/models"
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ReadCitiesCSV(filename string) []models.City {
	var cities []models.City

	// Read the CSV file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		city := models.City{
			CityCode:     line[0],
			ProvinceCode: line[1],
			CountryCode:  line[2],
			CityName:     line[3],
			ProvinceName: line[4],
			CountryName:  line[5],
		}
		cities = append(cities, city)
	}

	return cities
}

func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}

	return strings.ToLower(string(rr))
}

func AddDistributor(id *int, DistributorsList *[]models.Distributor) {
	*id++
	var name string
	fmt.Println("")
	fmt.Println("->Enter Distributor Name: ")
	fmt.Scanln(&name)
	var distributor models.Distributor
	distributor.ID = *id
	distributor.Name = name
	*DistributorsList = append(*DistributorsList, distributor)
	fmt.Println("->Now Add Permissions for ", distributor.Name)
	for {
		var permission string
		fmt.Println("Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA")
		// Permissions for DISTRIBUTOR1
		// INCLUDE: INDIA
		// INCLUDE: UNITEDSTATES
		// EXCLUDE: KARNATAKA-INDIA
		// EXCLUDE: CHENNAI-TAMILNADU-INDIA
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		permission = scanner.Text()
		if permission == "4" {
			(*DistributorsList)[*id-1].Include = distributor.Include
			(*DistributorsList)[*id-1].Exclude = distributor.Exclude
			break
		}

		data := strings.Split(permission, ":")
		prefix := strings.TrimSpace(data[0])
		sufix := strings.TrimSpace(data[1])

		switch prefix {
		case "INCLUDE":
			distributor.Include = append(distributor.Include, sufix)
		case "EXCLUDE":
			distributor.Exclude = append(distributor.Exclude, sufix)
		default:
			fmt.Println("Invalid Choice, Try Again!")
		}
	}
}

func PrintDistList(DistributorsList *[]models.Distributor) {
	fmt.Println("Distributor List: ")
	for _, distributor := range *DistributorsList {
		fmt.Println(distributor.ID, ") "+distributor.Name+" has permission to access: ")
		fmt.Println("INCLUDE: " + strings.Join(distributor.Include, ", "))
		fmt.Println("EXCLUDE: " + strings.Join(distributor.Exclude, ", "))
		fmt.Println("Distribution Parent: " + distributor.Parent)
		fmt.Println("Distribution Children: " + distributor.Child)
		fmt.Println("")
	}
}

func CheckDistributorPermission(cities *[]models.City, DistributorsList *[]models.Distributor) {

	fmt.Println("Check Permission for a Distributor")
	for {
		var name string
		fmt.Println("Enter Distributor Name: or press 4 for Main menu")
		fmt.Scanln(&name)

		if name == "4" {
			break
		}
		ans, err := CheckPermission(*cities, *DistributorsList, name)
		if err != nil {
			fmt.Println(err)
		}
		if ans {
			fmt.Println("YES")
			fmt.Println("Distributor: ", name, " has permission!")
		} else {
			fmt.Println("NO")
			fmt.Println("Distributor: ", name, " doesn't has permission!")
		}
	}
}

func CheckPermission(cities []models.City, DistributorsList []models.Distributor, distName string) (bool, error) {
	var id int
	for i, distributor := range DistributorsList {
		if distributor.Name == distName {
			id = i
			break
		}
	}

	// Check if distributor exists
	if id == 0 && DistributorsList[id].Name != distName {
		return false, errors.New("Distributor not found")
	}

	fmt.Println("-> Make a entry for the city you want to check: | ex: INDIA, FRANCE, KARNATKA-INDIA, etc...")
	// City Code,Province Code,Country Code,City Name,Province Name,Country Name
	// format -> CITY-PROVINCECODE-COUNTRYCODE-CITYNAME-PROVINCENAME-COUNTRYNAME
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	checkString := scanner.Text()
	var data []string
	if strings.Contains(checkString, "-") {
		data = strings.Split(checkString, "-")
	} else {
		data = append(data, checkString)
	}

	if len(data) == 1 {
		fmt.Println("Checking for Country")
		for _, city := range cities {
			city.CountryName = RemoveSpace(city.CountryName)
			data[0] = RemoveSpace(data[0])
			if city.CountryName == data[0] {
				for _, include := range DistributorsList[id].Include {
					include = RemoveSpace(include)
					city.CountryName = RemoveSpace(city.CountryName)
					if include == city.CountryName {
						return true, nil
					}
				}
				for _, exclude := range DistributorsList[id].Exclude {
					exclude = RemoveSpace(exclude)
					city.CountryName = RemoveSpace(city.CountryName)
					if exclude == city.CountryName {
						return false, nil
					}
				}
			}
		}
	} else if len(data) == 2 {
		fmt.Println("Checking for ProvinceName & CountryName")
		for _, city := range cities {
			city.ProvinceName = RemoveSpace(city.ProvinceName)
			city.CountryName = RemoveSpace(city.CountryName)
			data[0] = RemoveSpace(data[0])
			data[1] = RemoveSpace(data[1])
			if city.ProvinceName == data[0] && city.CountryName == data[1] {
				for _, include := range DistributorsList[id].Include {
					include = RemoveSpace(include)
					field := RemoveSpace(city.ProvinceName) + "-" + RemoveSpace(city.CountryName)
					if include == field {
						return true, nil
					}
				}
				for _, exclude := range DistributorsList[id].Exclude {
					exclude = RemoveSpace(exclude)
					field := RemoveSpace(city.ProvinceName) + "-" + RemoveSpace(city.CountryName)
					if exclude == field {
						return false, nil
					}
				}
			}
		}
	} else if len(data) == 3 {
		fmt.Println("Checking for CityName, ProvinceName & CountryName")
		for _, city := range cities {
			city.CityName = RemoveSpace(city.CityName)
			city.ProvinceName = RemoveSpace(city.ProvinceName)
			city.CountryName = RemoveSpace(city.CountryName)
			data[0] = RemoveSpace(data[0])
			data[1] = RemoveSpace(data[1])
			data[2] = RemoveSpace(data[2])
			if city.CityName == data[0] && city.ProvinceName == data[1] && city.CountryName == data[2] {
				for _, include := range DistributorsList[id].Include {
					include = RemoveSpace(include)
					field := RemoveSpace(city.CityName) + "-" + RemoveSpace(city.ProvinceName) + "-" + RemoveSpace(city.CountryName)
					if include == field {
						return true, nil
					}
				}
				for _, exclude := range DistributorsList[id].Exclude {
					exclude = RemoveSpace(exclude)
					field := RemoveSpace(city.CityName) + "-" + RemoveSpace(city.ProvinceName) + "-" + RemoveSpace(city.CountryName)
					if exclude == field {
						return false, nil
					}
				}
			}
		}
	}
	return false, nil
}

func DistributorNetwork(DistributorsList *[]models.Distributor) {
	fmt.Println("")
	fmt.Println("### Create a network between Distributors ###")
	fmt.Println("-> Enter the name of the Distributors which you want to connect in this FORMAT: childDistributor<-parentDistributor | ex : Dist2<-Dist1")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	// Split the input into parent and child distributors
	var parentName, childName string
	if strings.Contains(name, "<-") {
		data := strings.Split(name, "<-")
		parentName = strings.TrimSpace(data[1])
		childName = strings.TrimSpace(data[0])
	} else {
		fmt.Println("Invalid Format")
		return
	}

	// Find the parent and child distributors in the list
	var parentDistributor, childDistributor *models.Distributor
	for i := range *DistributorsList {
		d := &(*DistributorsList)[i]
		if d.Name == parentName {
			parentDistributor = d
		}
		if d.Name == childName {
			childDistributor = d
		}
	}

	// If either the parent or child distributors couldn't be found, return an error
	if parentDistributor == nil {
		fmt.Println("Parent distributor not found")
		return
	}
	if childDistributor == nil {
		fmt.Println("Child distributor not found")
		return
	}

	childDistributor.Include = append(childDistributor.Include, parentDistributor.Include...)
	childDistributor.Exclude = append(childDistributor.Exclude, parentDistributor.Exclude...)
	childDistributor.Parent = parentDistributor.Name
	parentDistributor.Child = childDistributor.Name
	fmt.Println("")
	fmt.Printf("Added network connection between %s and %s\n", parentDistributor.Name, childDistributor.Name)
}
