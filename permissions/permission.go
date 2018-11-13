package permissions

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/gowtham/challenge2016_New/model"
)


type Permission struct{
    permission map[string]*model.Permission_list
}



func (p Permission) Assign_permission(Distributor_map map[string]*model.Distributor, distributors_arr []string, cities map[string]*model.Country_list){
    ch := 0
    for ch != 4{
        fmt.Print("\n1.INCLUDE\n2.EXCLUDE\n3.Check Permission\n4.Main Menu\nEnter your choice=")
        fmt.Scanln(&ch)
        scanner := bufio.NewScanner(os.Stdin)
        if ch == 1{
            var include string
            fmt.Print("Enter the include string=")
            scanner.Scan()
            include = strings.ToLower(scanner.Text())
            include_arr := strings.Split(include, "-")
            p.include_permission(Distributor_map, distributors_arr, include_arr, cities)
        }else if ch == 2{
            fmt.Print("Enter the exclude string=")
            scanner.Scan()
            exclude := strings.ToLower(scanner.Text())
            exclude_arr := strings.Split(exclude, "-")
            p.exclude_permission(Distributor_map, distributors_arr, exclude_arr)
        }else if ch == 3{
           var check_perm string
           fmt.Print("Enter the permission to check=")
           scanner.Scan()
           check_perm = strings.ToLower(scanner.Text())
           check_perm_arr := strings.Split(check_perm, "-")
           p.check_permission(Distributor_map, check_perm_arr, distributors_arr)
       }
    }

}


func (p Permission) check_permission(Distributor_map map[string]*model.Distributor, check_perm []string, distributors []string){
    for i := len(distributors)-1; i>=0; i--{
        _,ok := Distributor_map[distributors[i]]
        if ok{
          if i==0{
            // distributor found
            if len(check_perm) == 1{
                // compare cluntry name
                _, ok := Distributor_map[distributors[i]].Permission[check_perm[0]]
                if ok{
                    fmt.Println("*** YES ****")
                }else{
                    fmt.Println("*** NO ****")
                }
            }else if len(check_perm) == 2{
                // compare cluntry name and province name
                _, ok := Distributor_map[distributors[i]].Permission[check_perm[1]].Province[check_perm[0]]
                if ok{
                    fmt.Println("*** YES ****")
                }else{
                    fmt.Println("*** NO ****")
                }
            }else if len(check_perm) == 3{
                // compare cluntry name and province name
                _, ok := Distributor_map[distributors[i]].Permission[check_perm[2]].Province[check_perm[1]].City[check_perm[0]]
                if ok{
                    fmt.Println("*** YES ****")
                }else{
                    fmt.Println("*** NO ****")
                }
            }
          }
          Distributor_map = Distributor_map[distributors[i]].Sub_distributor
        }else{
            fmt.Println("!!Distributor:",distributors[i],"is not found")
        }
    }

}

func (p Permission)exclude_permission(Distributor_map map[string]*model.Distributor, distributors []string, exclude_arr []string){
    if len(exclude_arr) == 1{
        // remove country
        // exclude will affect sud-distributors also
        p.exclude_country(Distributor_map, distributors, exclude_arr[0])
    }else if len(exclude_arr) == 2{
        // remove province
        p.exclude_province(Distributor_map, exclude_arr, distributors)
    }else if len(exclude_arr) == 3{
        p.exclude_city(Distributor_map, exclude_arr, distributors)
   }
}


func (p Permission)exclude_subdistributor_city(Distributor_map map[string]*model.Distributor, exclude_arr []string){
    if Distributor_map == nil {
        return
    }
    for key,_ := range Distributor_map{
        // exclude from sub-distributors also
        delete(Distributor_map[key].Permission[exclude_arr[2]].Province[exclude_arr[1]].City, exclude_arr[0])
        p.exclude_subdistributor_city(Distributor_map[key].Sub_distributor, exclude_arr)
    }
}


func (p Permission)exclude_city(Distributor_map map[string]*model.Distributor, exclude_arr []string, distributors []string){
    for i := len(distributors)-1; i>=0; i--{
        _,ok := Distributor_map[distributors[i]]
        if ok{
          if i==0{
              delete(Distributor_map[distributors[i]].Permission[exclude_arr[2]].Province[exclude_arr[1]].City, exclude_arr[0])
          }
          Distributor_map = Distributor_map[distributors[i]].Sub_distributor

        }else{
            fmt.Println("!!Distributor:",distributors[i],"is not found")
        }
    }
    p.exclude_subdistributor_city(Distributor_map, exclude_arr)
}

