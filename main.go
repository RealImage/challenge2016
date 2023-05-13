package main

import (
	"fmt"
	"strings"
	"encoding/csv"
	"os"
	"bufio"
)

type Permission struct {
	include []string
	exclude []string
}

type Location struct {
	City    string
	State   string
	Country string
}

type Distributor struct {
	name        string
	permission  Permission
	subDistributors []Distributor
}

var DistributerData map[string]Distributor
var PlacesMap map[string]Location
var subDistsByName map[string]string


func (p *Permission) IsAuthorized(region string,type_req string) bool {
	location, found := PlacesMap[region]
	new_reg1 := ""
	new_reg2 := ""
	new_reg3 := ""
	if !found {
		fmt.Println("Location not found")
	} else {
		if type_req == "city"{
			new_reg1 = region +"-"+location.State+"-"+location.Country
			new_reg2 = location.State+"-"+location.Country
			new_reg3 = location.Country
		}else if type_req == "state"{
			new_reg1 = location.State+"-"+location.Country
			new_reg2 = location.Country
			new_reg3 = location.Country
		}else if type_req == "country"{
			new_reg1 = location.Country
			new_reg2 = location.Country
			new_reg3 = location.Country
		}
	}
	for _, e := range p.exclude {
		if strings.HasPrefix(new_reg1, e) || strings.HasPrefix(new_reg2, e) || strings.HasPrefix(new_reg3, e) {
			return false
		}
	}

	for _, i := range p.include {
		if strings.HasPrefix(new_reg1, i) || strings.HasPrefix(new_reg2, i) || strings.HasPrefix(new_reg3, i){
			return true
		}
	}

	return false
}



func (d *Distributor) IsAuthorized(region string,type_req string) bool {
	if d.permission.IsAuthorized(region,type_req) {
		return true
	}

	return false
}


