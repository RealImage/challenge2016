package util

import(
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/gowtham/challenge2016/model"
)

type Tree struct{
    Country *model.Country_list
}

type Distributor_tree struct{
    distributor *model.Distributor
}


func (root *Tree)Create_tree(cities model.Data){
    if root.Country == nil{
        // creating country list
        root.Country = &model.Country_list{
            Country_name: cities.Country_name,
            Country_code: cities.Country_code,
        }
        // Creating province list
        root.Country.Province = &model.Province_list{
            Province_name: cities.Province_name,
            Province_code: cities.Province_code,
        }
        // Creating citiy list
        root.Country.Province.City = &model.City_list{
            City_name: cities.City_name,
            City_code: cities.City_code,
        }
    }else{
        // Update country list
        ptr_country := root.Update_country(cities)
        // Update province list
        ptr_province := root.Update_province(cities, ptr_country)
        // update city list
        root.Update_city(cities, ptr_province)
    }
}


func (root *Tree)Update_country(cities model.Data)(*model.Country_list){
    ptr_country := root.Country
    for ptr_country.Next != nil{
        if ptr_country.Country_name != cities.Country_name{
            ptr_country = ptr_country.Next
        }else{
            // country already present then break the loop
            break
        }
    }
    if (ptr_country.Country_name != cities.Country_name) && (ptr_country.Next == nil){
        Country := &model.Country_list{
            Country_name: cities.Country_name,
            Country_code: cities.Country_code,
        }
        ptr_country.Next = Country
        ptr_country = ptr_country.Next
    }
    return ptr_country
}


func (root *Tree)Update_province(cities model.Data, ptr_country *model.Country_list)(*model.Province_list){
    var ptr_province *model.Province_list
    Province := &model.Province_list{
        Province_name: cities.Province_name,
        Province_code: cities.Province_code,
    }
    if ptr_country.Province == nil{
        ptr_country.Province = Province
        ptr_province = ptr_country.Province
    }else{
        ptr_province = ptr_country.Province
        for ptr_province.Next != nil{
            if ptr_province.Province_name != cities.Province_name{
                ptr_province = ptr_province.Next
            }else{
                break
            }
        }
        if ptr_province.Province_name != cities.Province_name && ptr_province.Next == nil{
            ptr_province.Next = Province
            ptr_province = ptr_province.Next
        }
    }
    return ptr_province 
}

func (root *Tree)Update_city(cities model.Data, ptr_province *model.Province_list){
    var ptr_city *model.City_list
    City := &model.City_list{
        City_name: cities.City_name,
        City_code: cities.City_code,
    }
    if ptr_province.City == nil{
        ptr_province.City = City
        ptr_city = ptr_province.City
    }else{
        ptr_city = ptr_province.City
        for ptr_city.Next != nil{
            if ptr_city.City_name != cities.City_name{
                ptr_city = ptr_city.Next
            }else{
                break
            }
        }
        if ptr_city.City_name != cities.City_name && ptr_city.Next == nil{
            ptr_city.Next = City
        }
    }
}


func (d *Distributor_tree)Manage_distributor(cities_tree Tree){
    fmt.Print("Enter the distributor name=")
    scanner := bufio.NewScanner(os.Stdin)
    var distributor string
    scanner.Scan()
    distributor = scanner.Text()
    distributors := strings.Split(distributor, "<")
    for index, ele := range distributors{
        distributors[index] = strings.TrimSpace(ele)
    }
    distributor_ptr := d.Update_distributor(distributors)
    if distributor_ptr != nil{
        ch := 0
        for ch != 4{
            fmt.Print("\n1.INCLUDE\n2.EXCLUDE\n3.Check Permission\n4.Main Menu\nEnter your choice=")
            fmt.Scanln(&ch)
            if ch == 1{
                var include string 
                fmt.Print("Enter the include string=")
                scanner.Scan()
                include = scanner.Text()
                include_arr := strings.Split(include, "-")
                    if len(distributors) == 1{
                        // Root level distributor
                        assign_permission(distributor_ptr, cities_tree.Country, include_arr)
                    }else if len(distributors) > 1{
                        // sub distributors
                        assign_permission(distributor_ptr, distributor_ptr.Parent_tree, include_arr)
                    }
            }else if ch == 2{
                var include string
                fmt.Print("Enter the exclude string=")
                scanner.Scan()
                include = scanner.Text()
                include_arr := strings.Split(include, "-")
                remove_premission(distributor_ptr, include_arr)
            }
            if ch == 3{
                var include string
                fmt.Print("Enter the permission to check=")
                scanner.Scan()
                include = scanner.Text()
                include_arr := strings.Split(include, "-")
                check_permission(distributor_ptr, include_arr)
            }
        }
    }
}

