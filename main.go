package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Authorization struct {
	Patricia *PatriciaNode
}

type PatriciaNode struct {
	children   map[string]*PatriciaNode
	isExcluded bool
	isIncluded bool
	isLeaf     bool
}

type DistributorAuthorizations map[string]Authorization

func NewPatriciaNode() *PatriciaNode {
	return &PatriciaNode{
		children:   make(map[string]*PatriciaNode),
		isExcluded: false,
		isIncluded: false,
		isLeaf:     false,
	}
}

func (root *PatriciaNode) Insert(region string, included bool) {
	node := root
	parts := strings.Split(region, "-")

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else {
			newNode := NewPatriciaNode()
			node.children[part] = newNode
			node = newNode
		}
	}

	node.isLeaf = true
	if included {
		node.isIncluded = true
	} else {
		node.isExcluded = true
	}
}

func (root *PatriciaNode) Search(region string) bool {
	node := root
	parts := strings.Split(region, "-")

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else {
			return false
		}

		if node.isExcluded {
			return false
		}
	}

	return node.isIncluded || node.isLeaf
}

func LoadAuthorizations(filePath string) (DistributorAuthorizations, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	authorizations := make(DistributorAuthorizations)
	for _, record := range records {
		distributor := record[0]
		permission := record[1]
		region := record[2]

		auth, exists := authorizations[distributor]
		if !exists {
			auth = Authorization{
				Patricia: NewPatriciaNode(),
			}
			authorizations[distributor] = auth
		}

		if permission == "INCLUDE" {
			auth.Patricia.Insert(region, true)
		} else if permission == "EXCLUDE" {
			auth.Patricia.Insert(region, false)
		}
	}

	return authorizations, nil
}

func HasPermission(authorizations DistributorAuthorizations, distributor, region string) bool {
	auth, exists := authorizations[distributor]
	if !exists {
		return false
	}

	return auth.Patricia.Search(region)
}

func main() {
	authorizations, err := LoadAuthorizations("storage-data.csv")
	if err != nil {
		log.Fatal("Error loading authorizations:", err)
	}

	distributor1 := "DISTRIBUTOR1"
	distributor2 := "DISTRIBUTOR2"
	distributor3 := "DISTRIBUTOR3"

	region1 := "CHICAGO-ILLINOIS-UNITEDSTATES"
	region2 := "CHENNAI-TAMILNADU-INDIA"
	region3 := "BANGALORE-KARNATAKA-INDIA"

	fmt.Println(HasPermission(authorizations, distributor1, region1)) // true
	fmt.Println(HasPermission(authorizations, distributor1, region2)) // false
	fmt.Println(HasPermission(authorizations, distributor1, region3)) // false

	fmt.Println(HasPermission(authorizations, distributor2, region1)) // true
	fmt.Println(HasPermission(authorizations, distributor2, region2)) // false
	fmt.Println(HasPermission(authorizations, distributor2, region3)) // false

	fmt.Println(HasPermission(authorizations, distributor3, region1)) // false
	fmt.Println(HasPermission(authorizations, distributor3, region2)) // false
	fmt.Println(HasPermission(authorizations, distributor3, region3)) // false
}
