package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Distribution struct {
	Distributor *DistributorNode
}

type DistributorNode struct {
	children map[string]*DistributorNode
	exclude  bool
	include  bool
	leaf     bool
}

type DistributorPermissions map[string]Distribution

func NewDistributorNode() *DistributorNode {
	return &DistributorNode{
		children: make(map[string]*DistributorNode),
		exclude:  false,
		include:  false,
		leaf:     false,
	}
}

func (root *DistributorNode) Insert(region string, included bool) {
	node := root
	parts := strings.Split(region, "-")

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else {
			newNode := NewDistributorNode()
			node.children[part] = newNode
			node = newNode
		}
	}

	node.leaf = true
	if included {
		node.include = true
	} else {
		node.exclude = true
	}
}

func (root *DistributorNode) Search(region string) bool {
	node := root
	parts := strings.Split(region, "-")

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else {
			return false
		}

		if node.exclude {
			return false
		}
	}

	return node.include || node.leaf
}

func LoadDistributions(filepath string) (DistributorPermissions, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	permissions := make(DistributorPermissions)
	for _, record := range records {
		distributor := record[0]
		permission := record[1]
		region := record[2]

		distribution, exists := permissions[distributor]
		if !exists {
			distribution = Distribution{
				Distributor: NewDistributorNode(),
			}
			permissions[distributor] = distribution
		}

		if permission == "INCLUDE" {
			distribution.Distributor.Insert(region, true)
		}
	}

	return permissions, nil
}

func checkPermission(permissions DistributorPermissions, distributor string, region string) bool {
	permission, exists := permissions[distributor]
	if !exists {
		return false
	}

	return permission.Distributor.Search(region)
}

func main() {
	distributions, err := LoadDistributions("data.csv")

	if err != nil {
		log.Fatal("Error loading distributions:", err)
	}

	distributor1 := "DISTRIBUTOR1"
	distributor2 := "DISTRIBUTOR2"
	distributor3 := "DISTRIBUTOR3"

	region1 := "CHICAGO-ILLINOIS-UNITEDSTATES"
	region2 := "CHENNAI-TAMILNADU-INDIA"
	region3 := "BANGALORE-KARNATAKA-INDIA"

	fmt.Println(checkPermission(distributions, distributor1, region1)) // outputs true
	fmt.Println(checkPermission(distributions, distributor1, region2)) // outputs false
	fmt.Println(checkPermission(distributions, distributor1, region3)) // outputs false

	fmt.Println(checkPermission(distributions, distributor2, region1)) // outputs false
	fmt.Println(checkPermission(distributions, distributor2, region2)) // outputs false
	fmt.Println(checkPermission(distributions, distributor2, region3)) // outputs true

	fmt.Println(checkPermission(distributions, distributor3, region1)) // outputs false
	fmt.Println(checkPermission(distributions, distributor3, region2)) // outputs false
	fmt.Println(checkPermission(distributions, distributor3, region3)) // outputs false
}
