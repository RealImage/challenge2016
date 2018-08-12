package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Dist holds dist
type Dist struct {
	Name      string
	Parent    string
	Country   map[string]Country
	ExCountry map[string]Country
}

type permission struct {
	Country  string
	State    string
	Province string
}

//Country hold contry
type Country struct {
	Code  string
	Name  string
	State map[string]State
}

//State hold state info
type State struct {
	Code     string
	Name     string
	Province map[string]Province
}

//Province holds province
type Province struct {
	Name string
	Code string
}

//Data holds csv Data

var (
	//Data holds the CSV data
	Data map[string]Country
	//DistData holds Distdat
	DistData = map[string]Dist{}
)

func main() {
	//reading csv file
	err := read("cities.csv")
	if err != nil {
		log.Fatal(err)
	}

	var exitFlag bool
	for !exitFlag {
		var action int
		writer := bufio.NewWriter(os.Stdin)
		scanner := bufio.NewScanner(os.Stdin)
		writer.Write([]byte("\n\nSELECT AN OPTION\n\nADD DIST\t\t:\t\t1\nADD SUB DIST\t\t:\t\t2\nCHECK PERMMISION\t:\t\t3\nExit\t\t\t:\t\t0\n\nPlease select:\t"))
		writer.Flush()
		fmt.Scanf("%d", &action)
		switch action {
		case 0: // EXIT
			exitFlag = true
		case 1: //ADD DIST
			writer.Write([]byte("Please enter the details in the given format\nName\nInclude Permission\nExclude Permission\n\nExample:\nprabesh\nIN,KL\nIN,KL,VRKLA\n\n"))
			writer.Flush()
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			inc := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			exc := strings.TrimSpace(scanner.Text())
			//checking if the given permission are valid or not
			if !(Valid(getPermission(inc)) && Valid(getPermission(exc))) {
				fmt.Println("Not a valid permission for the given data set")
				break
			}

			if getPermission(inc).Country == "" {
				fmt.Println("All permission assignment is not allowed")
				break
			}
			getDist(name).Add(getPermission(inc), getPermission(exc))
			fmt.Println("Permission succesfully added")

		case 2: //ADD SUB DISTRIBUTOR
			writer.Write([]byte("Please enter the details in the given format\nName\nParent Distributor Name\nInclude Permission\nExclude Permission\n\nExample:\nprabesh\nprajesh\nIN,KL\nIN,KL,VRKLA\n\n"))
			writer.Flush()
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			parentName := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			inc := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			exc := strings.TrimSpace(scanner.Text())
			if !AddSubDistributor(name, parentName, getPermission(inc), getPermission(exc)) {

			}

		case 3: //CHECK FOR PERMISSION
			writer.Write([]byte("\nPlease enter the details in the given format\nName\nPermission\n\nExample:\nprabesh\nIN,KL,VRKLA\n\n"))
			writer.Flush()
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			per := strings.TrimSpace(scanner.Text())
			if getDist(name).Check(getPermission(per)) {
				fmt.Printf("\n%s has permission\n", name)
			} else {
				fmt.Printf("\n%s has no permission\n", name)
			}

		default:
			fmt.Println("Invalid input")

		}
	}

}
func getPermission(s string) permission {
	g := strings.Split(s, ",")
	if len(g) == 3 {
		return permission{Country: g[0], State: g[1], Province: g[2]}
	} else if len(g) == 2 {
		return permission{Country: g[0], State: g[1]}
	} else if len(g) == 1 {
		return permission{Country: g[0]}
	}
	return permission{}
}

func read(file string) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	return process(handle)
}
func process(handle io.Reader) error {
	reader := csv.NewReader(bufio.NewReader(handle))
	Data = make(map[string]Country)
	type data struct {
		provinceCode string
		stateCode    string
		countryCode  string
		province     string
		state        string
		country      string
	}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		d := data{line[0], line[1], line[2], line[3], line[4], line[5]}
		if _, ok := Data[d.countryCode]; ok {
			if _, ok := Data[d.countryCode].State[d.stateCode]; ok {
				Data[d.countryCode].State[d.stateCode].Province[d.provinceCode] = Province{d.province, d.provinceCode}
			} else {
				Data[d.countryCode].State[d.stateCode] = State{Code: d.stateCode, Name: d.state,
					Province: map[string]Province{
						d.provinceCode: Province{d.province, d.provinceCode},
					}}
			}
		} else {
			Data[d.countryCode] = Country{Code: d.countryCode, Name: d.country,
				State: map[string]State{
					d.stateCode: State{Code: d.stateCode, Name: d.state,
						Province: map[string]Province{
							d.provinceCode: Province{d.province, d.provinceCode},
						}},
				}}
		}

	}
	return nil
}

