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
	Country   map[string]Country
	ExCountry map[string]Country
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

			break
		case 5: //CHECK PERMISSION
			var name, p string
			fmt.Println("Please enter the name of ditributor")
			fmt.Scanf("%s", &name)
			fmt.Println("Please enter the permission in a comma seperated way (country,state,province)")
			fmt.Scanf("%s", &p)
			if !Valid(strings.Split(p, ",")) {
				fmt.Println("In Valid permission")
				break
			}
			break
		default:
			fmt.Println("Invalid input")

		}
	}

}

//Check if permission exists or not
func (per Dist) Check(p string) bool {
	return false
}

//Add adds a distributor
func (per Dist) Add(inc, ex, name string) {

}

//AddSub adds a distributor
func (per Dist) AddSub(inc, ex, name string) {

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