func check_permission(distributor_ptr *model.Distributor, include_arr []string){
   if len(include_arr)-1 >=0 {
       // check country
       country_ptr := reference_country_pointer(distributor_ptr.Permission, include_arr[len(include_arr)-1])
       if country_ptr != nil{
           if len(include_arr)-2 >=0 {
               province_ptr := reference_province_pointer(country_ptr.Province, include_arr[len(include_arr)-2])
               if province_ptr != nil{
                   if len(include_arr)-3 >=0 {
                       city_ptr := reference_city_pointer(province_ptr.City, include_arr[len(include_arr)-3])
                       if city_ptr != nil{
                           fmt.Println("*** YES ***")
                       }else{
                           fmt.Println("*** NO ***")
                       }
                   }else{
                       fmt.Println("*** YES ***")
                   }
               }else{
                   fmt.Println("*** NO ***")
               }
           }else{
               fmt.Println("*** YES ***")
           }
       }else{
           fmt.Println("*** NO ***")
       }
       
   }
}

func remove_premission(distributor_ptr *model.Distributor, include_arr []string){
    // Exclude will affects its sub-distributors also
    if len(include_arr) == 1{
        // Remove country pointer reference
        remove_country_pointer(distributor_ptr, include_arr[len(include_arr)-1])
    }else if len(include_arr) == 2{
        // Remove province
        remove_province_pointer(distributor_ptr, include_arr)
    }else if len(include_arr) == 3{
        // remove city
        remove_city_pointer(distributor_ptr, include_arr)
    }
}

func remove_city_pointer(distributor_ptr *model.Distributor, include_arr []string){
    if distributor_ptr == nil{
        return
    }
    remove_city_pointer(distributor_ptr.Sub_distributor, include_arr)
    // find country
    country_ptr := reference_country_pointer(distributor_ptr.Permission, include_arr[len(include_arr)-1])
    if country_ptr != nil{
        province_ptr := reference_province_pointer(country_ptr.Province, include_arr[len(include_arr)-2])
        if province_ptr != nil{
            ptr := province_ptr.City
            prev := province_ptr.City
            for ptr != nil{
                if strings.ToLower(ptr.City_name) == strings.ToLower(include_arr[len(include_arr)-3]){
                    if ptr == prev{
                        // remove header
                        province_ptr.City = ptr.Next
                    }else{
                        prev.Next = ptr.Next
                    }
                    // if all cities are removed then remove provision also
                    if province_ptr.City == nil{
                        remove_province_pointer(distributor_ptr, include_arr)
                    }
                    return
                }
                prev = ptr
                ptr = ptr.Next
            }
            if ptr == nil{
                fmt.Println("!! City name not present in a tree to exclude")
            }
        }else{
            fmt.Println("!! Province name not present in a tree")
        }
    }else{
        fmt.Println("!! Country name not present in a tree")
    }
}

func remove_province_pointer(distributor_ptr *model.Distributor, include_arr []string){
    if distributor_ptr == nil{
        return
    }
    remove_province_pointer(distributor_ptr.Sub_distributor, include_arr)
    // find country pointer
    country_ptr := reference_country_pointer(distributor_ptr.Permission, include_arr[len(include_arr)-1])
    if country_ptr != nil{
        ptr := country_ptr.Province
        prev := country_ptr.Province
        for ptr != nil{
            if strings.ToLower(ptr.Province_name) == strings.ToLower(include_arr[len(include_arr)-2]){
                if ptr == prev{
                    // remove header
                    country_ptr.Province = ptr.Next
                }else{
                    prev.Next = ptr.Next
                }
                // if all province are removed then remove country also
                if country_ptr.Province == nil{
                    remove_country_pointer(distributor_ptr, include_arr[len(include_arr)-1])
                }
                return
            }
            prev = ptr
            ptr = ptr.Next
        }
        if ptr == nil{
            fmt.Println("!! Provision name not present in a tree to exclude")
        }
    }else{
        fmt.Println("!! Country name not present in a tree")
    }
}

