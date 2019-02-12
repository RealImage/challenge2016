# Real Image Challenge 2016 Solution - Thamizh

`cities.csv` contains two blank fields`(B16737 and B77227)` in Province code column.
Created `cities1.csv` which is same as cities.csv except I gave some code to the above blank fields
I have used country, province and city names instead of codes as I saw some of the codes are duplicated.

To run the program 
MacOS/Linux
```
go get ./...
go run *.go

```

Windows
```
go get ./...
go run ./

```

Solution is designed as API server with four endpoints.
Two roles are used
```
1. admin - Can view all users, create admin and distributor users but dont have distribution rights.
2. distributor - Can view self, users created by them and have distribution rights only to the "includes" region.
                 Can create only distributors and assign them include or exclude regions they are entitled to.

```
The data structures used can be viewed from types.go file.
Error message are saved in constants.go file.
Postman colletion is attahced for both roles with sample requests.
cities1.csv info is turned into n-ary tree and that json format is stored in output.json for reference.




