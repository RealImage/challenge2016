package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chrislusf/glow/flow"
)

type Permission struct {
	id      string
	parent  *Permission
	include []string
	exclude []string
}

type Address struct {
	city    string
	country string
	state   string
}

var (
	fInput         = flow.New()
	distributerMap map[string]*Permission
)

/*
This method called before main().
*/

func getAddress(key string) (a Address) {
	array := strings.Split(key, "-")
	if len(array) >= 3 {
		a.city = array[0]
		a.state = array[1]
		a.country = array[2]
	} else if len(array) >= 2 {
		a.state = array[0]
		a.country = array[1]
	} else if len(array) >= 1 {
		a.country = array[0]
	}
	return
}

func isExcluded(key string, permission *Permission) bool {
	return isInList(key, permission.exclude)
}

func isIncluded(key string, permission *Permission) bool {
	return isInList(key, permission.include)
}

func isInList(key string, list []string) bool {
	a := getAddress(key)
	for index := range list {
		//fmt.Println("method:isInList", key, list)
		listAdd := getAddress(list[index])
		if listAdd.ToCompare(a) {
			return true
		}
	}

	return false
}

func (src Address) ToCompare(dest Address) bool {
	if src.city != "" {
		if strings.ToLower(src.city) != strings.ToLower(dest.city) {
			return false
		}
	}
	// fmt.Println("city matched")
	if src.state != "" {
		if strings.ToLower(src.state) != strings.ToLower(dest.state) {
			return false
		}
	}
	//fmt.Println("state matched")
	if src.country != "" {
		//println("src:  ", src.country, "::", dest.country)
		if strings.ToLower(src.country) == strings.ToLower(dest.country) {
			return true
		}
	}

	//fmt.Println("nothing matched")
	return false
}

func isAddressAllow(key string, permission *Permission) bool {
	if isExcluded(key, permission) {
		return false
	}
	return isIncluded(key, permission)
}

func hasPermission(key string, permission *Permission) bool {
	var permissionBool bool
	//fmt.Println("hasPermission", "key:", key)
	if permission.parent != nil {
		permissionBool = hasPermission(key, permission.parent)
		if permissionBool == false {
			//fmt.Println("hasPermission", "parent does not have permission")
			return false
		}
	}
	return isAddressAllow(key, permission)
}

// initialies data transformation using map-reduce-filter
func foundinList(filePath string, address string) {
	fInput.TextFile(
		filePath, 2,
	).Map(func(line string, out chan flow.KeyValue) {

		array := strings.Split(line, ",")
		//marks, _ := strconv.Atoi(array[2])
		key := strings.ToUpper(array[3]) + "-" + strings.ToUpper(array[4]) + "-" + strings.ToUpper(array[5])
		out <- flow.KeyValue{Key: key, Value: line}
	}).Filter(func(key string, value string) bool {
		return strings.ToLower(key) == strings.ToLower(address)
	}).Map(func(key string, value string) string {
		fmt.Println(key, "available in the cities list as well")
		return key
	})
}

func init() {
	permission1 := &Permission{}
	permission1.id = "DISTRIBUTOR1"
	permission1.include = []string{"INDIA", "UNITED STATES"}
	permission1.exclude = []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"}

	permission2 := &Permission{}
	permission2.id = "DISTRIBUTOR2"
	permission2.parent = permission1
	permission2.include = []string{"INDIA"}
	permission2.exclude = []string{"TAMILNADU-INDIA"}

	permission3 := &Permission{}
	permission3.id = "DISTRIBUTOR3"
	permission3.parent = permission2
	permission3.include = []string{"HUBLI-KARNATAKA-INDIA"}

	distributerMap = make(map[string]*Permission)
	distributerMap["DISTRIBUTOR1"] = permission1
	distributerMap["DISTRIBUTOR2"] = permission2
	distributerMap["DISTRIBUTOR3"] = permission3
	//	distributerMap["DISTRIBUTOR4"] = permission4

	argsWithProg := os.Args
	// fmt.Println(os.Args)

	distributionKey := "DISTRIBUTOR3"

	//key := "CHICAGO-ILLINOIS-UNITED STATES"
	// key := "CHENNAI-TAMILNADU-INDIA"
	// key := "BANGALORE-KARNATAKA-INDIA"
	key := "HUBLI-KARNATAKA-INDIA"

	if len(argsWithProg) > 1 {
		distributionKey = argsWithProg[1]
	}

	if len(argsWithProg) > 2 {
		key = argsWithProg[2]
	}

	permission := distributerMap[distributionKey]
	if hasPermission(key, permission) {
		fmt.Println(distributionKey, " has Permission for ", key)

	} else {
		fmt.Println(distributionKey, " does not has Permission for ", key)
	}

	foundinList("cities.csv", key)
}

func main() {
	// We need to add this magic line to compute it in cluster.
	flag.Parse()
	flow.Ready()
	// start computing
	fInput.Run()
}

// command
// go run realimage.go DISTRIBUTOR1 CHICAGO-ILLINOIS-UNITEDSTATES
// go run realimage.go DISTRIBUTOR3 "HUBLI-KARNATAKA-INDIA"
// go run realimage.go DISTRIBUTOR2 "CHICAGO-ILLINOIS-UNITEDSTATES"
