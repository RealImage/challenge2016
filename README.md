# DISTRIBUTORS and PERMISSIONS
 SOlution for the task https://github.com/RealImage/challenge2016

## PREREQUISITES

1. Golang must be installed
2. Run

```
go mod download
```

## SOLUTION

Created a simple API to interact with the program using "Gin" framework.

For running the program, use,
```
$ go run .
```
command in CLI.

It will start the program and Port "8090" will be open to receive API calls. 

## Available Endpoints

### 1. POST /create-distributor
    To create a new distributor.
    JSON body to be passed,

        "distributor" : "name of the distributor", // required
        "parentdistributor" : "name of the parent distributor(if any)",
        "include" : [permissions to be added seperated by comma]  // required
        "exclude" : [permissions to be excluded seperated by comma]

        include/exclude format - "CityCode-ProvinceCode-CountryCode" or "ProvinceCode-CountryCode" or "CountryCode"
            Example : CENAI-TN-IN

        Example body:
            {
                "distributor": "distributor1",
                "parentdistributor": "",
                "include": [
                    "IN",
                    "FL-US"
                ],
                "exclude": [
                    "TN-IN"
                ]
            }

### 2. GET /distributors

    For listing the available distributors and ther Include, Exclude, Parent Distributor details(if any).

### 3. GET /permission-check/:distributor/:permission
    To check a particular distributor have access to a particular region.
    URL params,
        :distributor - name of the distributor.
        :permission - name of the permission to be checked.

### For reference a postman collection of APIs added in the repository.