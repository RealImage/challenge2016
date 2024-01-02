package controller

import (
	"errors"
	"example/models"
	"strings"
)

func AddDistributer(input models.InputModel, countryMap models.CountryMap, distributerMap models.DistributerMap) (string, error) {
	var resultErr error
	result := "sucess"
	if validity(input.Permission, countryMap) {
		countPC := strings.Count(input.Name, "<")
		pcArr := strings.Split(input.Name, "<")

		switch countPC {
		case 0:
			distributer := models.Distributer{
				Name:            input.Name,
				InputPermission: input.Permission,
				AuthType:        input.AuthType,
			}
			resultErr = addAnyDist(distributer, distributerMap, 0, nil)

		default:

			distributer := models.Distributer{
				Name:            pcArr[0],
				InputPermission: input.Permission,
				AuthType:        input.AuthType,
			}
			parentPermission := distributerMap[pcArr[1]]
			result, err := checkPermission(distributer, parentPermission, countryMap)
			if err != nil {
				resultErr = err
			} else if result {
				resultErr = addAnyDist(distributer, distributerMap, 1, pcArr)
			}

		}
	} else {
		resultErr = errors.New("Invalid Input permission...." + input.Permission)
	}
	if resultErr != nil {
		result = resultErr.Error()
	}
	return result, resultErr
}
func directIncCheck(InputPermission string, parentPermission models.Permission) bool {
	_, checkInclude := parentPermission.IncludeMap[InputPermission]
	return checkInclude
}
func directExcCheck(InputPermission string, parentPermission models.Permission) bool {
	_, checkExclude := parentPermission.ExcludeMap[InputPermission]
	return checkExclude
}
func validity(InputPermission string, countryStateMap models.CountryMap) bool {
	plist, ptype := getPermissionType(InputPermission)
	var isvalid bool
	switch ptype {
	case models.CountryType:
		_, isvalid = countryStateMap[InputPermission]
	case models.StateType:
		stmap := countryStateMap[plist[1]]
		_, isvalid = stmap[plist[0]]
	case models.CityType:
		stmap := countryStateMap[plist[2]]
		ctmap := stmap[plist[1]]
		_, isvalid = ctmap[plist[0]]
	}
	return isvalid

}
func checkByPermissionTypeInc(InputPermission string, parentPermission models.Permission, countryStateMap models.CountryMap) (bool, error) {
	var result bool
	var resultErr error
	plist, ptype := getPermissionType(InputPermission)
	switch ptype {
	case models.CountryType:
		if _, ccheck := countryStateMap[plist[0]]; ccheck {
			result = true
		} else {
			resultErr = errors.New("invalid input country " + InputPermission)
		}
	case models.StateType:
		if _, childCountrycheck := parentPermission.IncludeMap[plist[1]]; childCountrycheck {
			smap := countryStateMap[plist[1]]
			if _, scheck := smap[plist[0]]; scheck {
				result = true
			} else {
				resultErr = errors.New("invalid input state - " + InputPermission)
			}
		}

	case models.CityType:
		if _, childStateCheck := parentPermission.IncludeMap[plist[1]+"-"+plist[2]]; childStateCheck {
			smap := countryStateMap[plist[2]]
			if _, scheck := smap[plist[1]]; scheck {
				result = true
			} else {
				resultErr = errors.New("invalid input city l1 - " + InputPermission)
			}
		} else if _, childCountryCheck := parentPermission.IncludeMap[plist[2]]; childCountryCheck {
			smap := countryStateMap[plist[2]]
			if citymap, scheck := smap[plist[1]]; scheck {
				if _, cityCheck := citymap[plist[0]]; cityCheck {
					result = true
				} else {
					resultErr = errors.New("invalid input city l3 - " + InputPermission)
				}

			} else {
				resultErr = errors.New("invalid input city l2 - " + InputPermission)
			}
		}

	}
	return result, resultErr
}
func checkByPermissionTypeExc(InputPermission string, parentPermission models.Permission, countryStateMap models.CountryMap) (bool, error) {
	var result bool
	var resultErr error
	plist, ptype := getPermissionType(InputPermission)
	switch ptype {
	case models.CountryType:
		if _, ccheck := countryStateMap[plist[0]]; ccheck {
			result = true
		} else {
			resultErr = errors.New("invalid input country" + InputPermission)
		}
	case models.StateType:
		if _, childCountrycheck := parentPermission.ExcludeMap[plist[1]]; childCountrycheck {
			smap := countryStateMap[plist[1]]
			if _, scheck := smap[plist[0]]; scheck {
				result = true
			} else {
				resultErr = errors.New("invalid input state -" + InputPermission)
			}
		}

	case models.CityType:
		if _, childStateCheck := parentPermission.ExcludeMap[plist[1]+"-"+plist[2]]; childStateCheck {
			smap := countryStateMap[plist[2]]
			if _, scheck := smap[plist[1]]; scheck {
				result = true
			} else {
				resultErr = errors.New("invalid input city Ex l1 -  " + InputPermission)
			}
		} else if _, childCountryCheck := parentPermission.ExcludeMap[plist[2]]; childCountryCheck {
			smap := countryStateMap[plist[2]]
			if citymap, scheck := smap[plist[1]]; scheck {
				if _, cityCheck := citymap[plist[0]]; cityCheck {
					result = true
				} else {
					resultErr = errors.New("invalid input city Ex l3 -" + InputPermission)
				}

			} else {
				resultErr = errors.New("invalid input city Ex l2 - " + InputPermission)
			}
		}

	}
	return result, resultErr
}
func checkPermission(child models.Distributer, parentPermission models.Permission, countryStateMap models.CountryMap) (bool, error) {
	var result bool
	var resultErr error
	var directCheck bool
	if child.AuthType == models.Include {
		directCheck = directIncCheck(child.InputPermission, parentPermission)

		if directCheck {
			result = true
		} else {
			resultExc, excErr := checkByPermissionTypeExc(child.InputPermission, parentPermission, countryStateMap)
			if excErr != nil {
				resultErr = excErr
			} else if resultExc {
				result = false
				resultErr = errors.New("Parent distributer dont have access to grant permission- " + child.InputPermission)
			}
			resultInc, incErr := checkByPermissionTypeInc(child.InputPermission, parentPermission, countryStateMap)
			if incErr != nil {
				resultErr = incErr
			} else if resultInc {
				result = true
			}
		}
	} else {
		directCheck = directExcCheck(child.InputPermission, parentPermission)

		if directCheck {
			result = true
		} else {
			resultExc, excErr := checkByPermissionTypeExc(child.InputPermission, parentPermission, countryStateMap)
			if excErr != nil {
				resultErr = excErr
			} else if resultExc {
				result = true
			}
			resultInc, incErr := checkByPermissionTypeInc(child.InputPermission, parentPermission, countryStateMap)
			if incErr != nil {
				resultErr = incErr
			} else if resultInc {
				result = true
			}
		}
	}

	return result, resultErr
}

