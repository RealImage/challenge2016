// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Distributor struct {
	CountryIncludes []string
	CountryExcludes []string
	StateIncludes   []string
	StateExcludes   []string
	CityIncludes    []string
	CityExcludes    []string
}

var distributorMapper map[string]Distributor

func main() {
	ruleFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error while loading the file: " + err.Error())
		return
	}

	defer ruleFile.Close()
	distributorMapper = make(map[string]Distributor)
	getDistributorPermissions(ruleFile)

	inputFile, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Error while loading the input file: " + err.Error())
		return
	}
	defer inputFile.Close()
	computeAndWriteAnswers(inputFile, os.Args[3])
}

func computeAndWriteAnswers(inputFile io.Reader, outputFilePath string) {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Couldn't create the output file: " + err.Error())
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputText := scanner.Text()
		if inputText == "" {
			return
		}
		answer := getAnswer(inputText)
		_, err = io.WriteString(outputFile, answer+"\n")
		if err != nil {
			fmt.Println("Couldn't write to the output file: " + err.Error())
			return
		}
	}
}

func getAnswer(inputText string) string {
	inputTokens := strings.Split(inputText, " ")
	distributor := inputTokens[0]
	location := inputTokens[1]

	locationTokens := strings.Split(location, "-")
	if len(locationTokens) == 1 {
		return checkCountryPermit(distributor, location)
	}
	if len(locationTokens) == 2 {
		return checkStatePermit(distributor, location)
	}
	if len(locationTokens) == 3 {
		return checkCityPermit(distributor, location)
	}
	return ""
}

func checkCountryPermit(distributor string, country string) string {
	for _, v := range distributorMapper[distributor].CountryExcludes {
		if v == country {
			return "NO"
		}
	}
	if len(distributorMapper[distributor].CountryIncludes) == 1 && distributorMapper[distributor].CountryIncludes[0] == "All" {
		return "YES"
	}

	for _, v := range distributorMapper[distributor].CountryIncludes {
		if v == country {
			return "YES"
		}
	}
	return "NO"
}

func checkStatePermit(distributor string, location string) string {
	locationTokens := strings.Split(location, "-")
	countryPermit := checkCountryPermit(distributor, locationTokens[1])
	if countryPermit == "NO" {
		return "NO"
	}

	for _, v := range distributorMapper[distributor].StateExcludes {
		if v == location {
			return "NO"
		}
	}
	for _, v := range distributorMapper[distributor].StateIncludes {
		if v == location {
			return "YES"
		}
	}
	return "YES"
}

func checkCityPermit(distributor string, location string) string {
	locationTokens := strings.Split(location, "-")
	countryPermit := checkCountryPermit(distributor, locationTokens[2])
	if countryPermit == "NO" {
		return "NO"
	}

	statePermit := checkStatePermit(distributor, locationTokens[1]+"-"+locationTokens[2])
	if statePermit == "NO" {
		return "NO"
	}

	for _, v := range distributorMapper[distributor].CityExcludes {
		if v == location {
			return "NO"
		}
	}
	for _, v := range distributorMapper[distributor].CityIncludes {
		if v == location {
			return "YES"
		}
	}

	return "YES"
}

func getDistributorPermissions(inputFile io.Reader) {
	var currentDistributor string
	var distributorHierarchy string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputText := scanner.Text()
		if inputText == "" {
			continue
		}
		if inputText[0] == 'P' {
			if currentDistributor != "" {
				setIncludeInformation(currentDistributor, distributorHierarchy)
			}
			currentDistributor = initializeDistributor(inputText)
			distributorHierarchy = inputText
		}
		if inputText[0] == 'I' {
			parseIncludeTokens(currentDistributor, inputText)
		}
		if inputText[0] == 'E' {
			parseExcludeTokens(currentDistributor, inputText)
		}
	}
	setIncludeInformation(currentDistributor, distributorHierarchy)
}

