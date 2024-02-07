package usecase

import (
	"challenge.com/domain"
	"fmt"
	"sync"
)

func DeleteDistributorDetails(name string, mx *sync.RWMutex) {
	_, ok := domain.DistributorMap[name]
	if ok {
		mx.Lock()
		delete(domain.DistributorMap, name)
		mx.Unlock()
		fmt.Println("Distributor " + name + " deleted successfully")
		return
	} else {
		fmt.Println("Distributor not exist for delete")
	}

	fmt.Println()
	fmt.Println()

	return
}
