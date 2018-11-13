package distributors

import(
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/gowtham/challenge2016_New/model"
)

type Distributors struct{
    Distributor map[string]*model.Distributor
}

func (d Distributors)Manage_distributors(cities map[string]*model.Country_list)([]string, bool){
    fmt.Print("Enter the distributor name=")
    scanner := bufio.NewScanner(os.Stdin)
    var distributor string
    scanner.Scan()
    distributor = scanner.Text()
    distributors := strings.Split(distributor, "<")
    for index, ele := range distributors{
        distributors[index] = strings.TrimSpace(ele)
    }
    err := d.Update_distributor(d.Distributor, distributors, len(distributors)-1)
    return distributors, err
}


func (d Distributors) Update_distributor(distributors_map map[string]*model.Distributor, distributors []string, index int) bool{
    if index < 0{
        return false
    }
    _, ok := distributors_map[distributors[index]]
    // Parent distributor not found then sub-distributor creation should fail
    if ok == false && index !=0{
        fmt.Println("!!Parent distributor:", distributors[index], "not found")
        return true
    }else if  ok == false && index ==0{
        // Distributor not found, create new one
        distributors_map[distributors[index]] = &model.Distributor{
            Distributor_name: distributors[index],
        }
        return false
    }else if ok == true && index !=0{
        if distributors_map[distributors[index]].Sub_distributor != nil{
            d.Update_distributor(
                distributors_map[distributors[index]].Sub_distributor, distributors, index-1)
            index --
            return false
        }else if (index -1) == 0{
            distributors_map[distributors[index]].Sub_distributor = make(map[string]*model.Distributor)
            d.Update_distributor( distributors_map[distributors[index]].Sub_distributor, distributors, index-1)
            index --
            return false
        }else{
            fmt.Print("!! parent distributor:", distributors[index], " not found")
            return true
        }
    }else if ok == true && index ==0{
        fmt.Println("Note: distributor:", distributors[index], "is already present")   
    }
    return false
}

