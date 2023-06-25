package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Permission struct {
	Distributor string
	Permission  string
	Region      string
}

func main() {
	filePath := "permissions.csv"

	permissions, err := readCSV(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the distributor: ")
	var distributor string
	_, err = fmt.Scanln(&distributor)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the region: ")
	var region string
	_, err = fmt.Scanln(&region)
	if err != nil {
		log.Fatal(err)
	}

	hasPermission := checkPermission(permissions, distributor, region)
	if hasPermission {
		fmt.Println("YES, distribution is allowed in", region)
	} else {
		fmt.Println("NO, distribution is not allowed in", region)
	}
}

func readCSV(filePath string) ([]Permission, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var permissions []Permission
	for _, row := range rows {
		permission := Permission{
			Distributor: row[0],
			Permission:  row[1],
			Region:      row[2],
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func checkPermission(permissions []Permission, distributor, region string) bool {
	for _, permission := range permissions {
		if permission.Distributor == distributor && permission.Permission == "INCLUDE" && permission.Region == region {
			return true
		}
		if permission.Distributor == distributor && permission.Permission == "EXCLUDE" && permission.Region == region {
			return false
		}
	}
	return true
}
