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
	Name    string
	Country map[string]Country
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

	csvFile, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	Data = make(map[string]Country)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if _, ok := Data[line[2]]; ok {
			if _, ok := Data[line[2]].State[line[1]]; ok {
				Data[line[2]].State[line[1]].Province[line[0]] = Province{line[3], line[0]}
			} else {
				Data[line[2]].State[line[1]] = State{Code: line[1], Name: line[4],
					Province: map[string]Province{
						line[0]: Province{line[3], line[0]},
					}}
			}
		} else {
			Data[line[2]] = Country{Code: line[2], Name: line[5],
				State: map[string]State{
					line[1]: State{Code: line[1], Name: line[4],
						Province: map[string]Province{
							line[0]: Province{line[3], line[0]},
						}},
				}}
		}

	}
Loop:
	for {
		var action int
		fmt.Println("*************************")
		fmt.Println()
		fmt.Println("SELECT AN OPTION        ")
		fmt.Println("")
		fmt.Println("ADD DIST       : ", 1)
		fmt.Println("DEL DIST       : ", 2)
		fmt.Println("ADD SUB DIST       : ", 4)
		fmt.Println("CHECK PERMMISION   : ", 5)
		fmt.Println("EXIT           : ", 0)
		fmt.Println()
		fmt.Println("*************************")
		fmt.Scanf("%d", &action)
		switch action {
		case 0: // EXIT
			break Loop
		case 1: //ADD DIST
			var name, incl, excl string
			fmt.Println("Please enter the name ")
			fmt.Scanf("%s", &name)
			fmt.Println("Please enter the inlucde permissions in a comma seperated way ")
			fmt.Scanf("%s", &incl)
			fmt.Println("Please enter the exclude permissions in a comma seperated way ")
			fmt.Scanf("%s", &excl)
			if incl == "" {
				fmt.Println("All permission assignment is not allowed")
				break
			}
			getDist(name).Add(incl, excl, name)

		case 2: //DEL A DISTRIBUTOR
			var name string
			fmt.Println("Please enter the name ")
			fmt.Scanf("%s", &name)
			if _, ok := DistData[name]; ok {
				delete(DistData, name)
			} else {
				fmt.Printf("Distributor with given name %s is not present in the list\n", name)
			}
			break

		case 4:
			var name, pname, incl, excl string
			fmt.Println("Please enter name ")
			fmt.Scanf("%s", &name)
			fmt.Println("Please enter parent Dist name ")
			fmt.Scanf("%s", &pname)
			fmt.Println("Please enter the inlucde permissions in a comma seperated way ")
			fmt.Scanf("%s", &incl)
			fmt.Println("Please enter the exclude permissions in a comma seperated way ")
			fmt.Scanf("%s", &excl)
			if incl == "" {
				fmt.Println("All permission is not allowed")
				break
			}
			str := incl
			if excl != "" {
				str = incl + "," + excl
			}
			s := strings.Split(str, ",")
			if !Valid(s) {
				fmt.Println("Not a Valid permission")
				break
			}
			if !getDist(pname).Check(str) {
				fmt.Printf("%s has no permission to distribute\n", pname)
				break
			}
			dist := Dist{name, getDist(pname).copyMap()}
			dist.Add(incl, excl, name)
			break
		case 5: //CHECK PERMISSION
			var name, p string
			fmt.Println("Please enter the name of ditributor")
			fmt.Scanf("%s", &name)
			fmt.Println("Please enter the permission in a comma seperated way (country,state,province)")
			fmt.Scanf("%s", &p)
			s := strings.Split(p, ",")
			if !Valid(s) || len(s) < 3 {
				fmt.Println("Not a Valid permission")
				break
			}
			if !getDist(name).Check(p) {
				fmt.Printf("%s has no permission\n", name)
				break
			}
			fmt.Printf("%s has permission\n", name)
			break

		default:
			fmt.Println("Invalid input")

		}
	}

}

func getDist(name string) Dist {
	if v, ok := DistData[name]; ok {
		return v
	}
	return Dist{Name: name, Country: make(map[string]Country)}
}

