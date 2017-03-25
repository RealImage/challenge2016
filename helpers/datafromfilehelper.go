package helpers

import (
	"os"
	"bufio"
	"log"
	"strings"
)

func DataFromFile(fileLocation string) ([][]string, error) {

	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Println(err)
	}

	var mainSlice [][]string
	for i := 0; i < len(lines); i++ {
		delimiterCount := strings.Count(lines[i], ",")
		var subSlice []string
		var subString string
		for j := 0; j < delimiterCount; j++ {
			endingIndex:=strings.Index(lines[i], ",")
			subString = lines[i][0 : endingIndex]
			subSlice = append(subSlice, subString)
			lines[i] = lines[i][(endingIndex+1):len(lines[i])]
			if strings.Count(lines[i], ",") == 0 {
				subSlice = append(subSlice, lines[i])
			}
		}
		if i != 0 {
			mainSlice = append(mainSlice, subSlice)
		}
	}


	return mainSlice, scanner.Err()
}
