package distributer

import (
	"log"
	"strconv"
	"testing"


	"github.com/eramit9006/challenge2016/csvreader"
	"github.com/eramit9006/challenge2016/models"
)

func TestStaticInput(t *testing.T) {
	csvFileName := "../cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)
	csvreader.MakeDataStore(csvFileName, countryStateMap)

	input := models.InputModel{
		Name:       UpperCaseNoSpace("distributer"),
		Permission: "India",
		AuthType:   models.Include,
	}

	_, err := AddDistributer(input, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	input1 := models.InputModel{
		Name:       UpperCaseNoSpace("distributer"),
		Permission: "Tamil Nadu-India",
		AuthType:   models.Exclude,
	}

	_, err = AddDistributer(input1, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	input2 := models.InputModel{
		Name:       UpperCaseNoSpace("distributer1 < distributer"),
		Permission: "Keelakarai-Tamil Nadu-India",
		AuthType:   models.Include,
	}

	_, err = AddDistributer(input2, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}
}

func TestAddDistributer(t *testing.T) {

	testScenrio := []struct {
		input          models.InputModel
		expectedResult string
	}{
		{
			input: models.InputModel{
				Name:       UpperCaseNoSpace("distributer"),
				Permission: "India",
				AuthType:   models.Include,
			},
			expectedResult: "sucess",
		}, {
			input: models.InputModel{
				Name:       UpperCaseNoSpace("distributer"),
				Permission: "Tamil Nadu-India",
				AuthType:   models.Exclude,
			},
			expectedResult: "sucess",
		}, {
			input: models.InputModel{
				Name:       UpperCaseNoSpace("distributer1 < distributer"),
				Permission: "Keelakarai-Tamil Nadu-India",
				AuthType:   models.Include,
			},
			expectedResult: "Parent distributer dont have access to grant permission- Keelakarai-Tamil Nadu-India",
		},
	}
	csvFileName := "../cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)

	csvreader.MakeDataStore(csvFileName, countryStateMap)

	for i, scenrio := range testScenrio {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			result, _ := AddDistributer(scenrio.input, countryStateMap, distributerMap)
			//3.1expects error message
			if want, got := scenrio.expectedResult, result; want != got {
				t.Errorf("expected result %#v, but got %#v", want, got)
				return
			} else if result != "" {
				return //expected error; done here
			}

		})
	}
}