func remove_country_pointer(distributor_ptr *model.Distributor, country_name string){
    if distributor_ptr == nil{
        return
    }
    remove_country_pointer(distributor_ptr.Sub_distributor, country_name)
    ptr := distributor_ptr.Permission
    prev := distributor_ptr.Permission
    for ptr != nil{
        if strings.ToLower(ptr.Country_name) == strings.ToLower(country_name){
            if ptr == prev{
                // remove header
                distributor_ptr.Permission = ptr.Next
            }else{
                prev.Next = ptr.Next
            }
            return
        }
        prev = ptr
        ptr = ptr.Next
    }
    if ptr == nil{
        fmt.Println("!! Country name not present in a tree to exclude")
    }
}


func assign_permission(distributor_ptr *model.Distributor, country_tree *model.Country_list, include_arr []string){
    if len(include_arr) == 1{
        country_ptr :=  reference_country_pointer(country_tree, include_arr[len(include_arr) -1])
        if country_ptr == nil{
            fmt.Println("!! Given country name not present in a tree")
            return
        }
        if distributor_ptr.Permission == nil{
            distributor_ptr.Permission = &model.Country_list{
                Country_name: country_ptr.Country_name,
                Country_code: country_ptr.Country_code,
            }
        }else{
            ptr := distributor_ptr.Permission
            for ptr != nil{
                if strings.ToLower(ptr.Country_name) == strings.ToLower(include_arr[len(include_arr) -1]){
                    fmt.Println("!! Given country is already included for this distributor")
                    return
                }
                ptr = ptr.Next
            }
            if ptr == nil{
                country := &model.Country_list{
                    Country_name: country_ptr.Country_name,
                    Country_code: country_ptr.Country_code,
                }
                country.Next = distributor_ptr.Permission
                distributor_ptr.Permission = country
            }
        }
        ptr := country_ptr.Province
        for ptr != nil{
            if distributor_ptr.Permission.Province == nil{
                distributor_ptr.Permission.Province = &model.Province_list{
                    Province_name: ptr.Province_name,
                    Province_code: ptr.Province_code,
                }
            }else{
                province_ptr := &model.Province_list{
                    Province_name: ptr.Province_name,
                    Province_code: ptr.Province_code,
                }
                province_ptr.Next = distributor_ptr.Permission.Province
                distributor_ptr.Permission.Province = province_ptr
            }
            ptr1 := ptr.City
            for ptr1 != nil{
                if distributor_ptr.Permission.Province.City == nil{
                    distributor_ptr.Permission.Province.City = &model.City_list{
                        City_name: ptr1.City_name,
                        City_code: ptr1.City_code,
                    }
                }else{
                    city_ptr := &model.City_list{
                        City_name: ptr1.City_name,
                        City_code: ptr1.City_code,
                    }
                    city_ptr.Next = distributor_ptr.Permission.Province.City
                    distributor_ptr.Permission.Province.City = city_ptr
                }
                ptr1 = ptr1.Next
            }
            ptr = ptr.Next
            
        }
    }else if len(include_arr) == 2{
        country_ptr :=  reference_country_pointer(country_tree, include_arr[len(include_arr) -1])
        if country_ptr == nil{
            fmt.Println("!! Given country name not present in a tree")
            return
        }
        init := distributor_ptr.Permission
        if distributor_ptr.Permission == nil{
            distributor_ptr.Permission = &model.Country_list{
                Country_name: country_ptr.Country_name,
                Country_code: country_ptr.Country_code,
            }
        }else{
            ptr := distributor_ptr.Permission
            prev := distributor_ptr.Permission
            for ptr != nil{
                if strings.ToLower(ptr.Country_name) != strings.ToLower(include_arr[len(include_arr) -1]){
                    prev = ptr
                    ptr = ptr.Next
                }else{
                   break
                }
            }
            if ptr == nil{
                country := &model.Country_list{
                    Country_name: country_ptr.Country_name,
                    Country_code: country_ptr.Country_code,
                }
                country.Next = distributor_ptr.Permission
                distributor_ptr.Permission = country
            }else if ptr!= prev{
                prev.Next = ptr.Next
                ptr.Next = distributor_ptr.Permission
                distributor_ptr.Permission = ptr
                init = distributor_ptr.Permission
            }
        }
        province_ptr :=  reference_province_pointer(country_ptr.Province, include_arr[len(include_arr) -2])
        if province_ptr == nil{
            fmt.Println("!! Given province name not present in country or it is excluded")
            distributor_ptr.Permission = init
            return
        }
        if distributor_ptr.Permission.Province == nil{
            distributor_ptr.Permission.Province = &model.Province_list{
                Province_name: province_ptr.Province_name,
                Province_code: province_ptr.Province_code,
            }
        }else{
            ptr := distributor_ptr.Permission.Province
            prev := distributor_ptr.Permission.Province
            for ptr != nil{
                if strings.ToLower(ptr.Province_name) != strings.ToLower(include_arr[len(include_arr) -2]){
                    prev = ptr
                    ptr = ptr.Next
                }else{
                    break
                }
            }
            if ptr == nil{
                province := &model.Province_list{
                    Province_name: province_ptr.Province_name,
                    Province_code: province_ptr.Province_code,
                }
                province.Next = distributor_ptr.Permission.Province
                distributor_ptr.Permission.Province = province
            }else if ptr != prev{
                prev.Next = ptr.Next
                ptr.Next = distributor_ptr.Permission.Province
                distributor_ptr.Permission.Province = ptr
            }
        } 
        city_ptr := province_ptr.City
        for city_ptr != nil{
            if distributor_ptr.Permission.Province.City == nil{
                distributor_ptr.Permission.Province.City = &model.City_list{
                    City_name: city_ptr.City_name,
                    City_code: city_ptr.City_code,
                }
            }else{
                city := &model.City_list{
                    City_name: city_ptr.City_name,
                    City_code: city_ptr.City_code,
                }
                city.Next = distributor_ptr.Permission.Province.City
                distributor_ptr.Permission.Province.City = city
            }
            city_ptr = city_ptr.Next
        }
    }else if len(include_arr) == 3{
        country_ptr :=  reference_country_pointer(country_tree, include_arr[len(include_arr) -1])
        if country_ptr == nil{
            fmt.Println("!! Given country name not present in a tree")
            return
        }
        country_init := distributor_ptr.Permission
        if distributor_ptr.Permission == nil{
            distributor_ptr.Permission = &model.Country_list{
                Country_name: country_ptr.Country_name,
                Country_code: country_ptr.Country_code,
            }
        }else{
            ptr := distributor_ptr.Permission
            prev := distributor_ptr.Permission
            for ptr != nil{
                if strings.ToLower(ptr.Country_name) != strings.ToLower(include_arr[len(include_arr) -1]){
                    prev = ptr
                    ptr = ptr.Next
                }else{
                   break
                }
            }
            if ptr == nil{
                country := &model.Country_list{
                    Country_name: country_ptr.Country_name,
                    Country_code: country_ptr.Country_code,
                }
                country.Next = distributor_ptr.Permission
                distributor_ptr.Permission = country
            }else if ptr!= prev{
                prev.Next = ptr.Next
                ptr.Next = distributor_ptr.Permission
                distributor_ptr.Permission = ptr
                country_init = distributor_ptr.Permission
            }
        }
        province_ptr :=  reference_province_pointer(country_ptr.Province, include_arr[len(include_arr) -2])
        if province_ptr == nil{
            fmt.Println("!! Given province name not present in country or it is excluded")
            distributor_ptr.Permission = country_init
            return
        }
        province_init := distributor_ptr.Permission.Province
        if distributor_ptr.Permission.Province == nil{
            distributor_ptr.Permission.Province = &model.Province_list{
                Province_name: province_ptr.Province_name,
                Province_code: province_ptr.Province_code,
            }
        }else{
            ptr := distributor_ptr.Permission.Province
            prev := distributor_ptr.Permission.Province
            for ptr != nil{
                if strings.ToLower(ptr.Province_name) != strings.ToLower(include_arr[len(include_arr) -2]){
                    prev = ptr
                    ptr = ptr.Next
                }else{
                    break
                }
            }
            if ptr == nil{
                province := &model.Province_list{
                    Province_name: province_ptr.Province_name,
                    Province_code: province_ptr.Province_code,
                }
                province.Next = distributor_ptr.Permission.Province
                distributor_ptr.Permission.Province = province
            }else if ptr != prev{
                prev.Next = ptr.Next
                ptr.Next = distributor_ptr.Permission.Province
                distributor_ptr.Permission.Province = ptr
                province_init = distributor_ptr.Permission.Province
            }
        }
        city_ptr :=  reference_city_pointer(province_ptr.City, include_arr[len(include_arr) -3])
        if city_ptr == nil{
            fmt.Println("!!Given city name not present under province")
            distributor_ptr.Permission.Province = province_init
            distributor_ptr.Permission = country_init
            return
        }
        if distributor_ptr.Permission.Province.City == nil{
            distributor_ptr.Permission.Province.City = &model.City_list{
                City_name: city_ptr.City_name,
                City_code: city_ptr.City_code,
            }
        }else{
            ptr := distributor_ptr.Permission.Province.City
            for ptr != nil{
                if strings.ToLower(ptr.City_name) != strings.ToLower(include_arr[len(include_arr) -3]){
                    ptr = ptr.Next
                }else{
                    break
                }
            }
            if ptr == nil{
                city := &model.City_list{
                    City_name: city_ptr.City_name,
                    City_code: city_ptr.City_code,
                }
                city.Next = distributor_ptr.Permission.Province.City
                distributor_ptr.Permission.Province.City = city
            }
        }
    }

}

