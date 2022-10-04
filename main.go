package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"qube-cinemas/config"
	"qube-cinemas/utils"
)

var rows [][]string
var matchFound [][]string
var excluded [][]string
var subInclude [][]string
var subExclude [][]string

func main() {
	distributer := utils.NewDistributer{}
	distributer.Name = config.MainDistributer
	distributer.Include = config.Include
	distributer.Exclude = config.Exclude
	distributer.Check = config.Check

	distributer.Sub.Name = config.SubDistributer
	distributer.Sub.Include = config.SubInclude
	distributer.Sub.Check = config.SubCheck

	csvFile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	// var matches []string
	for _, val := range distributer.Include {
		matchFound = find(csvLines, val, true)
	}

	for _, val := range distributer.Exclude {
		excluded = find(matchFound, val, false)
	}
	allowed := check(excluded, distributer.Check)
	if distributer.Check != "" {
		fmt.Println("Primary distributer allowed in this region = ", allowed)
	}
	for _, val := range distributer.Sub.Include {
		subInclude = find(excluded, val, true)
	}
	for _, val := range distributer.Sub.Exclude {
		subExclude = find(excluded, val, true)
	}
	allowedSub := check(excluded, distributer.Sub.Check)
	if distributer.Sub.Check != "" {
		fmt.Println("Sub distributer allowed in this region = ", allowedSub)
	}
}

func find(records [][]string, val string, task bool) [][]string {
	if task == true {
		for _, row := range records {
			if row[5] == val {
				rows = append(rows, row)
			}
			if row[4] == val {
				rows = append(rows, row)
			}
			if row[3] == val {
				rows = append(rows, row)
			}
			// fmt.Println(row[3], val)
		}
	} else {
		for i, row := range records {
			if row[5] == val {
				rows = remove(rows, i)
			}
			if row[4] == val {
				rows = remove(rows, i)
			}
			if row[3] == val {
				rows = remove(rows, i)
			}
		}
	}
	return rows
}

func check(records [][]string, val string) (allowed bool) {
	for _, row := range records {
		if row[5] == val {
			return true
		}
		if row[4] == val {
			return true
		}
		if row[3] == val {
			return true
		}
	}
	return false
}

func remove(s [][]string, i int) [][]string {
	return append(s[:i], s[i+1:]...)
}
