package controllers

import (
	"fmt"
	"strings"

	"realImage.com/m/model"
	"realImage.com/m/utility"
)

func AddSubDistributor() {
	parentName := ""
	fmt.Println("Enter the name of the Parent for whom child distributor will be added: ")
	if scanner.Scan() {
		parentName = scanner.Text()
	}

	parentDistributor, relation := utility.CheckDistributor(DistributorMap, parentName)
	if relation == "" {
		fmt.Println("Distributor does not exist")
		return
	}

	name := ""
	fmt.Println("Enter the name of the Sub Distributor to be added: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	subDistributor := model.Distributor{
		Name:            name,
		IncludeRegions:  []string{},
		ExcludeRegions:  []string{},
		SubDistributors: []*model.Distributor{},
	}

	UpdateIncludeAndExcludePermissions(&subDistributor, parentName, true)
	(DistributorMap[parentDistributor.Name])[subDistributor.Name] = subDistributor
	DistributorMap[subDistributor.Name] = make(map[string]model.Distributor)
	fmt.Println("Sub Distributor added under Parent Distributor ", parentName)
}

func AddDistributor() {
	name := ""
	fmt.Println("Enter name of the distributor to be added: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	if name == "" {
		fmt.Println("Distributor name cannot be empty")
		return
	}
	distributor := model.Distributor{
		Name:            name,
		IncludeRegions:  []string{},
		ExcludeRegions:  []string{},
		SubDistributors: []*model.Distributor{},
	}
	UpdateIncludeAndExcludePermissions(&distributor, "", false)
	DistributorMap[name] = map[string]model.Distributor{
		name: distributor,
	}
}

func CheckForAccess() {
	name := ""
	fmt.Println("Enter the name of the distributor or Subdistributor for which we want to check access: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	distributor, rel := utility.CheckDistributor(DistributorMap, name)
	if rel == "" {
		fmt.Println("Distributor does not exist")
		return
	}

	var city, state, country string

	fmt.Println("Search by Location ::::::::::")
	fmt.Println("1. Search By City")
	fmt.Println("2. Search By state")
	fmt.Println("3. Search By country")
	fmt.Println()
	var selection string
	if scanner.Scan() {
		selection = scanner.Text()
	}

	switch selection {
	case "1":
		fmt.Println("Enter City to be checked: ")
		if scanner.Scan() {
			city = scanner.Text()
		}
		city = strings.TrimSpace(city)
		if utility.HasAccess(distributor, city, "city", LocMap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	case "2":
		fmt.Println("Enter State to be checked: ")
		if scanner.Scan() {
			state = scanner.Text()
		}
		if utility.HasAccess(distributor, state, "state", LocMap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	case "3":
		fmt.Println("Enter Country to be checked: ")
		if scanner.Scan() {
			country = scanner.Text()
		}
		if utility.HasAccess(distributor, country, "country", LocMap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	default:
		fmt.Println("Invalid option.")
	}
}

func UpdateIncludeAndExcludePermissions(distributor *model.Distributor, parent string, isSub bool) {
	inclusions := distributor.IncludeRegions
	exclusions := distributor.ExcludeRegions

out:
	for {
		fmt.Println("Choose Option to execute")
		fmt.Println("1. Add Inclusions")
		fmt.Println("2. Add Exclusions")
		fmt.Println("3. Exit")
		var selection string
		fmt.Scanln(&selection)
		switch selection {
		case "1", "2":
			country := ""
			state := ""
			city := ""
			fmt.Println("Enter Country:")

			if scanner.Scan() {
				country = scanner.Text()
			}
			if country != "" {
				fmt.Println("Enter State:")

				if scanner.Scan() {
					state = scanner.Text()
				}
				if state != "" {
					fmt.Println("Enter City:")

					if scanner.Scan() {
						city = scanner.Text()
					}
				}
			}
			perms := ""
			if country != "" {
				perms += strings.ToLower(country)
			}
			if state != "" {
				perms = "_" + perms
				perms = strings.ToLower(state) + perms
			}
			if city != "" {
				perms = "_" + perms
				perms = strings.ToLower(city) + perms
			}
			if selection == "1" {
				inclusions = append(inclusions, perms)
			} else {
				exclusions = append(exclusions, perms)
			}

			if isSub {
				Parent := utility.FindParentDistributor(parent, DistributorMap)
				approved := false
				if country != "nil" && country != "" {
					if utility.HasAccess(Parent, country, "country", LocMap) {
						approved = true
					}
				}
				if state != "nil" && state != "" {
					if utility.HasAccess(Parent, state, "state", LocMap) {
						approved = true

					}
				}
				if city != "nil" && city != "" {
					if utility.HasAccess(Parent, city, "city", LocMap) {
						approved = true
					}
				}
				if approved {
					distributor.IncludeRegions = inclusions
					distributor.ExcludeRegions = exclusions
				} else {
					fmt.Println("Parent is NOT AUTHORISED in the region provided")
				}
			} else {
				distributor.IncludeRegions = inclusions
				distributor.ExcludeRegions = exclusions
			}
		case "3":
			fmt.Println("Exiting Permissions Update section")
			break out
		default:
			fmt.Println("Invalid option.")
			continue
		}
	}
}

func UpdatePermissions() {
	name := ""
	fmt.Println("Enter the name of the distributor to be added ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	distributor, rel := utility.CheckDistributor(DistributorMap, name)
	var isSub bool
	if rel == "child" {
		isSub = true
	}
	if rel != "" {
		UpdateIncludeAndExcludePermissions(&distributor, "", isSub)
		(DistributorMap[name])[distributor.Name] = distributor
	} else {
		fmt.Println(" Distributor does not exist")
	}
}
