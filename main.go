package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	permissions, err := loadPermissions("cities.csv")
	if err != nil {
		fmt.Printf("Error loading permissions: %v\n", err)
		return
	}

	checkPermissions("check.csv", permissions)
}

func loadPermissions(filename string) (map[string]map[string]Action, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	permissions := make(map[string]map[string]Action)

	for _, record := range data {
		if len(record) < 6 {
			continue
		}

		region := fmt.Sprintf("%s-%s-%s", record[0], record[1], record[2])
		distributorName := record[3]
		action := Action(record[4])

		if _, ok := permissions[distributorName]; !ok {
			permissions[distributorName] = make(map[string]Action)
		}

		permissions[distributorName][region] = action
	}

	return permissions, nil
}

func checkPermissions(filename string, permissions map[string]map[string]Action) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening check file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading check file: %v\n", err)
		return
	}

	for _, record := range data {
		if len(record) < 3 {
			continue
		}

		distributorName := record[0]
		action := Action(record[1])
		region := record[2]

		if hasPermission(permissions, distributorName, action, region) {
			fmt.Printf("%s has permission to distribute in %s: YES\n", distributorName, region)
		} else {
			fmt.Printf("%s does not have permission to distribute in %s: NO\n", distributorName, region)
		}
	}
}

type Action string

func hasPermission(permissions map[string]map[string]Action, distributorName string, action Action, region string) bool {
	if perms, ok := permissions[distributorName]; ok {
		if act, ok := perms[region]; ok && act == action {
			return true
		}
	}
	return false
}
