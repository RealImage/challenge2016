package models

type Cities struct {
	City_Code     string
	Province_Code string
	Country_Code  string
	City_Name     string
	Province_Name string
	Country_Name  string
}

type Distributors struct {
	Id              int
	Distributorname string
	Included        string
	Excluded        string
	Subdistributor  string
	Seniority       string
}

type Permissions struct {
	Included string
	Excluded string
}

type Data struct {
	Distributorname string
	Permission      string
}