func reference_city_pointer(city_tree *model.City_list, city string)(*model.City_list){
    ptr := city_tree
    for ptr != nil{
        if strings.ToLower(ptr.City_name) != strings.ToLower(city){
            ptr = ptr.Next
        }else{
            break
        }
    }
    if ptr!= nil{
        return ptr
    }else{
        return nil
    }
}

func reference_province_pointer(province_tree *model.Province_list, province string)(*model.Province_list){
    ptr := province_tree
    for ptr != nil{
        if strings.ToLower(ptr.Province_name) != strings.ToLower(province){
            ptr = ptr.Next
        }else{
            break
        }
    }
    if ptr!= nil{
        return ptr
    }else{
        return nil
    }
}

func reference_country_pointer(country_tree *model.Country_list, country string)(*model.Country_list){
    ptr := country_tree
    for ptr != nil{
        if strings.ToLower(ptr.Country_name) != strings.ToLower(country){
            ptr = ptr.Next
        }else{
            break
        }
    }
    if ptr!= nil{
        return ptr
    }else{
        return nil
    }
    
}

func (d *Distributor_tree)Update_distributor(distributors []string)(*model.Distributor){
    // Check distributor found and hierarchy is correct
    ptr := d.distributor
    if ptr == nil{
        // Parent distributor should present already,
        // then only it will allow to add sub distributors
        if len(distributors) == 1{
            // distributor not in list and no sub distributor given
            d.distributor = &model.Distributor{
                Distributor_name: distributors[0], 
            }
            ptr = d.distributor
        }else{
            fmt.Println("Create parent distributor: ",distributors[len(distributors)-1], " first then only it will alow to create sub-distributors")
        }
    }else{
        distributor, err := update_sub_distributor(ptr, len(distributors)-1, distributors)
        if distributor != nil && err == false{
            ptr = distributor
        }else{
            ptr = nil
        }
    }
    return ptr
}

