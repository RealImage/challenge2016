package usecase

import (
	"challenge.com/domain"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// CheckValidAreas Check input city is valid city and add it
func CheckValidAreas(permissions []string) []string {
	var validPermissions []string
	for _, permission := range permissions {
		tmp := strings.ToLower(permission)
		p := strings.Split(permission, "-")
		if len(p) > 0 {
			city := strings.ToLower(p[len(p)-1])
			cityData, ok := domain.CountryMap[city]
			if ok {
				for _, data := range cityData {
					dataString := strings.Join(data, "-")
					dataString = strings.ToLower(dataString)
					if strings.Contains(dataString, tmp) {
						validPermissions = append(validPermissions, permission)
						break
					}
				}
			}
		}
	}
	return validPermissions
}

func CheckUserPermissions(area string, name string, mx *sync.RWMutex) (bool, []string, []string, error) {
	area = strings.ToLower(area)
	validAreas := CheckValidAreas([]string{area})
	if len(validAreas) == 0 {
		fmt.Println("Not a valid area----")
		return false, nil, nil, errors.New("not valid area")
	}
	mx.RLock()
	distributor, ok := domain.DistributorMap[name]
	mx.RUnlock()
	if !ok {
		fmt.Println("Distributor not available ,need to create")
		return false, nil, nil, errors.New("distributor not available")
	}
	var includeFlag, excludeFlag bool
	for _, include := range distributor.IncludePermissions {
		if strings.Contains(area, include) {
			includeFlag = true
		}
	}
	for _, exclude := range distributor.ExcludePermissions {
		if strings.Contains(area, exclude) {
			excludeFlag = true
		}
	}
	if includeFlag && !excludeFlag {
		return true, distributor.IncludePermissions, distributor.ExcludePermissions, nil
	}
	return false, distributor.IncludePermissions, distributor.ExcludePermissions, nil
}

// CheckValidAreasWithParent Check input city is subset of parent distributor
func CheckValidAreasWithParent(permissions []string, ParentPermissions []string) []string {
	var validPermissions []string
	for _, permission := range permissions {
		tmp := strings.ToLower(permission)
		for _, data := range ParentPermissions {
			data = strings.ToLower(data)
			if strings.Contains(tmp, data) {
				validPermissions = append(validPermissions, permission)
				break
			}
		}

	}
	return validPermissions
}
