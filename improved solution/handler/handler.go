package handler

import (
	"RealImageSolution/models"
	"RealImageSolution/utils"
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type DistributorInterface interface {
	LoadCitiesFromCSV(filename string) (bool, error)
	AddDistributor(id *int)
	ListDistributors()
	CheckPermission()
	SetInclude(sufix string, id int)
	SetExclude(sufix string, id int)
	VerifyQuery(query string) bool
	CreateSubDistributorNetwork()
}

type DistributorsModel struct {
	CountryStateMap    map[string][]string
	StateCityMap       map[string][]string
	CurrentDistributor models.Distributor
	Distributors       []models.Distributor
}

// LoadCitiesFromCSV loads the cities from the csv file
func (d *DistributorsModel) LoadCitiesFromCSV(filename string) (bool, error) {
	// Read the CSV file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Error while reading the csv file")
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	// skip the header row
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	d.CountryStateMap = make(map[string][]string)
	d.StateCityMap = make(map[string][]string)

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		country := row[5]
		province := row[4]
		city := row[3]

		if _, ok := d.CountryStateMap[country]; !ok {
			d.CountryStateMap[country] = make([]string, 0)
		}

		if !utils.Contains(d.CountryStateMap[country], province) {
			d.CountryStateMap[country] = append(d.CountryStateMap[country], province)
		}

		if _, ok := d.StateCityMap[province]; !ok {
			d.StateCityMap[province] = make([]string, 0)
		}

		if !utils.Contains(d.StateCityMap[province], city) {
			d.StateCityMap[province] = append(d.StateCityMap[province], city)
		}
	}

	return true, nil
}

func (d *DistributorsModel) AddDistributor(id *int) {
	*id++
	var name string
	fmt.Println("")
	fmt.Println("->Enter Distributor Name: ")
	fmt.Scanln(&name)
	d.CurrentDistributor.ID = *id
	d.CurrentDistributor.Name = name
	d.Distributors = append(d.Distributors, d.CurrentDistributor)
	fmt.Println("->Now Add Permissions for ", d.CurrentDistributor.Name)
	for {
		var permission string
		fmt.Println("Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		permission = scanner.Text()
		if permission == "4" {
			break
		}

		data := strings.Split(permission, ":")
		prefix := strings.TrimSpace(data[0])
		sufix := strings.TrimSpace(data[1])

		switch prefix {
		case "INCLUDE":
			d.SetInclude(sufix, *id-1)
		case "EXCLUDE":
			d.SetExclude(sufix, *id-1)
		default:
			fmt.Println("Invalid Input")
		}
	}

	fmt.Println("->Distributor Added Successfully")
}

func (d *DistributorsModel) ListDistributors() {
	fmt.Println("->Distributor List: ")
	for _, distributor := range d.Distributors {
		fmt.Println("-> ", distributor.ID, ") "+distributor.Name+" has permission to access: ")
		fmt.Println("Permitted Places: ", len(distributor.PermittedPlaces))
		if distributor.SubDistributor {
			fmt.Println("Sub Distributor: YES")
		} else {
			fmt.Println("Sub Distributor: NO")
		}

		if distributor.Parent != "" {
			fmt.Println("Parent: " + distributor.Parent)
		} else {
			fmt.Println("Parent: NONE")
		}

		if len(distributor.Child) > 0 {
			fmt.Println("Children: " + strings.Join(distributor.Child, ","))
		} else {
			fmt.Println("Children: NONE")
		}
		fmt.Println("")
	}
}

func (d *DistributorsModel) CheckPermission() {

	for {
		fmt.Println("->Enter Distributor Name: or press 4 for Main menu")
		var name string
		fmt.Scanln(&name)

		if name == "4" {
			break
		}

		for _, dist := range d.Distributors {
			if dist.Name == name {
				d.CurrentDistributor = dist
			}
		}

		if d.CurrentDistributor.Name == "" {
			fmt.Println("Distributor not found")
			return
		}

		fmt.Println("->Enter your query to check permission: ")
		var query string
		fmt.Scanln(&query)

		ans := d.VerifyQuery(query)
		if ans {
			fmt.Println("")
			fmt.Println("YES")
			fmt.Println("")
		} else {
			fmt.Println("")
			fmt.Println("NO")
			fmt.Println("")
		}
	}
}

func (d *DistributorsModel) VerifyQuery(query string) bool {
	querySlice := strings.Split(query, "-")
	for _, include := range d.CurrentDistributor.PermittedPlaces {
		if include == querySlice[0] {
			return true
		}
	}

	return false
}

func (d *DistributorsModel) SetInclude(include string, id int) {

	for _, distributor := range d.Distributors {
		if distributor.ID == id {
			d.CurrentDistributor = distributor
		}
	}

	includeSlice := strings.Split(include, "-")

	switch len(includeSlice) {
	case 1:
		{
			log.Println("Inside len 1")
			d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, includeSlice[0])

			// storing the state in the distributor include
			d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, d.CountryStateMap[includeSlice[0]]...)

			// storing the city in the distributor include
			for _, state := range d.CountryStateMap[includeSlice[0]] {
				d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, d.StateCityMap[state]...)
			}
		}
	case 2:
		{
			// storing the state in the distributor include
			d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, includeSlice[0])

			// storing the city in the distributor include
			for _, state := range d.CountryStateMap[includeSlice[1]] {
				d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, d.StateCityMap[state]...)
			}
		}
	case 3:
		{
			// storing the state in the distributor include
			d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces, includeSlice[0])
		}
	default:
		fmt.Println("Invalid Input, Please try again")
	}

	d.Distributors[id].PermittedPlaces = d.CurrentDistributor.PermittedPlaces
}

