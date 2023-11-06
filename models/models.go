package models

import (
	"strings"
	"log"
)

type Distributor struct {
	Name   string
	Cities []string
	States []string
	Countries []string
	ExCities []string
	ExStates []string
	ExCountries []string
	Addedby string

}

type DistributorInput struct {
	Name       string   `json:"name"`
	Cities     []string `json:"cities"`
	States     []string `json:"states"`
	Countries   []string `json:"countries"`
	ExCities   []string `json:"exCities"`
	ExStates   []string `json:"exStates"`
	ExCountries []string `json:"exCountries"`
	Addedby		string 	`json:"addedby"`

	
}


var DistributerList = []Distributor{
	{
		Name:      "Aman001",
		Cities:    []string{"PUNCH-JK-IN", "KLRAI-TN-IN" , "PLMYR-NY-US"},
		States:    []string{"TN-IN","AZ-US"},
		Countries: []string{"IN","US"},
		ExCities: []string{},
		ExStates: []string{},
		ExCountries: []string{},
		Addedby: "Admin",
	},
	{
		Name:      "JohnDoe",
		Cities:    []string{"WLAMS-AZ-US","ADUCE-CA-US"},
		States:    []string{"RP-GE","UP-IN"},
		Countries: []string{"US"},
		Addedby: "Admin",

	},
	{
		Name:      "Kamlesh001",
		Cities:    []string{"PLMYR-NY-US"},
		States:    []string{},
		Countries: []string{"IN","US"},
		ExCities: []string{},
		ExStates: []string{},
		ExCountries: []string{"TN-IN"},
		Addedby: "Admin",
	},
}
func Getdistributer(name string) Distributor{
	var s Distributor
	for _,v := range DistributerList{
		if v.Name==name{
			return v
		}
	}
	return s 
}

func CountWords(location string) int {
	// Split the location string by commas.
	words := strings.Split(location, "-")

	// Return the number of words in the location.
	return len(words)
}

func containsString(slice []string, str string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}

func IsPermitted(d Distributor,location string) bool{
	status:=false
	n:=CountWords(location)
	if Excluded(d,n,location){
		return false
	}
	if n==1{
		//check in country
			if containsString(d.Countries, location) {
				return true
			}		
	}else if n==2{
		//check in state
			if containsString(d.States, location) {
				return true
			}else{
				parts := strings.Split(location, "-")
				if containsString(d.Countries, parts[1]){
					return true
				}
			}	
		
	}else if n==3{
		//check in city
		if containsString(d.Cities, location) {
			return true
		}else{
			parts := strings.SplitN(location, "-", 2)
			if containsString(d.States, parts[1]){
				return true
			}else{
				parts2 := strings.SplitN(parts[1], "-", 2)
				if containsString(d.Countries, parts2[1]){
					return true
				}
			}
		}	
	}else {
		log.Panicln("Not a valid location")
	}
	return status
}



func Excluded(d Distributor, n int,location string) bool{
	if n==1{
		//check if country is excluded
			if containsString(d.ExCountries, location) {
				return true
			}		
	}else if n==2{
		//check in state
			if containsString(d.ExStates, location) {
				return true
			}else{
				parts := strings.Split(location, "-")
				if containsString(d.ExCountries, parts[1]){
					return true
				}
			}	
		
	}else if n==3{
		//check in city
		if containsString(d.ExCities, location) {
			return true
		}else{
			parts := strings.SplitN(location, "-", 2)
			if containsString(d.ExStates, parts[1]){
				return true
			}else{
				parts2 := strings.SplitN(parts[1], "-", 2)
				if containsString(d.ExCountries, parts2[1]){
					return true
				}
			}
		}	
	}else {
		log.Panicln("Not a valid location")
	}
	return false
}




func CreateDistributor(d Distributor) bool{

	alllocation:=append(d.Cities,d.States...)
	alllocation=append(alllocation,d.Countries...)
	flag:=true

	//Add a func to get the diributer details by name
	if(d.Addedby!="Admin"){
		addDist:=Getdistributer(d.Addedby)
		
		// Add the new distributor to the DistributorList
		for _,v:= range alllocation{
			if IsPermitted(addDist,v){
		}else{
			flag=false
			break
		}
	  }
	   }
	   return flag
}
