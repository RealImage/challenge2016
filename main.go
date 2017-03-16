package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type city struct {
	Code         string
	ProvinceCode string
	CountryCode  string
	Name         string
	ProvinceName string
	CountryName  string
}

type node struct {
	Name string
	Code string

	Nodes []*node
}

func (c *city) RuleArea() *ruleArea {
	return &ruleArea{
		Name:         c.Name,
		ProvinceName: c.ProvinceName,
		CountryName:  c.CountryName,
	}
}

type ruleArea struct {
	Name         string
	ProvinceName string
	CountryName  string
}

func (ra *ruleArea) String() string {
	area := toInput(ra.Name) + "-" + toInput(ra.ProvinceName) + "-" + toInput(ra.CountryName)

	return strings.TrimPrefix(area, "-")
}

func subAreas(d *distributor) []string {
	var sas []string
	tnodes := topNodes

	var chain []*distributor
	currentNode := d

	for currentNode != nil {
		chain = append(chain, currentNode)
		currentNode = currentNode.Parent
	}

	for i := len(chain) - 1; i >= 0; i-- {
		bnodes := []*node{}
		tempDistributor := chain[i]
		tnodes = filter(tnodes, tempDistributor.Include, tempDistributor.Exclude)
		for _, tnode := range tnodes {
			bnodes = append(bnodes, tnode.Nodes...)
		}

		if len(bnodes) > 0 {
			tnodes = bnodes
		}
	}

	for _, c := range tnodes {
		sas = append(sas, c.Code)
	}

	return sas
}

func filter(tc []*node, includes, excludes []string) []*node {
	var fc, cc []*node
	for _, c := range tc {
		ok := false
		for _, allowed := range includes {
			if strings.HasSuffix(c.Code, allowed) {
				ok = true
				break
			}
		}
		if ok {
			cc = append(cc, c)
		}
	}

	for _, c := range cc {
		ok := false
		for _, exclude := range excludes {
			if strings.HasSuffix(c.Code, exclude) {
				ok = true
				break
			}
		}
		if !ok {
			fc = append(fc, c)
		}
	}

	return fc
}

type distributor struct {
	Name    string
	Include []string
	Exclude []string

	Parent   *distributor
	Children []*distributor
}

func (d *distributor) Allow(r *ruleArea) bool {
	sas := subAreas(d)
	for _, sa := range sas {
		if strings.HasSuffix(sa, r.String()) {
			return true
		}
	}

	return false
}

var topNodes []*node

func main() {
	file, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)

	_, err = r.Read() //header row
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic("oh no")
		}
		c := &city{Code: record[0], ProvinceCode: record[1], CountryCode: record[2], Name: toInput(record[3]), ProvinceName: toInput(record[4]), CountryName: toInput(record[5])}

		putCityInTree(c)
	}

	d1 := &distributor{
		Name:    "DISTRIBUTOR1",
		Include: []string{"INDIA", "UNITEDSTATES"},
		Exclude: []string{"TAMILNADU-INDIA"},
	}

	d2 := &distributor{
		Name:    "DISTRIBUTOR2",
		Include: []string{"INDIA"},
		Exclude: []string{"CHENNAI-TAMILNADU-INDIA"},
		Parent:  d1,
	}

	d1.Children = append(d1.Children, d2)

	d3 := &distributor{
		Name:    "DISTRIBUTOR3",
		Include: []string{"ANDHRAPRADESH-INDIA", "BENGALURU-KARNATAKA-INDIA"},
		Parent:  d2,
	}

	d2.Children = append(d2.Children, d3)

	area := &ruleArea{
		Name:         "Ayakudi",
		ProvinceName: "TAMILNADU",
		CountryName:  "India",
	}

	ok := d3.Allow(area)
	if ok {
		println("YES")
	} else {
		println("NO")
	}
}

func putCityInTree(c *city) {
	var cityCountry *node
	var cityState *node
	var cityArea *node

	cName := toInput(c.CountryName)
	var has bool
	for _, c := range topNodes {
		if c.Name == cName {
			cityCountry = c
			has = true
			break
		}
	}

	if !has {
		cityCountry = &node{
			Name: cName,
			Code: cName,
		}
		topNodes = append(topNodes, cityCountry)
	} else {
		has = false
		cState := toInput(c.ProvinceName)
		for _, s := range cityCountry.Nodes {
			if s.Name == cState {
				cityState = s
				has = true
				break
			}
		}
		if !has {
			cityState = &node{
				Name: cState,
				Code: cState + "-" + cityCountry.Code,
			}
			cityCountry.Nodes = append(cityCountry.Nodes, cityState)
		} else {
			has = false
			cArea := toInput(c.Name)
			for _, a := range cityState.Nodes {
				if a.Name == cArea {
					cityArea = a
					has = true
					break
				}
			}
			if !has {
				cityArea = &node{
					Name: cArea,
					Code: cArea + "-" + cityState.Code,
				}
				cityState.Nodes = append(cityState.Nodes, cityArea)
			}
		}

		cityCountry.Nodes = append(cityCountry.Nodes, cityState)
	}

}

func parseRule(r string) *ruleArea {
	parts := strings.Split(r, "-")
	switch len(parts) {
	case 1:
		return &ruleArea{
			CountryName: parts[0],
		}
	case 2:
		return &ruleArea{
			ProvinceName: parts[0],
			CountryName:  parts[1],
		}
	case 3:
		return &ruleArea{
			Name:         parts[0],
			ProvinceName: parts[1],
			CountryName:  parts[2],
		}
	default:
		return nil
	}
}

func toInput(s string) string {
	return strings.ToUpper(strings.Replace(s, " ", "", -1))
}