func update_sub_distributor(ptr *model.Distributor, index int, distributors[] string)(*model.Distributor, bool){
    err := false
    if ptr == nil{
        return nil, err
    }
    for ptr.Next != nil{
        if ptr.Distributor_name == distributors[index]{
            break
        }
        ptr = ptr.Next
    }
    if ptr.Distributor_name == distributors[index]{
        index --
        // if parent found for sub-distributor then check it is the
        // immediate parent or not 
        if index == 0 {
            // immediate parent found
            if ptr.Sub_distributor == nil{
                ptr.Sub_distributor = &model.Distributor{
                    Distributor_name: distributors[index],
                    Parent_tree: ptr.Permission,
                }
                ptr = ptr.Sub_distributor
            }else{
                sub_dis := ptr.Sub_distributor
                for sub_dis!= nil{
                    if sub_dis.Distributor_name == distributors[index]{
                        fmt.Println("Note: sub distributor:", distributors[index], "already present")
                        return sub_dis, err
                    }
                    sub_dis = sub_dis.Next
                }
                Sub_distributor := &model.Distributor{
                    Distributor_name: distributors[index],
                    Parent_tree: ptr.Permission,
                }
                Sub_distributor.Next = ptr.Sub_distributor
                ptr.Sub_distributor = Sub_distributor
                ptr = ptr.Sub_distributor
            }
        }else{ 
            // Need to finds its immediate parent, it will call all of its parents recursively
            // it is keep call this function until immediate parent found
            // if immediate parent found then assign sub distributor under it
            if index > 0{
                ptr, err = update_sub_distributor(ptr.Sub_distributor, index, distributors)
            }else{
                fmt.Println("Note: Distributor", ptr.Distributor_name, "is already present")
            }
        }
    }else if ptr.Next == nil && index == 0{
        // Always root level parent creation will happen here
        ptr.Next = &model.Distributor{
            Distributor_name: distributors[index],
        }
        ptr = ptr.Next
    }else{
        fmt.Println("Before sub-distributor: ", distributors[index], "parent distirbutor should be added")
        err = true
    }
    return ptr, err 
}