//Add adds a distributor
func (per Dist) Add(inc, ex, name string) {
	s := strings.Split(inc, ",")

	if len(s) == 2 { // case with two values in include
		if _, ok := per.Country[s[0]]; ok {
			if per.Country[s[0]].State != nil {
				if _, ok := per.Country[s[0]].State[s[1]]; ok {
					if per.Country[s[0]].State[s[1]].Province == nil {
						per.Country[s[0]].State[s[1]] = State{
							Province: map[string]Province{},
						}
					}
				} else {
					per.Country[s[0]].State[s[1]] = State{
						Province: map[string]Province{},
					}
				}
			} else {
				per.Country[s[0]].State[s[1]] = State{
					Province: map[string]Province{},
				}
			}
		} else {
			per.Country[s[0]] = Country{State: map[string]State{
				s[1]: State{
					Province: map[string]Province{},
				},
			}}
		}
		per.Country[s[0]].State[s[0]+":all"] = State{}
		if ex != "" {
			if per.Country[s[0]].State[s[1]].Province != nil {
				per.Country[s[0]].State[s[1]].Province[ex] = Province{}
			} else {
				per.Country[s[0]].State[s[1]] = State{
					Province: map[string]Province{
						ex: Province{},
					},
				}
			}
		}

	} else if len(s) == 1 { //case with a single include string

		if _, ok := per.Country[s[0]]; !ok {
			per.Country[s[0]] = Country{
				State: map[string]State{},
			}
		}
		if per.Country[s[0]].State != nil {

			delete(per.Country[s[0]].State, s[0]+":all")
		} else {
			per.Country[s[0]] = Country{
				State: map[string]State{},
			}
		}

		//adding exlude permission part
		if ex == "" {
			return
		}
		e := strings.Split(ex, ",")
		if len(e) == 1 {
			if _, ok := per.Country[s[0]].State[e[0]]; !ok {
				per.Country[s[0]].State[e[0]] = State{
					Province: map[string]Province{
						e[0] + ":all": Province{},
					},
				}
			}
		} else {
			if _, ok := per.Country[s[0]].State[e[0]]; !ok {
				per.Country[s[0]].State[e[0]] = State{
					Province: map[string]Province{
						e[1]:          Province{},
						e[0] + ":all": Province{},
					},
				}
			}
			//adding province details
			if per.Country[s[0]].State[e[0]].Province != nil {
				per.Country[s[0]].State[e[0]].Province[e[1]] = Province{}
				per.Country[s[0]].State[e[0]].Province[e[0]+":all"] = Province{}
			} else {
				per.Country[s[0]].State[e[0]] = State{
					Province: map[string]Province{
						e[1]:          Province{},
						e[0] + ":all": Province{},
					},
				}
			}
		}

	} else if len(s) == 3 { // with just inlcude string
		if _, ok := per.Country[s[0]]; ok {
			if per.Country[s[0]].State != nil {
				if _, ok := per.Country[s[0]].State[s[1]]; ok {
					if per.Country[s[0]].State[s[1]].Province != nil {
						delete(per.Country[s[0]].State[s[1]].Province, s[1]+":all")
						delete(per.Country[s[0]].State[s[1]].Province, s[2])
					}
				}
			}

		}
	}
	DistData[name] = per
}

//Valid checks if given permission is valid or not
func Valid(s []string) bool {
	if len(s) == 3 {
		if _, ok := Data[s[0]]; ok {
			if _, ok := Data[s[0]].State[s[1]]; ok {
				if _, ok := Data[s[0]].State[s[1]].Province[s[2]]; ok {
					return true
				}
			}
		}
	} else if len(s) == 2 {
		if _, ok := Data[s[0]]; ok {
			if _, ok := Data[s[0]].State[s[1]]; ok {
				return true
			}
		}
	} else if len(s) == 1 {
		if _, ok := Data[s[0]]; ok {
			return true
		}
	}
	return false
}

//Check if permission exists or not
func (per Dist) Check(p string) bool {

	s := strings.Split(p, ",")

	if len(s) < 3 {
		return false
	}

	if per.Country != nil {
		if _, ok := per.Country[s[0]]; ok {
			if per.Country[s[0]].State != nil {
				if _, ok := per.Country[s[0]].State[s[1]]; ok {
					if per.Country[s[0]].State[s[1]].Province != nil {
						if _, ok := per.Country[s[0]].State[s[1]].Province[s[2]]; !ok {
							if _, ok := per.Country[s[0]].State[s[1]].Province[s[1]+":all"]; !ok {
								return true
							}
						}
					} else {
						return true
					}
				} else {
					if _, ok := per.Country[s[0]].State[s[0]+":all"]; !ok {
						return true
					}
				}
			} else {
				return true
			}

		}

	}
	return false
}

func (per Dist) copyMap() map[string]Country {
	m := make(map[string]Country)
	for i := range per.Country {
		s := make(map[string]State)
		for j := range per.Country[i].State {
			p := make(map[string]Province)
			for x := range per.Country[i].State[j].Province {
				p[x] = per.Country[i].State[j].Province[x]
			}
			s[j] = State{Province: p}
		}
		m[i] = Country{State: s}
	}
	return m
}
