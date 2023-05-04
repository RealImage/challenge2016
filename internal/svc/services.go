package svc

import (
	"distribution-mgmnt/app"
	"distribution-mgmnt/pkg/cmaps"
	"strings"
)

const (
	NO  = "NO"
	YES = "YES"
)

func SaveAddDistributor(md *app.DistributorDetails) bool {
	if !strings.EqualFold(md.ParentDistributor, "") {
		some := app.DistributorDetails{
			Exclude: make([]app.Location, 0),
			Include: make([]app.Location, 0),
		}
		some, _ = AddAllPermissions(md.ParentDistributor, true, some)
		if len(some.Include) == 0 || !compareIncludedCntryOfParent(some.Include, md.Include, some.Exclude) {
			return false
		}
		if !compareExclude(some.Exclude, md.Exclude, some.Include, md.Include) {
			return false
		}
	}

	cmaps.DistributorMgmntDB.SetDistributionDB(md)
	return true
}

func GetPermissionsByName(distributorName string) (app.DistributorDetails, bool) {
	return cmaps.DistributorMgmntDB.GetDistributorDetails(distributorName)
}
func CheckPermissions(req app.CheckPermissionJSONBody) string {
	loc, ok := cmaps.DistributorMgmntDB.GetDistributorDetails(req.Name)
	if !ok || strings.EqualFold(req.Location.Country, "") {
		return NO
	}
	if !strings.EqualFold(loc.ParentDistributor, "") {
		loc2, ok2 := cmaps.DistributorMgmntDB.GetDistributorDetails(loc.ParentDistributor)
		if !ok2 {
			return NO
		}

		loc2.Exclude = append(loc2.Exclude, loc.Exclude...)
		loc2.Include = append(loc2.Include, loc.Include...)
		return CheckPermission(req, loc2)

	}

	return CheckPermission(req, loc)
}
func CheckPermission(req app.CheckPermissionJSONBody, loc app.DistributorDetails) string {
	// loc, ok := cmaps.DistributorMgmntDB.GetDistributorDetails(req.Name)
	// if !ok || strings.EqualFold(req.Location.Country, "") {
	// 	return NO
	// }
	prv, isCityAvlbl := checkCity(req.Location.City)
	_, isPrvnceAvlbl := checkProvince(req.Location.Province)
	cntry, isCntryAvlbl := checkCntry(req.Location.Country)
	if (!strings.EqualFold(req.Location.City, "") && !isCityAvlbl) || (!strings.EqualFold(req.Location.Province, "") && !isPrvnceAvlbl) || (!strings.EqualFold(req.Location.Country, "") && !isCntryAvlbl) {
		return NO
	}

	if strings.EqualFold(req.Location.City, "") && !strings.EqualFold(req.Location.Province, "") && !strings.EqualFold(req.Location.Country, "") {
		if !isPresentInArrayStruct(loc.Include, req.Location.Country) || checkExcludeList(loc.Exclude, req.Location.Province, "") {
			return NO
		} else if _, ok := cntry[req.Location.Province]; ok {
			return YES
		}
	} else if !strings.EqualFold(req.Location.City, "") && !strings.EqualFold(req.Location.Province, "") && !strings.EqualFold(req.Location.Country, "") {
		if checkExcludeList(loc.Exclude, req.Location.Province, req.Location.City) || !isPresentInArrayStruct(loc.Include, req.Location.Country) {
			return NO
		} else if isPresent(cntry[req.Location.Province], req.Location.City) {
			return YES
		}
	} else if !strings.EqualFold(req.Location.City, "") && strings.EqualFold(req.Location.Province, "") && !strings.EqualFold(req.Location.Country, "") {
		if !isPresentInArrayStruct(loc.Include, req.Location.Country) || checkExcludeList(loc.Exclude, req.Location.Province, req.Location.City) {
			return NO
		} else if isPresent(cntry[prv], req.Location.City) {
			return YES
		}
	} else if strings.EqualFold(req.Location.City, "") && strings.EqualFold(req.Location.Province, "") && !strings.EqualFold(req.Location.Country, "") {
		if !isPresentInArrayStruct(loc.Include, req.Location.Country) {
			return NO
		} else {
			return YES
		}
	}

	return NO
}

func checkIncludeList(include []string, cntry string) bool {
	count := 0
	for _, val := range include {
		if val == cntry {
			count++
		}
	}
	return count != 0
}

