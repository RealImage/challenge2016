package main

import "sync"

type country struct {
	Code      string      `json:"country_code,omitempty"`
	Name      string      `json:"country_name,omitempty"`
	Provinces []*province `json:"provinces,omitempty"`
}

type province struct {
	Code   string  `json:"province_code,omitempty"`
	Name   string  `json:"province_name,omitempty"`
	Cities []*city `json:"cities,omitempty"`
}

type city struct {
	Code string `json:"city_code,omitempty"`
	Name string `json:"city_name,omitempty"`
}

type location struct {
	CityName     string `json:"city_name,omitempty"`
	ProvinceName string `json:"province_name,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
}

type session struct {
	sessionMap map[string]credential
	mutex      sync.RWMutex
}
type credential struct {
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword []byte `json:"-"`
}

type credentials struct {
	credentialMap map[string]credential
	mutex         sync.RWMutex
}

type message struct {
	Items   interface{} `json:"items,omitempty"`
	Message string      `json:"message,omitempty"`
}

type errorMessage struct {
	ErrorMessage string `json:"error,omitempty"`
}

type user struct {
	Name     string     `json:"name,omitempty"`
	Role     string     `json:"role,omitempty"`
	Parent   string     `json:"parent,omitempty"`
	Includes []location `json:"includes_region,omitempty"`
	Excludes []location `json:"excludes_region,omitempty"`
	Children []*user    `json:"children,omitempty"`
}
