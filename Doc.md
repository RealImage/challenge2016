# Assumptions

* One city can have multiple distributors.
* One child distributor can't have multiple parent distributors.
* After adding child, parent can include more locations but can't exclude locations.

# Building Code

1. Building Server Code.
    * set GOPATH
    * go to $GOPATH/src/github.com/RealImage/challenge2016/location/cmd/locationService/server
    * open command prompt and run command go build
    * binary run --help for help

2. Building Location Client To add location
    * go to $GOPATH/src/github.com/RealImage/challenge2016/location/cmd/locationService/client
    * open command prompt and run command go build
    * run --help and set location server address:port.

# Running Server
* run generated binary with config
* example server.exe -l=false -p 8080 

# Adding Locations

* run location client binary with specifying csv file path and location server host:post.
* example client.exe -l=false -i "cities.csv" -host "localhost:8080"

# Adding Distributors

* send post request to url http://host:port/api/distributor/v1 with  json encoded body

```
{
    "parent_id":"", //specify if this distributor is child distributor else leave empty 
	"id":"", // distributor unique id
	"permission":"", // either Access Or Granted
	"country_code":"", // country code
	"state_code":"", // state code
	"city_code":""  //city code
}
```
* example: to do below you have to send 4 request 
```
Permissions for DISTRIBUTOR1
INCLUDE: INDIA
INCLUDE: UNITEDSTATES
EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA
```

```
1.  {
    	"parent_id":"",
    	"id":"DISTRIBUTOR1",
    	"permission":"Granted",
    	"country_code":"In"
    }

2.  {
    	"parent_id":"",
    	"id":"DISTRIBUTOR1",
    	"permission":"Granted",
    	"country_code":"US"
    }
3.  {
    	"parent_id":"",
    	"id":"DISTRIBUTOR1",
    	"permission":"Denied",
    	"country_code":"In",
    	"state_code":"KARN"
    }
4.  {
    	"parent_id":"",
    	"id":"DISTRIBUTOR1",
    	"permission":"Denied",
    	"country_code":"In",
    	"state_code":"TAML",
    	"city_code":"CHEN"
    }
```

* adding child(sub) distributor permission
 ```
Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: INDIA
EXCLUDE: TAMILNADU-INDIA
```
```
 1.  {
    	"parent_id":"DISTRIBUTOR2",
    	"id":"DISTRIBUTOR1",
    	"permission":"Granted",
    	"country_code":"In"
    }
 2.  {
    	"parent_id":"DISTRIBUTOR2",
    	"id":"DISTRIBUTOR1",
    	"permission":"Denied",
    	"country_code":"In",
    	"state_code":"TAML"
    }
```

*  success response code is 200 else error description is found in response body with json encoded error example : {"error":"Invalid Location"}

# Asking Location Permission
CHICAGO-ILLINOIS-UNITEDSTATES for distributor1
* send get request to url host:port/api/distributor/v1/permission with query string
```
    id=DISTRIBUTOR1
    country_code=US
    state_code=ILLI
    city_code=CHI
```    
* example : localhost:8080/api/distributor/v1/permission?id=DISTRIBUTOR1&country_code=US&state_code=ILLI&city_code=CHI
`