func (d *DistributorsModel) SetExclude(exclude string, id int) {

	for _, distributor := range d.Distributors {
		if distributor.ID == id {
			d.CurrentDistributor = distributor
		}
	}

	excludeSlice := strings.Split(exclude, "-")

	switch len(excludeSlice) {
	case 1:
		{
			for i, value := range d.CurrentDistributor.PermittedPlaces {
				if value == excludeSlice[0] {
					for _, state := range d.CountryStateMap[value] {
						for _, city := range d.StateCityMap[state] {
							for j, value := range d.CurrentDistributor.PermittedPlaces {
								if j >= 0 && j < len(d.CurrentDistributor.PermittedPlaces) {
									if value == city {
										d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:j], d.CurrentDistributor.PermittedPlaces[j+1:]...)
									}
								}
							}
						}
						for j, value := range d.CurrentDistributor.PermittedPlaces {
							if j >= 0 && j < len(d.CurrentDistributor.PermittedPlaces) {
								if value == state {
									d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:j], d.CurrentDistributor.PermittedPlaces[j+1:]...)
								}
							}
						}
					}
					if i >= 0 && i < len(d.CurrentDistributor.PermittedPlaces) {
						d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:i], d.CurrentDistributor.PermittedPlaces[i+1:]...)
					}
				}
			}
		}
	case 2:
		{
			for i, value := range d.CurrentDistributor.PermittedPlaces {
				if value == excludeSlice[0] {
					d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:i], d.CurrentDistributor.PermittedPlaces[i+1:]...)
				}
			}

			for _, city := range d.StateCityMap[excludeSlice[1]] {
				for i, value := range d.CurrentDistributor.PermittedPlaces {
					if value == city {
						d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:i], d.CurrentDistributor.PermittedPlaces[i+1:]...)
					}
				}
			}
		}
	case 3:
		{
			// we need to remove the city from the distributor include
			for i, value := range d.CurrentDistributor.PermittedPlaces {
				if value == excludeSlice[0] {
					d.CurrentDistributor.PermittedPlaces = append(d.CurrentDistributor.PermittedPlaces[:i], d.CurrentDistributor.PermittedPlaces[i+1:]...)
				}
			}
		}
	default:
		fmt.Println("Invalid Input, Please try again!")
	}

	d.Distributors[id].PermittedPlaces = d.CurrentDistributor.PermittedPlaces
}

func (d *DistributorsModel) CreateSubDistributorNetwork() {

	for {

		fmt.Println("-> Your are currently in the Distributor Network Creation Mode")
		fmt.Println("-> Create Sub Distributor")
		fmt.Println("-> Enter Name of the Sub Distributor: or press 'q' to exit")

		var name string
		fmt.Scanln(&name)

		if name == "q" {
			break
		}

		d.CurrentDistributor = models.Distributor{
			ID:              len(d.Distributors),
			Name:            name,
			SubDistributor:  true,
			PermittedPlaces: d.Distributors[len(d.Distributors)-1].PermittedPlaces,
		}

		fmt.Println("")
		fmt.Println("### Create a network between Distributors ###")
		fmt.Println("-> Enter the name of the Distributors which you want to connect in this FORMAT: childDistributor<-parentDistributor | ex : Dist2<-Dist1")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name2 := scanner.Text()

		if name == "q" {
			break
		}

		// Split the input into parent and child distributors
		var parentName, childName string
		if strings.Contains(name2, "<-") {
			data := strings.Split(name2, "<-")
			parentName = strings.TrimSpace(data[1])
			childName = strings.TrimSpace(data[0])
		} else {
			fmt.Println("Invalid Format")
			return
		}

		d.CurrentDistributor.Parent = parentName
		d.Distributors = append(d.Distributors, d.CurrentDistributor)

		// Check if the parent distributor exists
		d.Distributors[len(d.Distributors)-1].Child = append(d.Distributors[len(d.Distributors)-1].Child, childName)
		d.Distributors[len(d.Distributors)-1].SubDistributor = true

		fmt.Println("")
		fmt.Printf("Added network connection between parent -> %s and child->  %s\n", parentName, childName)
	}
}