func main() {
	InitPlacesData()	
	main_menu:
		fmt.Println("Distributer List : ",DistributerData)
		fmt.Println("Select Options \n 1. Add new Distributor \n 2. Add Permission \n 3. Check distibutions \n 4. Add SubDistributors")
		arg := 0
		fmt.Scan(&arg)
		scanner := bufio.NewScanner(os.Stdin)
		switch arg {
			case 1:
				fmt.Println("Enter Distributer's Name: ")
				name := ""
			    scanner.Scan()
			    name = scanner.Text()
				fmt.Println("name",name)
				name = strings.ReplaceAll(strings.ToLower(name), " ", "_")
				_, ok := DistributerData[name]
				if ok {
					fmt.Println("Distributer already exit ")
					goto main_menu
				}else{

					distributer := &Distributor{
									name:name,
									permission: Permission{},
									subDistributors: []Distributor{},
								}
					DistributerData[name] = *distributer
					include_things := []string{}
					exclude_things := []string{}
					permissionOpt:
						fmt.Println("Adding Permissions..\n  1. Include \n 2. Exclude \n 3. Main menu")
						permissions := 0
						fmt.Scan(&permissions)
						switch permissions {
							case 1,2:
								fmt.Println("Enter Country name(enter nil if not needed) : ")
								country,state,city := "","",""
			
							    scanner.Scan()
							    country = scanner.Text()
								
								if country != "nil" && country != ""{
									fmt.Println("Enter State name(enter nil if not needed) :  ")
				
								    scanner.Scan()
								    state = scanner.Text()
									if state != "nil" && state != ""{
										fmt.Println("Enter City name(enter nil if not needed) : ")
					
									    scanner.Scan()
									    city = scanner.Text()
									}
								}
								include_perm := ""
								if (country != "nil" && country != "") || (state != "nil" && state != "") || (city != "nil" && city != ""){
									if (country != "nil" && country != ""){
										include_perm += strings.ReplaceAll(strings.ToLower(country), " ", "_")
									}
									if (state != "nil" && state != ""){
										include_perm = "-" + include_perm
										include_perm = strings.ReplaceAll(strings.ToLower(state), " ", "_") + include_perm
									}
									if (city != "nil" && city != ""){
										include_perm = "-" + include_perm
										include_perm = strings.ReplaceAll(strings.ToLower(city), " ", "_") + include_perm
									}
								}
								if permissions == 1{
									include_things = append(include_things,include_perm)
								}else{
									exclude_things = append(exclude_things,include_perm)
								}

								goto permissionOpt
							case 3:
								distributer.permission.include = include_things
								distributer.permission.exclude = exclude_things
								DistributerData[name] = *distributer
								goto main_menu
						}
				} 


			case 2:
				add_perm:
					fmt.Println("Enter Distributer/Sub Distributor's  Name: ")
					name := ""

				    scanner.Scan()
				    name = scanner.Text()
					fmt.Println("name",name)
					name = strings.ReplaceAll(strings.ToLower(name), " ", "_")
					var sub bool
					data, ok := DistributerData[name]
					dataa := Distributor{}
					if !ok{
						if para,ex := subDistsByName[name];ex{
							dataa, ok = DistributerData[para]
							if ok{
								ok,data = isSubDistributorOf(dataa,name)
								sub = true

							}
						}

					}
					if ok{
						include_things := []string{}
						exclude_things := []string{}
						permissionOpt2:
							fmt.Println("Adding Permissions..\n 1. Include \n 2. Exclude \n 3. Main menu")
							permissions := 0
							fmt.Scan(&permissions)
							switch permissions {
								case 1,2:
									fmt.Println("Enter Country name(enter nil if not needed) : ")
									country,state,city := "nil","nil","nil"
				
								    scanner.Scan()
								    country = scanner.Text()
									
									if country != "nil" && country != ""{
										country = strings.ReplaceAll(strings.ToLower(country), " ", "_")
										fmt.Println("Enter State name(enter nil if not needed) :  ")
					
									    scanner.Scan()
									    state = scanner.Text()
										if state != "nil" && state != ""{
											state = strings.ReplaceAll(strings.ToLower(state), " ", "_")
											fmt.Println("Enter City name(enter nil if not needed) : ")
						
										    scanner.Scan()
										    city = scanner.Text()
										    if city !="nil" && state != ""{
										    	city = strings.ReplaceAll(strings.ToLower(city), " ", "_")
										    }
										}
									}
									include_perm := ""
									auth := false
									if (country != "nil" && country != "") || (state != "nil" && state != "") || (city !="nil" && city != ""){
										if sub{
											if (city !="nil" && city != ""){
												auth = dataa.IsAuthorized(city,"city")
											}else if (state != "nil" && state != ""){
												auth = dataa.IsAuthorized(state,"state")

											}else if (country != "nil" && country != ""){
												auth = dataa.IsAuthorized(country,"country")

											}
										}
										if !sub || auth{	
											if (country != "nil" && country != ""){
												include_perm += country
											}
											if (state != "nil" && state != ""){
												include_perm = "-" + include_perm
												include_perm = state + include_perm
											}
											if (city !="nil" && city != ""){
												include_perm = "-" + include_perm
												include_perm = city + include_perm
											}
										}
									}
									if permissions == 1{
										include_things = append(include_things,include_perm)
									}else{
										exclude_things = append(exclude_things,include_perm)
									}

									if sub && !auth{
										fmt.Println("You Parent dont have the access in this place \n")
									}

									goto permissionOpt2
								case 3:
									data.permission.include = append(data.permission.include,include_things...)
									data.permission.exclude = append(data.permission.exclude,exclude_things...)
									if sub{
										dataa.subDistributors = []Distributor{}
										dataa.subDistributors = append(dataa.subDistributors,data)
										DistributerData[strings.ToLower(dataa.name)] = dataa
									}else{
										DistributerData[strings.ToLower(name)] = data
									}
									goto main_menu
							}
					}else{
						fmt.Println("Enter valid Distributer \n")
						goto add_perm
					}

			case 3:
					check_new_dist:
						fmt.Println("Enter Distributer/Sub Distributor's Name to check  the permission: \n")
						name := ""
	
					    scanner.Scan()
					    name = scanner.Text()
						name = strings.ReplaceAll(strings.ToLower(name), " ", "_")

						data, ok := DistributerData[strings.ToLower(name)]

						if !ok{
							if para,ex := subDistsByName[name];ex{
								dataa, okk := DistributerData[para]
								if okk{
									ok,data = isSubDistributorOf(dataa,name)
								}
							}

						}

						if ok{
							check_dist:
								fmt.Println("\nSearch by \n ")
								fmt.Println(" 1. City \n 2. State \n 3. Country  \n 4. Check again \n 5. Main menu\n")
								input := 0
								fmt.Scan(&input)
								region := ""
								if input == 1 || input == 2 || input == 3 {
									fmt.Println("Enter the Place \n")
				
								    scanner.Scan()
								    region = scanner.Text()

									region = strings.ReplaceAll(strings.ToLower(region), " ", "_")
								}

								type_req := ""
								switch input {
									case 1:
										type_req = "city"
									case 2:
										type_req = "state"
									case 3:
										type_req = "country"
									case 4:
										goto check_dist
									case 5:
										goto main_menu
									default:
										fmt.Println("Enter Valid input")
										goto check_dist

								}
							auth := data.IsAuthorized(region,type_req)
							if auth{
								fmt.Printf("%s Can Distribute in %s",data.name,region)
							}else{
								fmt.Printf("%s Can't Distribute in %s",data.name,region)

							}
							goto check_dist

						}else{
							fmt.Println("\n Enter valid Distributer")
							goto check_new_dist
						}

			case 4:
				fmt.Println("Enter Parent Distributer's Name  ")
				name := ""
			    scanner.Scan()
			    name = scanner.Text()
				name = strings.ReplaceAll(strings.ToLower(name), " ", "_")
				// sub := false
				data, ok := DistributerData[strings.ToLower(name)]

				if ok{
					fmt.Println("Enter sub Distributer's Name: ")
					sub_name := ""

				    scanner.Scan()
				    sub_name = scanner.Text()
					fmt.Println("sub_name",sub_name)
					sub_name = strings.ReplaceAll(strings.ToLower(sub_name), " ", "_")

					Value,_ := isSubDistributorOf(data,sub_name)
					if Value{
						fmt.Println("Sub Distributer is already there")

					}else{
						distributer := Distributor{
							name:sub_name,
							permission: Permission{},
							subDistributors: []Distributor{},
						}
						if subDistsByName == nil{
							subDistsByName = make(map[string]string)
						}
						subDistsByName[sub_name] = name
						
						data.subDistributors = append(data.subDistributors,distributer)
						DistributerData[name] = data	

					}
				}else{
					fmt.Println("NO distributer Found")
				}
				goto main_menu
		}

}