func checkExcludeList(exclude []app.Location, province, city string) bool {
	for _, val := range exclude {
		if !strings.EqualFold(province, "") && province == val.Province && (strings.EqualFold(city, "") || strings.EqualFold(val.City, "")) {
			return true
		}
		if !strings.EqualFold(city, "") && city == val.City && (strings.EqualFold(province, "") || strings.EqualFold(val.Province, "")) {
			return true
		}
		if !strings.EqualFold(province, "") && !strings.EqualFold(city, "") && province == val.Province && city == val.City {
			return true
		}
	}
	return false
}

func isPresent(sl []string, value string) bool {
	for _, val := range sl {
		if val == value {
			return true
		}
	}
	return false
}
func isPresentInArrayStruct(arr []app.Location, str string) bool {
	for _, val := range arr {
		if strings.EqualFold(val.Country, str) {
			return true
		}
		if strings.EqualFold(val.Province, str) {
			return true
		}
		if strings.EqualFold(val.City, str) {
			return true
		}
	}
	return false

}

func checkCity(city string) (string, bool) {
	res, ok := cmaps.DistributorMgmntDB.GetProvinceFromCityMap(city)
	return res, ok
}

func checkProvince(province string) (string, bool) {
	res, ok := cmaps.DistributorMgmntDB.GetCountryFromProvinceMap(province)
	return res, ok
}

func checkCntry(cntry string) (map[string][]string, bool) {
	res, ok := cmaps.DistributorMgmntDB.GetLocationFromCountryMap(cntry)
	return res, ok
}

func compareIncludedCntryOfParent(parent, child, ex []app.Location) bool {
	for i := 0; i < len(child); i++ {
		count := 0
		for j := 0; j < len(parent); j++ {
			if !strings.EqualFold(child[i].City, "") {
				prv1, _ := checkCity(child[i].City)
				cntry1, _ := checkProvince(prv1)
				if !strings.EqualFold(parent[i].Province, "") {
					if strings.EqualFold(prv1, parent[i].Province) && !isPresentInArrayStruct(ex, child[i].Province) {
						count++
					}
				}
				if !strings.EqualFold(parent[i].Country, "") {
					if strings.EqualFold(cntry1, parent[i].Country) && !isPresentInArrayStruct(ex, child[i].City) {
						count++
					}
				}
			}

			if !strings.EqualFold(child[i].Province, "") {
				cntry1, _ := checkProvince(child[i].Province)
				if !strings.EqualFold(parent[i].Country, "") {
					if strings.EqualFold(cntry1, parent[i].Country) && !isPresentInArrayStruct(ex, child[i].Province) {
						count++
					}
				}
			}
			if !strings.EqualFold(child[i].Country, "") {
				if strings.EqualFold(child[i].Country, parent[i].Country) && strings.EqualFold(child[i].Province, "") && strings.EqualFold(child[i].City, "") {
					count++
				}
			}
			if count != 0 {
				break
			}
		}
		if count != 0 {
			return true
		}
	}
	return false
}

func compareExclude(prnt []app.Location, child, parentIn, childIn []app.Location) bool {
	for i := 0; i < len(child); i++ {

		if !strings.EqualFold(child[i].Province, "") {

			cntry, ok := checkProvince(child[i].Province)
			if !ok || !checkExcludeList(prnt, child[i].Province, "") && (!isPresentInArrayStruct(childIn, cntry) || !isPresentInArrayStruct(parentIn, cntry)) || !strings.EqualFold(cntry, child[i].Country) {
				return false
			}

		}
		if !strings.EqualFold(child[i].City, "") {
			prv, _ := checkCity(child[i].City)
			cntry, ok := checkProvince(prv)
			if !ok || !checkExcludeList(prnt, "", child[i].City) && (!isPresentInArrayStruct(childIn, cntry) || !isPresentInArrayStruct(parentIn, cntry)) || !strings.EqualFold(cntry, child[i].Country) || !strings.EqualFold(cntry, child[i].Country) {
				return false
			}
		}

	}
	return true
}

func AddAllPermissions(ParentDistributor string, isTrue bool, res app.DistributorDetails) (app.DistributorDetails, bool) {
	if !isTrue {
		return res, isTrue
	}

	temp, tr := cmaps.DistributorMgmntDB.GetDistributorDetails(ParentDistributor)
	if tr {
		res.Exclude = append(res.Exclude, temp.Exclude...)
		res.Include = append(res.Include, temp.Include...)
		res.ParentDistributor = temp.ParentDistributor

	}
	isTrue = tr
	return AddAllPermissions(res.ParentDistributor, isTrue, res)
}
