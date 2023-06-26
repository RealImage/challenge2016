/*
	permissions.csv file have the permissions we can edit it according to our requirement

 	build the executable with "go build film-distribution.go" command

	execute the executable file by "./film-distribution permissions.csv DISTRIBUTOR1 CHENNAI-TAMILNADU-INDIA"

	The program will output either "YES" or "NO" indicating whether the distributor has permission
	for the specified region.

*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Permissions struct {
	Include []string
	Exclude []string
}

type Distributor struct {
	Name        string
	Permissions Permissions
	Children    []*Distributor
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: film-distribution <permissions_file> <distributor_name> <region>")
	}

	permissionsFile := os.Args[1]
	distributorName := os.Args[2]
	region := os.Args[3]

	distributorMap := make(map[string]*Distributor)

	// Load and process the permissions file
	err := loadDistributors(permissionsFile, distributorMap)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Check if the distributor has permission for the specified region
	hasPermission := checkPermission(distributorMap, distributorName, region)
	if hasPermission {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func loadDistributors(filename string, distributorMap map[string]*Distributor) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error: Unable to close the file")
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		distributorName := strings.TrimSpace(record[0])
		permissionType := strings.TrimSpace(record[1])
		region := strings.TrimSpace(record[2])

		perm := Permissions{}
		switch permissionType {
		case "INCLUDE":
			perm.Include = append(perm.Include, region)
		case "EXCLUDE":
			perm.Exclude = append(perm.Exclude, region)
		}

		// Check if the distributor is a sub-distributor
		if strings.Contains(distributorName, "<") {
			parentName := strings.TrimSpace(strings.Split(distributorName, "<")[1])
			subDistributorName := strings.TrimSpace(strings.Split(distributorName, "<")[0])
			subDistributor := getOrCreateDistributor(distributorMap, subDistributorName)
			parentDistributor := getOrCreateDistributor(distributorMap, parentName)
			parentDistributor.Children = append(parentDistributor.Children, subDistributor)
			subDistributor.Permissions = mergePermissions(subDistributor.Permissions, perm)
		} else {
			// Create or get the distributor node and assign permissions
			node := getOrCreateDistributor(distributorMap, distributorName)
			node.Permissions = mergePermissions(node.Permissions, perm)
		}
	}

	return nil
}

func getOrCreateDistributor(distributorMap map[string]*Distributor, name string) *Distributor {
	if node, ok := distributorMap[name]; ok {
		return node
	}

	node := &Distributor{
		Name:        name,
		Permissions: Permissions{},
		Children:    make([]*Distributor, 0),
	}
	distributorMap[name] = node
	return node
}

func mergePermissions(permissions, newPermissions Permissions) Permissions {
	merged := Permissions{
		Include: append(permissions.Include, newPermissions.Include...),
		Exclude: append(permissions.Exclude, newPermissions.Exclude...),
	}
	return merged
}

func checkPermission(distributorMap map[string]*Distributor, distributorName, region string) bool {
	distributor, ok := distributorMap[distributorName]
	if !ok {
		return false
	}

	return checkRegionPermission(distributor, region)
}

func checkRegionPermission(distributor *Distributor, region string) bool {
	permissions := distributor.Permissions

	for _, included := range permissions.Include {
		if included == region {
			return true
		}
	}

	for _, excluded := range permissions.Exclude {
		if excluded == region {
			return false
		}
	}

	for _, child := range distributor.Children {
		if checkRegionPermission(child, region) {
			return true
		}
	}

	return false
}
