package usecase

import (
	"bufio"
	"challenge.com/domain"
	"errors"
	"fmt"
	"os"
	"sync"
)

func GetDistributorInput() ([]string, []string) {
	var includePermissions, excludePermissions []string
	flag2 := true
	for flag2 {
		fmt.Println("Permissions should be of format 'CHENNAI-TAMIL NADU-INDIA'  - is mandatory")
		fmt.Println("The choices are :")
		fmt.Println("1.enter include permissions for distributor")
		fmt.Println("2.enter exclude permissions for distributor")
		fmt.Println("3.exit permissions choice")
		var permissionChoice int
		fmt.Println("enter permission choice for distributor")
		_, err := fmt.Scanln(&permissionChoice)
		if err != nil {
			fmt.Println("Error in getting permissionChoice")
		}
		switch permissionChoice {
		case 1:
			fmt.Println("enter include permissions for distributor")
			var includePermission string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				fmt.Println("Error in getting included areas")
			} else {
				includePermission = scanner.Text()
				includePermissions = append(includePermissions, includePermission)
			}
		case 2:
			fmt.Println("enter exclude permissions for distributor")
			var excludePermission string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				fmt.Println("Error in getting excluded areas")
			} else {
				excludePermission = scanner.Text()
				excludePermissions = append(excludePermissions, excludePermission)
			}
		case 3:
			flag2 = false
		}

	}
	return includePermissions, excludePermissions
}
func CreateDistributor(name string, includePermissions, excludePermissions []string, mx *sync.RWMutex) error {
	includePermissions = CheckValidAreas(includePermissions)
	excludePermissions = CheckValidAreas(excludePermissions)
	if len(includePermissions) == 0 && len(excludePermissions) == 0 {
		return errors.New("enter valid cities to the distributor ")
	} else {
		_, ok := domain.DistributorMap[name]
		if !ok {
			distributor := domain.Distributor{
				Name:               name,
				IncludePermissions: includePermissions,
				ExcludePermissions: excludePermissions,
			}
			mx.Lock()
			domain.DistributorMap[name] = distributor
			mx.Unlock()
			if len(includePermissions) > 0 {
				fmt.Println("Valid Include cities added ", includePermissions, " for distributor ", name)
			}
			if len(excludePermissions) > 0 {
				fmt.Println("Valid Include cities added ", excludePermissions, " for distributor ", name)
			}
		} else {
			return errors.New("distributor with this name already created")
		}
	}
	fmt.Println()
	fmt.Println()
	return nil
}

func CreateSubDistributor(name string, parentName string, includePermissions, excludePermissions []string, mx *sync.RWMutex) error {
	ParentDistributor, ok := domain.DistributorMap[parentName]
	if !ok {
		return errors.New("parent distributor entry not available")
	}
	includePermissions = CheckValidAreasWithParent(includePermissions, ParentDistributor.IncludePermissions)
	excludePermissions = CheckValidAreasWithParent(excludePermissions, ParentDistributor.ExcludePermissions)
	if len(includePermissions) == 0 && len(excludePermissions) == 0 {
		return errors.New("enter valid cities to the sub distributor that are allowed by parent distributor ")
	} else {
		_, ok := domain.DistributorMap[name]
		if !ok {
			distributor := domain.Distributor{
				Name:               name,
				IncludePermissions: includePermissions,
				ExcludePermissions: excludePermissions,
				ParentDistributor:  &ParentDistributor,
			}
			mx.Lock()
			domain.DistributorMap[name] = distributor
			mx.Unlock()
			if len(includePermissions) > 0 {
				fmt.Println("Valid Include cities added ", includePermissions, " for distributor ", name)
			}
			if len(excludePermissions) > 0 {
				fmt.Println("Valid Include cities added ", excludePermissions, " for distributor ", name)
			}
		} else {
			return errors.New("sub distributor with this name already created")
		}
	}
	fmt.Println()
	fmt.Println()
	return nil
}
