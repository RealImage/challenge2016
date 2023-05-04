package util

import (
	"distribution-mgmnt/app"
	"strings"
)

func ConvertSliceToUpper(slice []string) []string {
	for i, s := range slice {
		slice[i] = RemoveSpacesAndToUpper(s)
	}
	return slice
}

func ConvertSliceOfStructToUpper(Exclude []app.Location) []app.Location {
	temp := make([]app.Location, 0)
	for _, val := range Exclude {
		t := app.Location{
			City:     RemoveSpacesAndToUpper(val.City),
			Country:  RemoveSpacesAndToUpper(val.Country),
			Province: RemoveSpacesAndToUpper(val.Province),
		}
		temp = append(temp, t)
	}
	return temp
}

func RemoveSpacesAndToUpper(str string) string {
	return strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(str)), " ", "")
}
