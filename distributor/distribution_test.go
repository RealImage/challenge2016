package distribution

import (
	// "log"
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DistributorSuite struct {
	suite.Suite
	cities []Cities
}

func (ds *DistributorSuite) SetupSuite() {
	ds.cities = PrepareCitiesJson("../cities.csv")
}

func (ds *DistributorSuite) TestPrepareRootUser() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: INDIA",
		"INCLUDE: CHINA",
		"EXCLUDE: KERALA_INDIA",
		"EXCLUDE: TAMIL NADU_INDIA",
		"INCLUDE: Al-Hamdaniya_Ninawa_Iraq",
		"INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran",
		"EXCLUDE: GOOTY_ANDHRA PRADESH_INDIA",
		"INCLUDE: Basel-Landschaft_Switzerland",
	}
	expectedResult := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}
	actualResult := PrepareRootUser(permissions, ds.cities)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestPrepareRootUserTwo() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: INDIA",
		"INCLUDE: CHINA",
		"EXCLUDE: KERALA_INDIA",
		"EXCLUDE: TAMIL NADU_INDIA",
		"INCLUDE: Al-Hamdaniya_Ninawa_Iraq",
		"INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran",
		"EXCLUDE: GOOTY_ANDHRA PRADESH_INDIA",
		"INCLUDE: Basel-Landschaft_Switzerland",
		"INCLUDE: chennai_tamil nadu_india",
	}
	expectedResult := map[string]interface{}{
		"err": "[INCLUDE: chennai_tamil nadu_india] Not permitted, please try again\n",
	}
	actualResult := PrepareRootUser(permissions, ds.cities)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestPrepareSubUser() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: Al-Hamdaniya_Ninawa_Iraq",
		"INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran",
		"EXCLUDE: Maharashtra_India",
		"INCLUDE: Chhattisgarh_india",
		"INCLUDE: Fujian_China",
		"INCLUDE: Shanhaiguan_Hebei_China",
		"EXCLUDE: Yunnan_China",
		"EXCLUDE: Songmai_Sichuan_China",
	}
	root := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}

	expectedResult := map[string]interface{}{
		"excluded_states": []string{"Maharashtra_India", "Yunnan_China"},
		"excluded_cities": []string{"Songmai_Sichuan_China"},
		"included_states": []string{"Chhattisgarh_India", "Fujian_China"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran", "Shanhaiguan_Hebei_China"},
	}
	actualResult := PrepareSubUser(permissions, ds.cities, root)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestPrepareSubUserTwo() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: Umarkot_Chhattisgarh_India",
		"INCLUDE: Takhatpur_Chhattisgarh_India",
		"EXCLUDE: Zhangwan_Fujian_China",
		"EXCLUDE: Yongning_Fujian_China",
	}
	root := map[string]interface{}{
		"excluded_states": []string{"Maharashtra_India", "Yunnan_China"},
		"excluded_cities": []string{"Songmai_Sichuan_China"},
		"included_states": []string{"Chhattisgarh_India", "Fujian_China"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran", "Shanhaiguan_Hebei_China"},
	}

	expectedResult := map[string]interface{}{
		"excluded_states": []string{},
		"excluded_cities": []string{"Zhangwan_Fujian_China", "Yongning_Fujian_China"},
		"included_states": []string{},
		"included_cities": []string{"Umarkot_Chhattisgarh_India", "Takhatpur_Chhattisgarh_India"},
	}

	actualResult := PrepareSubUser(permissions, ds.cities, root)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestPrepareSubUserThree() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: Al-Hamdaniya_Ninawa_Iraq",
		"INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran",
		"Maharashtra_India",
		"INCLUDE: Chhattisgarh_india",
		"INCLUDE: Fujian_China",
		"INCLUDE: Shanhaiguan_Hebei_China",
		"EXCLUDE: Yunnan_China",
		"EXCLUDE: Songmai_Sichuan_China",
	}
	root := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}

	expectedResult := map[string]interface{}{
		"err":             "[Maharashtra_India] Not permitted - mention the INCLUDE/EXCLUDE operation\n",
		"excluded_states": []string{"Yunnan_China"},
		"excluded_cities": []string{"Songmai_Sichuan_China"},
		"included_states": []string{"Chhattisgarh_India", "Fujian_China"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran", "Shanhaiguan_Hebei_China"},
	}
	actualResult := PrepareSubUser(permissions, ds.cities, root)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestcheckUserPerm() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: chennai_tamil nadu_india",
	}

	root := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}

	expectedResult := map[string]interface{}{
		"err": "[INCLUDE: chennai_tamil nadu_india] Not permitted\n",
	}
	actualResult := checkUserPerm(permissions, ds.cities, root)
	//fmt.Printf("RESULT: %v\n", actualResult)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestcheckUserPermTwo() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: Saint-Antoine-de-Tilly_Quebec_Canada",
	}

	root := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}

	expectedResult := map[string]interface{}{}
	actualResult := checkUserPerm(permissions, ds.cities, root)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func (ds *DistributorSuite) TestcheckUserPermThree() {
	t := ds.T()

	permissions := []string{
		"INCLUDE: Basel-Landschaft_Switzerland",
	}

	root := map[string]interface{}{
		"countries":       []string{"India", "China"},
		"excluded_states": []string{"Kerala_India", "Tamil Nadu_India"},
		"excluded_cities": []string{"Gooty_Andhra Pradesh_India"},
		"included_states": []string{"Basel-Landschaft_Switzerland"},
		"included_cities": []string{"Al-Hamdaniya_Ninawa_Iraq", "Istgah-e Rah Ahan-e Garmsar_Semnan_Iran"},
	}

	expectedResult := map[string]interface{}{}
	actualResult := checkUserPerm(permissions, ds.cities, root)
	assert.Equal(t, expectedResult, actualResult, "Actual result doesn't matches")
}

func TestDistributorSuite(t *testing.T) {
	suite.Run(t, new(DistributorSuite))
}
