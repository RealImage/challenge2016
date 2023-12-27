package main

import "fmt"

func ViewDistributors() {
	fmt.Println("Distributors:")
	fmt.Println("--------------------------------------------------\n")
	for _, distributor := range distributors {
		fmt.Println(distributor.Name)
		fmt.Println("Permissions:", distributor.Permissions)
		isSubDistributor := (distributor.Parent != "")
		fmt.Println("isSubDistributor:", isSubDistributor)
		if isSubDistributor {
			fmt.Println("Parent:", distributor.Parent)
		}
		fmt.Println("--------------------------------------------------\n")
	}
}

func AddDistributor() {
	fmt.Println("Enter Distributor Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter Distributor Parent Name:")
	var parent string
	fmt.Scanln(&parent)

	fmt.Println("Enter Distributor Permissions:")
	var permissions Permission
	fmt.Println("Enter Include Regions:")
	var include []string
	for {
		var region string
		fmt.Scanln(&region)
		if region == "" {
			break
		}
		include = append(include, region)
	}
	permissions.Include = include

	fmt.Println("Enter Exclude Regions:")
	var exclude []string
	for {
		var region string
		fmt.Scanln(&region)
		if region == "" {
			break
		}
		exclude = append(exclude, region)
	}
	permissions.Exclude = exclude

	distributor := Distributor{
		Name:        name,
		Parent:      parent,
		Permissions: permissions,
	}

	distributors[name] = distributor
	distributorPermissions[name] = permissions
	fmt.Println("Distributor Added Successfully!!")
	fmt.Println("--------------------------------------------------\n")
}