func (p Permission)exclude_subdistributor_province(Distributor_map map[string]*model.Distributor, exclude_arr []string){
    if Distributor_map == nil {
        return
    }
    for key,_ := range Distributor_map{
        // exclude from sub-distributors also
        delete(Distributor_map[key].Permission[exclude_arr[1]].Province, exclude_arr[0])
        p.exclude_subdistributor_province(Distributor_map[key].Sub_distributor, exclude_arr)
    }
}

func (p Permission)exclude_province(Distributor_map map[string]*model.Distributor, exclude_arr []string, distributors []string){
    for i := len(distributors)-1; i>=0; i--{
        _,ok := Distributor_map[distributors[i]]
        if ok{
          if i==0{
              delete(Distributor_map[distributors[i]].Permission[exclude_arr[1]].Province, exclude_arr[0])
              break
          }
          Distributor_map = Distributor_map[distributors[i]].Sub_distributor

        }else{
            fmt.Println("!!Distributor:",distributors[i],"is not found")
        }
    }
    p.exclude_subdistributor_province(Distributor_map, exclude_arr)
}


func (p Permission)exclude_distributor(Distributor_map map[string]*model.Distributor, exclude_str string){
    if Distributor_map == nil {
        return
    }
    for key,_ := range Distributor_map{
        // exclude from sub-distributors also
        delete(Distributor_map[key].Permission, exclude_str)
        p.exclude_distributor(Distributor_map[key].Sub_distributor, exclude_str)
    }
}

func (p Permission)exclude_country(Distributor_map map[string]*model.Distributor, distributors []string, exclude_str string){
    for i := len(distributors)-1; i>=0; i--{
        _,ok := Distributor_map[distributors[i]]
        if ok{
          if i==0{
              delete(Distributor_map[distributors[i]].Permission, exclude_str)
          } 
          Distributor_map = Distributor_map[distributors[i]].Sub_distributor
           
        }else{
            fmt.Println("!!Distributor:",distributors[i],"is not found")
        }
    }
    p.exclude_distributor(Distributor_map, exclude_str)

}


