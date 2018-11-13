
package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
    "io"
    "log"
    "strings"

    "github.com/gowtham/challenge2016_New/distributors"
    "github.com/gowtham/challenge2016_New/model"
    "github.com/gowtham/challenge2016_New/permissions"
)

func main() {
    // Find relative path
    pwd, _ := os.Getwd()
    data, _ := os.Open(pwd + "/cities.csv")
    reader := csv.NewReader(bufio.NewReader(data))
    count := 0
    // Parent list
    var cities = map[string]*model.Country_list{}
    for {
        line, error := reader.Read()
        // avoid header from csv
        if count == 0 {
            count ++
            continue
        }
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        City_code :=  strings.ToLower(strings.Replace(line[0], " ", "",-1))
        Province_code := strings.ToLower(strings.Replace(line[1], " ", "",-1))
        Country_code := strings.ToLower(strings.Replace(line[2], " ", "",-1))
        City_name := strings.ToLower(strings.Replace(line[3], " ", "",-1))
        Province_name := strings.ToLower(strings.Replace(line[4], " ", "",-1))
        Country_name := strings.ToLower(strings.Replace(line[5], " ", "",-1))
        // create country key in map 
        _, ok := cities[Country_name]
        if ok == false{
            cities[Country_name] = &model.Country_list{
                Country_name: Country_name,
                Country_code: Country_code,
            
            }
        }
        _, ok = cities[Country_name].Province[Province_name]
        if ok == false {
            // append new province
            if cities[Country_name].Province == nil {
                cities[Country_name].Province = make(map[string]*model.Province_list)
            }
            cities[Country_name].Province[Province_name] = &model.Province_list{
                Province_name: Province_name,
                Province_code: Province_code,
            }
        }
        _, ok= cities[Country_name].Province[Province_name].City[City_name]
        if ok == false {
            // append new city
            if cities[Country_name].Province[Province_name].City == nil{
                cities[Country_name].Province[Province_name].City = make(map[string]*model.City_list)
            }
            cities[Country_name].Province[Province_name].City[City_name]= &model.City_list{
                City_name: City_name,
                City_code: City_code,
            }
        }
        
    }
    ch := 0
    distributor := distributors.Distributors{}
    permission_hash := permissions.Permission{}
    distributor.Distributor = make(map[string]*model.Distributor)
    for{
        fmt.Print("\n1.Distributor(Create or switch control)\n2.exit\nEnter your choice=")
        fmt.Scanln(&ch)
        if ch != 2{
            // Create distributos of update distributor
            distributors_arr, err := distributor.Manage_distributors(cities)
            if err != true{
                permission_hash.Assign_permission(distributor.Distributor, distributors_arr, cities)
            }
        }else{
            fmt.Println("!!Exit")
            break
        }
    }
    
}
