FILM DISTRIBUTION MANAGEMENT SYSTEM
========================================================
- [FILM DISTRIBUTION MANAGEMENT SYSTEM](#film-distribution-management-system)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [How to use it?](#how-to-use-it)
  - [Limitations](#limitations)
## Introduction
 This is an application to manage film distribution. It's a CLI based application written in Go. 
## Features
    1. Add a film
    2. Remove a film
    3. Add Distributor
    4. Remove Distributor
    5. Add region data from csv file (remote/local)
    6. Authorize distributor to distribute film in a region
    7. Check if a distributor is authorized to distribute a film in a region
## Prerequisites
    1. Go
    2. Git
## How to use it?
    1. Clone the repository
    2. Run the command `go run cmd/main.go`
## Limitations
    1. This application is case sensitive. It treats ABC and abc as different inputs.
    2. This is only tested to take CSV file as input data for regions. The format of data supported is this order(City Code,Province Code,Country Code,City Name,Province Name,Country Name)
   
