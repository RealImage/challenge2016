package distributor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type DistributorFixture struct {
	*gunit.Fixture
}

func TestDistributorFixture(t *testing.T) {
	gunit.RunSequential(new(DistributorFixture), t)
}

func (fixture *DistributorFixture) TestCreatingDistributor() {
	distributorsData := getDistributorsData()
	json, _ := json.Marshal(distributorsData)
	req, _ := http.NewRequest("POST", "/createDistributor", bytes.NewBuffer(json))
	w := httptest.NewRecorder()
	CreateDistributor(w, req)

	output := w.Body.String()

	fixture.So(output, should.NotBeBlank)
	fixture.AssertEqual(http.StatusOK, w.Code)
}

func (fixture *DistributorFixture) TestCreatingDistributorWithDuplicateData() {
	distributorsData := getDuplicateDistributorsData()
	json, _ := json.Marshal(distributorsData)
	req, _ := http.NewRequest("POST", "/createDistributor", bytes.NewBuffer(json))
	w := httptest.NewRecorder()
	CreateDistributor(w, req)

	output := w.Body.String()

	expectedError1 := `Unable to create distributor, Distributor1 already exist`
	expectedError2 := `Distributor4's parent Distributor5 doesn't exist`

	fixture.So(output, should.ContainSubstring, expectedError1)
	fixture.So(output, should.ContainSubstring, expectedError2)
	fixture.AssertEqual(http.StatusInternalServerError, w.Code)
}

func (fixture *DistributorFixture) TestCreatingDistributorError() {
	distributorsData := `{"%test%": "data"}`
	json, _ := json.Marshal(distributorsData)
	req, _ := http.NewRequest("POST", "/createDistributor", bytes.NewBuffer(json))
	w := httptest.NewRecorder()
	CreateDistributor(w, req)
	expectedOutput := `json: cannot unmarshal string into Go value of type []distributor.Distributor
`
	output := w.Body.String()

	fixture.AssertEqual(expectedOutput, output)
	fixture.AssertEqual(http.StatusInternalServerError, w.Code)
}

func (fixture *DistributorFixture) TestVerifyingDistributorRights() {
	distributionsData := getDistributionLocationsData()
	json, _ := json.Marshal(distributionsData)
	req, _ := http.NewRequest("POST", "/verifyDistribution", bytes.NewBuffer(json))
	w := httptest.NewRecorder()
	VerifyDistributorRights(w, req)

	output := w.Body.String()

	fixture.So(output, should.NotBeBlank)
}

///////////////////////////////////////// Fake Data ////////////////////////////////////////////

func getDistributorsData() []Distributor {
	fakeDistributorsData := make([]Distributor, 3)

	fakeDistributorsData[0] = getDistributor1Data()
	fakeDistributorsData[1] = getDistributor2Data()
	fakeDistributorsData[2] = getDistributor3Data()

	return fakeDistributorsData
}

func getDuplicateDistributorsData() []Distributor {
	fakeDistributorsData := make([]Distributor, 2)

	// adding duplicate data
	fakeDistributorsData[0] = getDistributor1Data()
	// parentDistributor doesn't exist
	fakeDistributorsData[1] = getDistributor4Data()

	return fakeDistributorsData
}

func getDistributor1Data() Distributor {
	distInclLocations := make([]Location, 2)
	distExclLocations := make([]Location, 2)
	distInclLocations[0] = getLocation("IN", "", "")
	distInclLocations[1] = getLocation("US", "", "")
	distExclLocations[0] = getLocation("IN", "KA", "")
	distExclLocations[1] = getLocation("IN", "TN", "CENAI")

	distributorData := Distributor{
		Name:              "Distributor1",
		ParentDistributor: "none",
		IncludedLocations: distInclLocations,
		ExcludedLocations: distExclLocations,
	}

	return distributorData
}

