package model

type City_list struct{
    City_name string
    City_code string
    Next *City_list
}

type Province_list struct{
    Province_name string
    Province_code string
    Next *Province_list
    City *City_list
}

type Country_list struct {
    Country_name string
    Country_code string
    Next *Country_list
    Province *Province_list
}

type Data struct{
    Country_name string  `json:"Country Name"`
    Country_code string  `json:"Country Code"`
    Province_name string `json:"Province Name"`
    Province_code string `json:"Province Code"`
    City_name string     `json:"City Name"`
    City_code string     `json:"City Code"`
}

type Distributor struct{
    Distributor_name string
    Sub_distributor *Distributor
    Parent_tree *Country_list
    Permission *Country_list
    Next *Distributor
}
