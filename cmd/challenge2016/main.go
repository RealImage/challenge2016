package main

import (
	"bufio"
	"challenge2016/pkg/model"
	"challenge2016/utils"
	"encoding/csv"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

/*
	*** Program Flow ***

- The distribution data would be given in csv. Load and parse the csv
- After Parsing, we have the entire distribution information loaded into a map
- Distribution Info map has Key: distributor's Name, Value: distributorInfo struct
- Input prompt is opened for the user for validation

- USAGE: From the repo base path, run the following commands
- go build -C cmd/challenge2016/
- ./cmd/challenge2016/challenge2016
*/

var log *slog.Logger

func init() {
	log = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func main() {
	log.Info("Welcome to challenge 2016!")
	defer func() {
		log.Info("Challenge completed!")
	}()

	/* Load the csv data file into the program */
	repoBasePath, _ := os.Getwd()
	inputDataPath := filepath.Join(repoBasePath, "data/input_data.csv")
	log.Info("Input data path", "path", inputDataPath)

	fileP, err := os.OpenFile(inputDataPath, os.O_RDONLY, 0400)
	if err != nil {
		log.Error("failed to load the csv distribution data", "error", err)
		return
	}
	defer fileP.Close()

	csvReader := csv.NewReader(fileP)
	inputCSVData, err := csvReader.ReadAll()
	if err != nil {
		log.Error("failed to read the distrubtion data", "error", err)
		return
	}

	// Parse the input csv data and get the distribution info map
	distributionInfo, err := utils.ParseCsv(log, inputCSVData)
	if err != nil {
		log.Error("failed to Parse the distribution data csv", "error", err)
		return
	}
	log.Info("Successfully parsed the distribution data ")
	// printDistributionInfo(log, distributionInfo)

	log.Info("Opening prompt for user validation!")
	for {
		reader := bufio.NewReader(os.Stdin)

		var inputData []string

		if line, err := reader.ReadString('\n'); err != nil {
			log.Error("Failed to read user input", "error", err)
			return
		} else if line = strings.TrimSpace(line); len(line) == 0 {
			return
		} else {
			inputData = strings.Split(line, " ")
		}

		if len(inputData) != 2 {
			log.Info("USAGE: Enter the following details (space seperated)" +
				"Name(distributor's name), " +
				"Region(city-state-country format) as the arguments",
			)
			reader.Reset(os.Stdin)
			continue
		}

		name, region := inputData[0], inputData[1]
		if utils.ValidateInput(log, name, region, distributionInfo) {
			log.Info("Distributor AUTHORIZED for region",
				"name", name,
				"region", region,
			)
		} else {
			log.Info("Distributor is UNAUTHORIZED for region",
				"name", name,
				"region", region,
			)
		}
	}
}

func printDistributionInfo(log *slog.Logger, distributionInfo model.DistributionInfo) {
	log.Info("Printing Distribution Info...")

	for _, record := range distributionInfo {
		distributor := record.(model.Distributor)
		data, err := distributor.Marshall()
		if err != nil {
			log.Error("failed to convert distributor data to json", "error", err)
		}

		log.Info("Distributor::",
			"name", distributor.Name,
			"data", data,
		)
	}
}
