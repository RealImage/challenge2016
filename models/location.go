package models

type Country struct{
	Name string
	Code string
}

type Province struct{
	Name string
	Code string
	Country *Country
}

type City struct{
	Name string
	Code string
	Province *Province
}