func addAnyDist(distributer models.Distributer, distributerMap models.DistributerMap, count int, pcAr []string) error {
	var resultErr error
	permission, ok := distributerMap[distributer.Name]
	_, pType := getPermissionType(distributer.InputPermission)

	if ok {

		if distributer.AuthType == models.Include {
			_, iok := permission.IncludeMap[distributer.InputPermission]
			if iok {
				resultErr = errors.New("permission Exist.. in include - " + distributer.InputPermission)
			} else {
				permission.IncludeMap[distributer.InputPermission] = pType
			}
		} else {
			_, eok := permission.ExcludeMap[distributer.InputPermission]
			if eok {
				resultErr = errors.New("permission Exist.. in exclude - " + distributer.InputPermission)
			} else {
				permission.ExcludeMap[distributer.InputPermission] = pType
			}
		}

	} else {
		var permission models.Permission
		if distributer.AuthType == models.Include {
			permission = models.Permission{
				IncludeMap: map[string]models.PermissionType{distributer.InputPermission: pType},
				ExcludeMap: map[string]models.PermissionType{},
			}
		} else {
			permission = models.Permission{
				IncludeMap: map[string]models.PermissionType{},
				ExcludeMap: map[string]models.PermissionType{distributer.InputPermission: pType},
			}
		}
		for k, v := range pcAr {
			if k == 0 {
				continue
			}
			if count == 1 {
				ss := distributerMap[v]
				valInc := ss.IncludeMap
				for k := range valInc {
					if _, ok := permission.IncludeMap[k]; !ok {
						permission.IncludeMap[k] = valInc[k]
					}

				}
				valExc := ss.ExcludeMap
				for k := range valExc {
					if _, ok := permission.ExcludeMap[k]; !ok {
						permission.ExcludeMap[k] = valExc[k]
					}
				}

			}
		}
		distributerMap[distributer.Name] = permission
	}
	return resultErr
}

func getPermissionType(permission string) ([]string, models.PermissionType) {
	decideNum := strings.Count(permission, "-")
	var ptype models.PermissionType
	switch decideNum {
	case 0:
		ptype = models.CountryType
	case 1:
		ptype = models.StateType
	case 2:
		ptype = models.CityType
	}
	pArr := strings.Split(permission, "-")
	return pArr, ptype
}

func UpperCaseNoSpace(input string) string {
	return strings.ToUpper(strings.Replace(input, " ", "", -1))
}