func initializeDistributor(distributorText string) string {
	distributors := strings.Replace(distributorText, "Permissions: ", "", 1)
	distributors = strings.Replace(distributors, " ", "", 4)
	distributorTokens := strings.Split(distributors, "<")
	distributor := new(Distributor)
	if len(distributorTokens) == 1 {
		return distributorTokens[0]
	}
	for i := 1; i < len(distributorTokens); i++ {
		distributor.CountryExcludes = append(distributor.CountryExcludes, distributorMapper[distributorTokens[i]].CountryExcludes...)
		distributor.StateExcludes = append(distributor.StateExcludes, distributorMapper[distributorTokens[i]].StateExcludes...)
		distributor.CityExcludes = append(distributor.CityExcludes, distributorMapper[distributorTokens[i]].CityExcludes...)
	}
	distributorMapper[distributorTokens[0]] = *distributor
	return distributorTokens[0]
}

func setIncludeInformation(distributor string, distributorHierarchy string) {
	if len(distributorMapper[distributor].CountryIncludes) > 0 || len(distributorMapper[distributor].StateIncludes) > 0 || len(distributorMapper[distributor].CityIncludes) > 0 {
		return
	}
	distributors := strings.Replace(distributorHierarchy, "Permissions: ", "", 1)
	distributors = strings.Replace(distributors, " ", "", 4)
	distributorTokens := strings.Split(distributors, "<")
	if len(distributorTokens) == 1 {
		tempDistributor := cloneDistributor(distributorMapper[distributor])
		tempDistributor.CountryIncludes = append(tempDistributor.CountryIncludes, "All")
		distributorMapper[distributor] = *tempDistributor
		return
	}

	for i := 1; i < len(distributorTokens); i++ {
		if len(distributorMapper[distributorTokens[i]].CountryIncludes) > 0 || len(distributorMapper[distributorTokens[i]].StateIncludes) > 0 || len(distributorMapper[distributorTokens[i]].CityIncludes) > 0 {
			tempDistributor := cloneDistributor(distributorMapper[distributor])
			tempDistributor.CountryIncludes = append(tempDistributor.CountryIncludes, distributorMapper[distributorTokens[i]].CountryIncludes...)
			tempDistributor.StateIncludes = append(tempDistributor.StateIncludes, distributorMapper[distributorTokens[i]].StateIncludes...)
			tempDistributor.CityExcludes = append(tempDistributor.CityIncludes, distributorMapper[distributorTokens[i]].CityIncludes...)
			distributorMapper[distributor] = *tempDistributor
			break
		}
	}
}

func parseIncludeTokens(distributor string, description string) {
	description = strings.Replace(description, "INCLUDE: ", "", 1)
	descriptionTokens := strings.Split(description, "-")
	tempDistributor := cloneDistributor(distributorMapper[distributor])
	if len(descriptionTokens) == 3 {
		tempDistributor.CityIncludes = append(tempDistributor.CityIncludes, description)
	}
	if len(descriptionTokens) == 2 {
		tempDistributor.StateIncludes = append(tempDistributor.StateIncludes, description)
	}
	if len(descriptionTokens) == 1 {
		tempDistributor.CountryIncludes = append(tempDistributor.CountryIncludes, description)
	}
	distributorMapper[distributor] = *tempDistributor
}

func parseExcludeTokens(distributor string, description string) {
	description = strings.Replace(description, "EXCLUDE: ", "", 1)
	descriptionTokens := strings.Split(description, "-")
	tempDistributor := cloneDistributor(distributorMapper[distributor])
	if len(descriptionTokens) == 3 {
		tempDistributor.CityExcludes = append(tempDistributor.CityExcludes, description)
	}
	if len(descriptionTokens) == 2 {
		tempDistributor.StateExcludes = append(tempDistributor.StateExcludes, description)
	}
	if len(descriptionTokens) == 1 {
		tempDistributor.CountryExcludes = append(tempDistributor.CountryExcludes, description)
	}
	distributorMapper[distributor] = *tempDistributor
}

func cloneDistributor(distributor Distributor) *Distributor {
	newDistributor := new(Distributor)
	newDistributor.CityExcludes = distributor.CityExcludes
	newDistributor.CityIncludes = distributor.CityIncludes
	newDistributor.CountryExcludes = distributor.CountryExcludes
	newDistributor.CountryIncludes = distributor.CountryIncludes
	newDistributor.StateExcludes = distributor.StateExcludes
	newDistributor.StateIncludes = distributor.StateIncludes
	return newDistributor
}
