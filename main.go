package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
    "io"
    "log"
    "strings"

    "github.com/gowtham/challenge2016/model"
    "github.com/gowtham/challenge2016/util" 
)

func main() {
    // Find relative path
    pwd, _ := os.Getwd()
    data, _ := os.Open(pwd + "/cities.csv")
    reader := csv.NewReader(bufio.NewReader(data))
    root := util.Tree{}
    count := 0
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
        cities := model.Data{
            City_code: strings.Replace(line[0], " ", "",-1),
            Province_code: strings.Replace(line[1], " ", "",-1),
            Country_code: strings.Replace(line[2], " ", "",-1),
            City_name: strings.Replace(line[3], " ", "",-1),
            Province_name: strings.Replace(line[4], " ", "",-1),
            Country_name: strings.Replace(line[5], " ", "",-1),
        }
        root.Create_tree(cities)
    }
    ch := 0
    distributor := util.Distributor_tree{}
    for{
        fmt.Print("\n1.Distributor\n2.exit\nEnter your choice=")
        fmt.Scanln(&ch)
        if ch != 2{
            distributor.Manage_distributor(root)
        }else{
            break
        }
    }
    
}
