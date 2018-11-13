package model

type City_list struct{
    City_name string
    City_code string
}

type Province_list struct{
    Province_name string
    Province_code string
    City map[string]*City_list
}

type Country_list struct {
    Country_name string
    Country_code string
    Province map[string]*Province_list
}

type Distributor struct{
    Distributor_name string
    Sub_distributor map[string]*Distributor
    Permission map[string]*Country_list
}

type Permission_list struct{
    Permission map[string]*Country_list
}