func getDistributor2Data() Distributor {
	distInclLocations := make([]Location, 1)
	distExclLocations := make([]Location, 2)
	distInclLocations[0] = getLocation("IN", "", "")
	distExclLocations[0] = getLocation("IN", "TN", "")
	distExclLocations[1] = getLocation("IN", "GJ", "")

	distributorData := Distributor{
		Name:              "Distributor2",
		ParentDistributor: "Distributor1",
		IncludedLocations: distInclLocations,
		ExcludedLocations: distExclLocations,
	}

	return distributorData
}

func getDistributor3Data() Distributor {
	distInclLocations := make([]Location, 4)
	distInclLocations[0] = getLocation("IN", "TN", "CENAI")
	distInclLocations[1] = getLocation("IN", "KA", "HBALI")
	distInclLocations[2] = getLocation("IN", "KL", "")
	distInclLocations[3] = getLocation("IN", "GJ", "NAVSR")

	distributorData := Distributor{
		Name:              "Distributor3",
		ParentDistributor: "Distributor2",
		IncludedLocations: distInclLocations,
	}

	return distributorData
}

// Parent Distributor (Distributor5) not available
func getDistributor4Data() Distributor {
	distInclLocations := make([]Location, 4)
	distInclLocations[0] = getLocation("IN", "TN", "CENAI")
	distInclLocations[1] = getLocation("IN", "KA", "HBALI")
	distInclLocations[2] = getLocation("IN", "KL", "")
	distInclLocations[3] = getLocation("IN", "GJ", "NAVSR")

	distributorData := Distributor{
		Name:              "Distributor4",
		ParentDistributor: "Distributor5",
		IncludedLocations: distInclLocations,
	}

	return distributorData
}

func getLocation(countryCode, stateCode, cityCode string) Location {
	return Location{
		CityCode:    cityCode,
		StateCode:   stateCode,
		CountryCode: countryCode,
	}
}

func getDistributionLocationsData() []DistributorPermissions {
	fakeDistributionLocationData := make([]DistributorPermissions, 3)

	fakeDistributionLocationData[0] = getDistributor1VerificationData()
	fakeDistributionLocationData[1] = getDistributor2VerificationData()
	fakeDistributionLocationData[2] = getDistributor3VerificationData()

	return fakeDistributionLocationData
}

func getDistributor1VerificationData() DistributorPermissions {
	distLocations := make([]Location, 5)

	distLocations[0] = getLocation("US", "CA", "AVEJO")
	distLocations[1] = getLocation("IN", "TN", "CENAI")
	distLocations[2] = getLocation("US", "CA", "FILMR")
	distLocations[3] = getLocation("IN", "KA", "MYSUR")
	distLocations[4] = getLocation("IN", "TN", "NMAKL")

	distributorVerification := DistributorPermissions{
		Name:     "Distributor1",
		Location: distLocations,
	}

	return distributorVerification
}

func getDistributor2VerificationData() DistributorPermissions {
	distLocations := make([]Location, 6)

	distLocations[0] = getLocation("US", "CA", "")
	distLocations[1] = getLocation("IN", "KA", "NGAML")
	distLocations[2] = getLocation("US", "CA", "FILMR")
	distLocations[3] = getLocation("IN", "KA", "MYSUR")
	distLocations[4] = getLocation("IN", "TN", "NMAKL")
	distLocations[5] = getLocation("IN", "TN", "SALEM")

	distributorVerification := DistributorPermissions{
		Name:     "Distributor2",
		Location: distLocations,
	}

	return distributorVerification
}

func getDistributor3VerificationData() DistributorPermissions {
	distLocations := make([]Location, 7)

	distLocations[0] = getLocation("US", "CA", "")
	distLocations[1] = getLocation("IN", "KL", "")
	distLocations[2] = getLocation("US", "GJ", "NAVSR")
	distLocations[3] = getLocation("IN", "KA", "MYSUR")
	distLocations[4] = getLocation("IN", "KA", "NGAML")
	distLocations[5] = getLocation("IN", "TN", "NMAKL")
	distLocations[6] = getLocation("IN", "TN", "SALEM")

	distributorVerification := DistributorPermissions{
		Name:     "Distributor3",
		Location: distLocations,
	}

	return distributorVerification
}