func (p Permission)include_permission(Distributor_map map[string]*model.Distributor, distributors_arr []string, include_arr []string, cities map[string]*model.Country_list){
    var parent_permission = map[string]*model.Country_list{}
    var distributor *model.Distributor
    parent_permission = make(map[string]*model.Country_list)
    if len(distributors_arr) == 1{
        parent_permission =  cities
        distributor = Distributor_map[distributors_arr[0]]
    }else{
        for i:=len(distributors_arr)-1; i>=1; i--{
            _, ok := Distributor_map[distributors_arr[i]]
            if ok{
                // need to assign sub-distributor permission using its immediate parent permission
                parent_permission = Distributor_map[distributors_arr[i]].Permission
                Distributor_map = Distributor_map[distributors_arr[i]].Sub_distributor

            }else{
                fmt.Println("!!Distributor ", distributors_arr[i], "is not present to assign permission")
                return
            }
        }
    }
    distributor = Distributor_map[distributors_arr[0]]
    if distributor.Permission == nil{
        distributor.Permission = make(map[string]*model.Country_list)
    }
    if len(include_arr) == 1{
        // add permission for specific country
        // check country parent in parent permission
        _, ok := parent_permission[include_arr[0]]
        if ok{
            // copy country
            _, ok := distributor.Permission[include_arr[0]]
            if ok==false{
                if distributor.Permission == nil{
                    distributor.Permission = make(map[string]*model.Country_list)
                }
                distributor.Permission[include_arr[0]] = &model.Country_list{
                    Country_code: parent_permission[include_arr[0]].Country_code,
                    Country_name: parent_permission[include_arr[0]].Country_name,
                }
            }
            // add province
            for key, p := range parent_permission[include_arr[0]].Province{
                if distributor.Permission[include_arr[0]].Province == nil{
                    distributor.Permission[include_arr[0]].Province = make(map[string]*model.Province_list)
                }
                distributor.Permission[include_arr[0]].Province[key] = &model.Province_list{
                    Province_name: p.Province_name,
                    Province_code: p.Province_code,
                }
                for k, c := range parent_permission[include_arr[0]].Province[key].City{
                    if distributor.Permission[include_arr[0]].Province[key].City == nil{
                        distributor.Permission[include_arr[0]].Province[key].City = make(map[string]*model.City_list)
                    }
                    distributor.Permission[include_arr[0]].Province[key].City[k] = &model.City_list{
                        City_name: c.City_name,
                        City_code: c.City_code,
                    }
                }
            }
            
        }else{
            fmt.Println("!!Given country not present in its parent list")
        }
       
    }else if len(include_arr) == 2{
        _, ok := parent_permission[include_arr[1]]
        if ok{
            _, ok := distributor.Permission[include_arr[1]]
            if ok==false{
                if distributor.Permission == nil{
                    distributor.Permission = make(map[string]*model.Country_list)
                }
                distributor.Permission[include_arr[1]] = &model.Country_list{
                    Country_code: parent_permission[include_arr[1]].Country_code,
                    Country_name: parent_permission[include_arr[1]].Country_name,
                }
            }
            // add province
            _, ok = parent_permission[include_arr[1]].Province[include_arr[0]]
            if ok{
                _, ok = distributor.Permission[include_arr[1]].Province[include_arr[0]]
               if ok == false{
                   if distributor.Permission[include_arr[1]].Province == nil{
                       distributor.Permission[include_arr[1]].Province = make(map[string]*model.Province_list)
                   }
                   distributor.Permission[include_arr[1]].Province[include_arr[0]] = &model.Province_list{
                       Province_name: parent_permission[include_arr[1]].Province[include_arr[0]].Province_name,
                       Province_code: parent_permission[include_arr[1]].Province[include_arr[0]].Province_code,
                   }
               }
               for k, c := range parent_permission[include_arr[1]].Province[include_arr[0]].City{
                   if distributor.Permission[include_arr[1]].Province[include_arr[0]].City == nil{
                       distributor.Permission[include_arr[1]].Province[include_arr[0]].City = make(map[string]*model.City_list)
                   }   
                   distributor.Permission[include_arr[1]].Province[include_arr[0]].City[k] = &model.City_list{
                        City_name: c.City_name,
                        City_code: c.City_code,
                    }
               }
                   
            }else{
                fmt.Println("!!Given Province not present in its parent list")
            }
        }else{
            fmt.Println("!!Given country not present in its parent list")
        }
            
    }else if len(include_arr) == 3{
        _, ok := parent_permission[include_arr[2]]
        if ok{
            _, ok := distributor.Permission[include_arr[2]]
            if ok==false{
                if distributor.Permission == nil{
                    distributor.Permission = make(map[string]*model.Country_list)
                }
                distributor.Permission[include_arr[2]] = &model.Country_list{
                    Country_code: parent_permission[include_arr[2]].Country_code,
                    Country_name: parent_permission[include_arr[2]].Country_name,
                }
            }
            // add province
            _, ok = parent_permission[include_arr[2]].Province[include_arr[1]]
            if ok{
                _, ok = distributor.Permission[include_arr[2]].Province[include_arr[1]]
               if ok == false{
                   if distributor.Permission[include_arr[2]].Province == nil{
                       distributor.Permission[include_arr[2]].Province = make(map[string]*model.Province_list)
                   }
                   distributor.Permission[include_arr[2]].Province[include_arr[1]] = &model.Province_list{
                       Province_name: parent_permission[include_arr[2]].Province[include_arr[1]].Province_name,
                       Province_code: parent_permission[include_arr[2]].Province[include_arr[1]].Province_code,
                   }
               }
               // add city
               _, ok = parent_permission[include_arr[2]].Province[include_arr[1]].City[include_arr[0]]
               if ok{
                    _, ok = distributor.Permission[include_arr[2]].Province[include_arr[1]].City[include_arr[0]]
                   if ok == false{
                       if distributor.Permission[include_arr[2]].Province[include_arr[1]].City == nil{
                           distributor.Permission[include_arr[2]].Province[include_arr[1]].City = make(map[string]*model.City_list)
                       }
                       distributor.Permission[include_arr[2]].Province[include_arr[1]].City[include_arr[0]] = &model.City_list{
                           City_name: parent_permission[include_arr[2]].Province[include_arr[1]].City[include_arr[0]].City_name,
                           City_code: parent_permission[include_arr[2]].Province[include_arr[1]].City[include_arr[0]].City_code,
                       }
                   }
               }else{
                   fmt.Println("!!Given city not present in its parent list")
              }
                
            }else{
                fmt.Println("!!Given Province not present in its parent list")
            }
        }else{
            fmt.Println("!!Given country not present in its parent list")
        }
    }
}

