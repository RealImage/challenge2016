# Distribution application

## Introduction
This application helps in creating a hierarchy of distributors, including and excluding the distribution areas and also helps in querying whether a distributor is allowed to distribute in an area or not.

## Installation

1. Install Go latest version
2. Install CSV parser
```
go get -u github.com/gocarina/gocsv
```

## Run application

```
go run main.go -file cities.csv
```

## TODO

1. Cover test cases to validate distributor scope
2. Add usage documentation

## Authors

1. Ilayaraja M (ilayaraja.edu@gmail.com)