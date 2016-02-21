package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputCsv        = "./cities.csv"
	distributorFile = "./distributors.rule"
)

const (
	permis    = "Permissions:"
	include   = "INCLUDE:"
	exclude   = "EXCLUDE:"
	separator = "<"
)

// Node is to sepcify single row data of csv
type Node struct {
	country string
	state   string
	city    string
}

// NewNode is to get new struct
func NewNode() *Node {
	return &Node{}
}

// Distributor is to define a new Distributor
type Distributor struct {
	Name              string
	Includes          []string
	Excludes          []string
	ParentDistributor *Distributor
}

// NewDistributor is to get a new struct of Distributor
func NewDistributor() *Distributor {
	return &Distributor{}
}

var dataStore []*Node
var distributors []*Distributor

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: Scriptname <PathToRuleFile> <Distributor> <Location> \n Will give " +
			"YES or NO")
		return
	}
	distributorFile = os.Args[1]
	d := os.Args[2]
	loc := os.Args[3]

	loadData()
	loadDistributorRule()
	validateCode(loc)
	city, state, country := getSplittedLocation(loc)
	res := make(chan bool, 1)
	go isDistributorAllowed(d, country, state, city, res)

	v := <-res
	if v {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func getSplittedLocation(loc string) (string, string, string) {
	aLoc := strings.Split(strings.Trim(loc, " \n"), "-")
	city, state, country := "", "", ""

	if len(aLoc) > 2 {
		city = aLoc[0]
		state = aLoc[1]
		country = aLoc[2]
	} else if len(aLoc) > 1 {
		state = aLoc[0]
		country = aLoc[1]
	} else {
		country = aLoc[0]
	}
	return city, state, country
}

func isDistributorAllowed(name string, country string, state string, city string, res chan bool) {
	dis, err := getDistributor(name)
	if err != nil {
		fmt.Print(err)
	}
	included := true
	for {
		// Uncomment this to see the rule order
		//fmt.Println(dis)
		for _, exc := range dis.Excludes {
			ci, s, co := getSplittedLocation(exc)
			if isSubset(ci, s, co, city, state, country) {
				res <- false
			}
		}

		innerLoopInc := false
		for _, ic := range dis.Includes {
			ci, s, co := getSplittedLocation(ic)
			if isSubset(ci, s, co, city, state, country) {
				innerLoopInc = true
				break
			}
		}
		included = included && innerLoopInc

		if dis.ParentDistributor == nil {
			break
		}
		dis = dis.ParentDistributor
	}
	res <- included

}

func isIncluded(name string, country string, state string, city string, dis *Distributor, res chan bool) {

}

func isSubset(ci string, s string, co string, ci1 string, s1 string, co1 string) bool {
	if co == co1 {
		if s == "" {
			return true
		} else if s == s1 {
			if ci == "" {
				return true
			} else if ci == ci1 {
				return true
			}
		}
	}
	return false
}

func loadData() {
	data, err := ioutil.ReadFile(inputCsv)
	if err != nil {
		fmt.Println("While reading file", err)
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Parsing csv", err)
			break
		}
		if len(rec) < 3 {
			continue
		}
		// DEVUA KL IN Devikulam Kerala India
		ci := rec[0]
		s := rec[1]
		c := rec[2]
		insertData(ci, s, c)
	}
}

func insertData(city string, state string, country string) {
	node := &Node{country: country, state: state, city: city}
	dataStore = append(dataStore, node)
}

func loadDistributorRule() {
	data, err := ioutil.ReadFile(distributorFile)
	if err != nil {
		fmt.Println("While reading distributors file", err)
	}

	r := bufio.NewReader(strings.NewReader(string(data)))
	lastName := ""
	var lastParent *string
	newGuy := false
	var name string
	var includes []string
	var excludes []string
	var parentDis *string
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			addDistributor(lastName, includes, excludes, lastParent)
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		if strings.Contains(line, permis) {
			r := strings.Split(line, permis)[1]
			if strings.Contains(r, separator) {

				dis := strings.Split(r, separator)
				name = strings.Trim(strings.ToUpper(dis[0]), "\n ")
				pd := strings.Trim(strings.ToUpper(dis[1]), "\n ")
				parentDis = &pd
			} else {
				name = strings.Trim(strings.ToUpper(r), "\n ")
			}
			if lastName == "" {
				// coming inside first time
				lastName = name
				lastParent = parentDis
			}
			newGuy = true
		}

		if strings.Contains(line, include) {
			r := strings.Split(line, include)
			ic := strings.Trim(strings.ToUpper(r[1]), "\n ")
			validateCode(ic)
			includes = append(includes, ic)
			newGuy = false
		}
		if strings.Contains(line, exclude) {
			r := strings.Split(line, exclude)
			ec := strings.Trim(strings.ToUpper(r[1]), "\n ")
			validateCode(ec)
			excludes = append(excludes, ec)
			newGuy = false
		}

		if !newGuy || (newGuy && len(includes) == 0 && len(excludes) == 0) {
			continue
		}
		addDistributor(lastName, includes, excludes, lastParent)
		lastName = name
		lastParent = parentDis
		// Resetting all the values.
		name = ""
		includes = nil
		excludes = nil
		parentDis = nil

	}
}

func addDistributor(name string, inc []string, exc []string, parentDis *string) {
	var parentDistributor *Distributor
	var err error
	if parentDis != nil {
		parentDistributor, err = getDistributor(*parentDis)
		if err != nil {
			fmt.Println("Error:", err, *parentDis)
		}
	}
	dis := &Distributor{Name: name, Includes: inc, Excludes: exc,
		ParentDistributor: parentDistributor}
	distributors = append(distributors, dis)
}

func getDistributor(name string) (*Distributor, error) {
	name = strings.Trim(strings.ToUpper(name), "\n ")
	for _, dis := range distributors {
		if dis.Name == name {
			return dis, nil
		}
	}
	return nil, errors.New("NotFound")
}

func validateCode(code string) {
	ci, s, co := getSplittedLocation(code)
	found := false
	for _, d := range dataStore {
		if isSubset(ci, s, co, d.city, d.state, d.country) {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Invalid code " + code + " is given.")
		os.Exit(1)
	}
}