func getDist(name string) *Dist {
	if v, ok := DistData[name]; ok {
		return &v
	}
	return &Dist{Name: name, Country: make(map[string]Country), ExCountry: make(map[string]Country)}
}

//AddSubDistributor adds subsdistributor
func AddSubDistributor(name, parentName string, inc, exc permission) bool {
	if !getDist(parentName).Check(inc) {
		return false
	}
	dis := getDist(name)
	dis.Parent = parentName
	dis.Add(inc, exc)
	return true
}

//Add adds a distributor
func (per *Dist) Add(inc, ex permission) {

	per.ExCountry = deleteInlcudedKeyfromMap(inc, per.ExCountry, true)
	per.Country = addKeyToMap(inc, per.Country)
	per.Country = deleteInlcudedKeyfromMap(ex, per.Country, false)
	per.ExCountry = addKeyToMap(ex, per.ExCountry)

	DistData[per.Name] = *per
}

func addKeyToMap(inc permission, country map[string]Country) map[string]Country {
	if c, ok := country[inc.Country]; ok {
		if inc.State != "" {
			if c.State != nil {
				if s, ok := c.State[inc.State]; ok {
					if inc.Province != "" {
						if s.Province != nil {
							country[inc.Country].State[inc.State].Province[inc.Province] = Province{}
						}
					} else {
						country[inc.Country].State[inc.State] = State{}
					}
				} else {
					country[inc.Country].State[inc.State] = State{Province: map[string]Province{
						inc.Province: Province{},
					}}
				}
			}
		} else {
			country[inc.Country] = Country{}
		}

	} else {
		if inc.Country != "" {
			country[inc.Country] = Country{}
			if inc.State != "" {
				country[inc.Country] = Country{State: map[string]State{
					inc.State: State{},
				}}
			}
			if inc.Province != "" {
				country[inc.Country].State[inc.State] = State{Province: map[string]Province{
					inc.Province: Province{},
				}}
			}
		}
	}

	return country
}

func deleteInlcudedKeyfromMap(inc permission, country map[string]Country, flag bool) map[string]Country {
	if c, ok := country[inc.Country]; ok {
		if c.State != nil {
			if s, ok := c.State[inc.State]; ok {
				if s.Province != nil {
					if _, ok := s.Province[inc.Province]; ok {
						delete(country[inc.Country].State[inc.State].Province, inc.Province)
						if len(country[inc.Country].State[inc.State].Province) == 0 {
							delete(country[inc.Country].State, inc.State)
							if len(country[inc.Country].State) == 0 {
								delete(country, inc.Country)
							}
						}
					}
				} else {
					if flag {
						delete(country[inc.Country].State, inc.State)
						if len(country[inc.Country].State) == 0 {
							delete(country, inc.Country)
						}
					}
				}
			}
		}
	}
	return country
}

//Valid checks if given permission is valid or not
func Valid(per permission) bool {
	if per.Country != "" {
		if _, ok := Data[per.Country]; !ok {
			return false
		}
	}
	if per.State != "" {
		if _, ok := Data[per.Country].State[per.State]; !ok {
			return false
		}
	}
	if per.Province != "" {
		if _, ok := Data[per.Country].State[per.State].Province[per.Province]; !ok {
			return false
		}
	}
	return true
}

//Check if permission exists or not
func (per Dist) Check(p permission) bool {
	if c, ok := per.ExCountry[p.Country]; ok {
		if c.State == nil {
			return false
		}
		if p.State == "" && len(c.State) >= 0 {
			return false
		}
		if s, ok := c.State[p.State]; ok {
			if s.Province == nil {
				return false
			}
			if p.Province == "" && len(s.Province) >= 0 {
				return false
			}
			if _, ok := s.Province[p.Province]; ok {
				return false
			}
		}
	}
	if c, ok := per.Country[p.Country]; ok {
		if c.State == nil {
			return true
		}
		if p.State == "" && len(c.State) == 0 {
			return true
		}
		if s, ok := c.State[p.State]; ok {
			if s.Province == nil {
				return true
			}
			if p.Province == "" && len(s.Province) == 0 {
				return true
			}
			if _, ok := s.Province[p.Province]; ok {
				return true
			}
		}
	}
	return false
}
