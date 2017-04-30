package helpers

import "strings"

func DelimitorSplitor(lines []string) ([][]string) {
	var mainSlice [][]string
	for i := 0; i < len(lines); i++ {
		delimiterCount := strings.Count(lines[i], ",")
		var subSlice []string
		if delimiterCount != 0 {
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
		} else {
			subSlice = append(subSlice, lines[i])
		}

		mainSlice = append(mainSlice, subSlice)
	}

	return mainSlice
}
