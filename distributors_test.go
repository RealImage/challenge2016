package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/bsyed6/challenge2016/utils"
)

const (
	// BadRequestCode - Server Response Code for failure
	BadRequestCode     = 400
	SuccessRequestCode = 200
)

// TestStruct - contains fields to test
type TestStruct struct {
	RequestBody        string
	ExpectedStatusCode int
	ResponseBody       string
	ObservedStatusCode int
}

func TestAssignDistributor(t *testing.T) {
	url := "http://localhost:8000/assign_distributor"

	tests := []TestStruct{
		{utils.WITHOUT_INCLUDES_FIELD, BadRequestCode, "", 0},

		{utils.WITHOUT_EXCLUDES_FIELD, BadRequestCode, "", 0},

		{utils.WITHOUT_DISTRIBUTOR_FIELD, BadRequestCode, "", 0},

		{utils.VALID_ASSIGN_DISTRIBUTOR_REQUEST, SuccessRequestCode, "", 0},

		{utils.WRONG_COUNTRY_CODE, BadRequestCode, "", 0},
	}

	checkTestCases(url, t, tests)

	displayTestCaseResults("AssignDistributor", tests, t)

}

func TestCheckPermission(t *testing.T) {
	url := "http://localhost:8000/check_permission"

	tests := []TestStruct{
		// //Check Permissions Test
		{utils.EXCLUDED_DISTRIBUTOR, BadRequestCode, "", 0},

		{utils.EXCLUDED_CITY, BadRequestCode, "", 0},

		{utils.EXCLUDED_COUNTRY, BadRequestCode, "", 0},

		{utils.EXCLUDED_PROVINCE, BadRequestCode, "", 0},

		{utils.UNKNOWN_DISTRIBUTOR, BadRequestCode, "", 0},

		{utils.WRONG_CITY_CODE, SuccessRequestCode, "", 0},

		{utils.VALID_CHECK_PERMISSION, SuccessRequestCode, "", 0},
	}

	checkTestCases(url, t, tests)
	displayTestCaseResults("CheckPermissions", tests, t)
}

func TestAssignSubDistributor(t *testing.T) {
	url := "http://localhost:8000/assign_sub_distributor"

	tests := []TestStruct{
		//Check Permissions Test
		{utils.EXCLUDED_FROM, BadRequestCode, "", 0},

		{utils.EXCLUDED_FOR, BadRequestCode, "", 0},

		{utils.EXCLUDED_SUB_INCLUDES, BadRequestCode, "", 0},

		{utils.EXCLUDED_EXCLUDES, BadRequestCode, "", 0},
		{utils.VALID_SUB_DISTRIBUTION, SuccessRequestCode, "", 0},

		{utils.INVALID_SUB_DISTRIBUTION, SuccessRequestCode, "", 0},
	}

	checkTestCases(url, t, tests)
	displayTestCaseResults("AssignSubDistributions", tests, t)
}

func TestCheckSubPermission(t *testing.T) {
	url := "http://localhost:8000/check_permission"

	tests := []TestStruct{
		{utils.VALID_PERMISSION, SuccessRequestCode, "", 0}, // checks if "SULUR" IS PRESENT IN "D3" Distributor

		{utils.VALID_DIFF_PROVINCE1, SuccessRequestCode, "", 0}, // Valid as the country is included in distributor D1
		{utils.VALID_DIFF_PROVINCE2, SuccessRequestCode, "", 0},

		{utils.INVALID_PERMISSION1, SuccessRequestCode, "", 0}, // checks if "BOZEN" IS PRESNET IN "D#" Distributor

		{utils.INVALID_PERMISSION2, SuccessRequestCode, "", 0},
	}

	checkTestCases(url, t, tests)
	displayTestCaseResults("CheckPermissions", tests, t)
}

// CheckTestCases - To test the post request with the parameters
func checkTestCases(url string, t *testing.T, tests []TestStruct) {
	for i, testCase := range tests {
		var reader io.Reader
		reader = strings.NewReader(testCase.RequestBody)
		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err)
		}

		body, _ := ioutil.ReadAll(res.Body)

		tests[i].ResponseBody = strings.TrimSpace(string(body))
		tests[i].ObservedStatusCode = res.StatusCode
	}

}

func displayTestCaseResults(operation string, tests []TestStruct, t *testing.T) {
	t.Logf("Functionality: %s\n", operation)
	for _, test := range tests {
		if test.ObservedStatusCode == test.ExpectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n\n\n", test.RequestBody, test.ExpectedStatusCode, test.ResponseBody, test.ObservedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n\n\n", test.RequestBody, test.ExpectedStatusCode, test.ResponseBody, test.ObservedStatusCode)

		}
	}
}
