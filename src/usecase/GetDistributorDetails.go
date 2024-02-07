package usecase

import (
	"challenge.com/domain"
	"fmt"
	"sync"
)

func GetDistributorDetails(name string, mx *sync.RWMutex) {
	mx.RLock()
	distributor, ok := domain.DistributorMap[name]
	mx.RUnlock()
	if ok {
		fmt.Println("Name: ", distributor.Name)
		fmt.Println("Include permissions: ", distributor.IncludePermissions)
		fmt.Println("Exclude Permissions: ", distributor.ExcludePermissions)

		if distributor.ParentDistributor != nil {
			fmt.Println("Parent Distributor Name: ", distributor.ParentDistributor.Name)
			fmt.Println("Parent Include permissions: ", distributor.ParentDistributor.IncludePermissions)
			fmt.Println("Parent Exclude Permissions: ", distributor.ParentDistributor.ExcludePermissions)

		}
	} else {
		fmt.Println("Distributor not exist")
	}
	fmt.Println()
	fmt.Println()
	return
}
