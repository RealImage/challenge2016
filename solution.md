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

When go run command is used server will be started. To stop the server presee `Cmd + C` or `Ctrl + C`
Solution is designed as API server running at 8080 port with four endpoints.

Two roles are used
```
1. admin - Can view all users, create admin and distributor users but dont have distribution rights.
2. distributor - Can view self, users created by them and have distribution rights only to the "includes" region.
                 Can create only distributors and assign them include or exclude regions they are entitled to.

```

A default admin user is created with below credentials.
```
Username: admin
Password: admin

```
Whenever a distributor is created with any name the default password is `distributor`

Location is a combination of country name, province name and city name.
```
Country name cannot be empty
City name must be empty when province name is empty.
```

Action Flow in application
```
1. Login to Get Authentication token by providing username and password
2. See users under the logged in user. (Admin can see all users.)
3. Create new user (If "parent" field is mentioned and parent user comes under logged in user, new user will be created under parent
                    user. If "parent field is blank, new user will be created under logged in user.)
4. Validate distribution rights for a location (If "username" field is blank, distribution rights will be checked for logged in user. If
                                                "username" is mentioned and if that user comes under logged in user, distribution rights 
                                                will be checked for that user.)
5. Logout by invalidating token

Note: While using postman, logout a particular user and then login with another user.
```

The data structures used can be viewed from types.go file.
Error message are saved in constants.go file.
Postman colletion is attahced for both roles with sample requests.
cities1.csv info is turned into n-ary tree and that json format is stored in output.json for reference.