func InitPlacesData(){
	DistributerData = make(map[string]Distributor, 0)
	sub_dist1 := Distributor{
									name:"Dist2",
									permission: Permission{},
									subDistributors: []Distributor{},
								}
	distributer := &Distributor{
									name:"Dist1",
									permission: Permission{include:[]string{"kerala-india"},exclude:[]string{"japan"}},
									subDistributors: []Distributor{sub_dist1},
								}

	DistributerData["dist1"] = *distributer

	subDistsByName = make(map[string]string)
	subDistsByName["dist2"] = "dist1"

	file, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Create data structure for quick lookups
	if PlacesMap == nil{
		PlacesMap = make(map[string]Location)
	}
	for _, row := range data {
		location := Location{
			City:    strings.ReplaceAll(strings.ToLower(row[3]), " ", "_"),
			State:   strings.ReplaceAll(strings.ToLower(row[4]), " ", "_"),
			Country: strings.ReplaceAll(strings.ToLower(row[5]), " ", "_"),
		}
		PlacesMap[location.City] = location
		PlacesMap[location.State] = location
		PlacesMap[location.Country] = location
	}
}

func isSubDistributorOf(parentDist Distributor, subDistName string) (bool,Distributor) {
    if parentDist.name == subDistName {
        return true,Distributor{}
    }
    for _, subDist := range parentDist.subDistributors {
        if subDist.name == subDistName {
            return true,subDist
        } 
    }
    return false,Distributor{}
}
