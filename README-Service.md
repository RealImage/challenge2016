# Distribution Service(Real Image Challenge)

## Introduction:
This Distribution REST Service helps you to create Distributors and Sub-Distributors Permissions and also Check the Permissions of any Distributor. You can also List the Distributors available with us.

## Installation:

1. Install Go latest version

## To Run the Service:

You can Run the Service using "go run" command. flags to choose the port on which the Service should Run and file flag to specify the CSV File Path are provided.
```
go run main.go -port 7770 -file cities.csv
```

The Above Command Starts the REST Service on port 7770 and reads the cities.csv file from your present directory.

If Go is Not Installed on your machine, you can still Run the Service using the below command. (Executable Binary File is Also Provided)
```
./main -port 7770 -file cities.csv
```

## EndPoints Available:

/addDist/ - This Takes the Array of JSON filled with Distributors Permissions and stores the permissions in its respective Distributor Template.(Sample JSON provided under the Sample JSON Structure)

/verifyDist/ - This lets you verify a particular Distributor Permissions by the taking the location you want to check through JSON.(Sample JSON provided under the Sample JSON Structure)

/listDist/ - On a GET Request to this Endpoint it will Return you an array with the list of Distributors stored.

(For Example : If the Service is Running on localhost and 7770 port then "localhost:7770/addDist/" || "localhost:7770/verifyDist/")

## Logging & Monitoring:

Server Logs INFO & ERROR Logs after every action. you can redirect the logs to the required file for Debugging Purposes.

## Sample JSON:
	---Sample JSON also provided in addDistsample.json & verifyDistsample.json---
	
For /addDist:

[
  {
    "parentDistributorName": "none",
    "distributorName": "Distributor1",
    "includeData": [
      {
        "countryCode": "IN",
        "provinceCode": "TN",
        "cityCode": "ERODE"
      },
      {
        "countryCode": "IN",
        "provinceCode": "TN",
        "cityCode": "COBTE"
      },
      {
        "countryCode": "IN",
        "provinceCode": "TN",
        "cityCode": "MDURI"
      },
      {
        "countryCode": "ES",
        "provinceCode": "*",
        "cityCode": "*"
      },
      {
        "countryCode": "JO",
        "provinceCode": "*",
        "cityCode": "*"
      },
      {
        "countryCode": "US",
        "provinceCode": "CA",
        "cityCode": "*"
      }
    ],
    "excludeData": [
      {
        "countryCode": "IN",
        "provinceCode": "TN",
        "cityCode": "NMAKL"
      },
      {
        "countryCode": "US",
        "provinceCode": "CA",
        "cityCode": "ALMED"
      },
      {
        "countryCode": "JO",
        "provinceCode": "MN",
        "cityCode": "*"
      }
    ]
  },
  {
    "parentDistributorName": "Distributor1",
    "distributorName": "Distributor2",
    "includeData": [
      {
        "countryCode": "ES",
        "provinceCode": "EX",
        "cityCode": "*"
      },
      {
        "countryCode": "JO",
        "provinceCode": "AM",
        "cityCode": "*"
      },
      {
        "countryCode": "US",
        "provinceCode": "CA",
        "cityCode": "AVEJO"
      }
    ],
    "excludeData": [
      {
        "countryCode": "ES",
        "provinceCode": "EX",
        "cityCode": "VLENA"
      },
      {
        "countryCode": "JO",
        "provinceCode": "AM",
        "cityCode": "WASIR"
      }
    ]
  }
]

For /verifyDist:

{
  "distributorName": "Distributor2",
  "locations": [
    {
      "countryCode": "JO",
      "provinceCode": "AM",
      "cityCode": "WASIR"
    },
    {
      "countryCode": "US",
      "provinceCode": "CA",
      "cityCode": "AVEJO"
    }
  ]
}

## Test Results:

Request-1: curl -XPOST "http://localhost:7770/addDist/" -d '[{"parentDistributorName":"none","distributorName":"Distributor1","includeData":[{"countryCode":"IN","provinceCode":"TN","cityCode":"ERODE"},{"countryCode":"IN","provinceCode":"TN","cityCode":"COBTE"},{"countryCode":"IN","provinceCode":"TN","cityCode":"MDURI"},{"countryCode":"ES","provinceCode":"*","cityCode":"*"},{"countryCode":"JO","provinceCode":"*","cityCode":"*"},{"countryCode":"US","provinceCode":"CA","cityCode":"*"}],"excludeData":[{"countryCode":"IN","provinceCode":"TN","cityCode":"NMAKL"},{"countryCode":"US","provinceCode":"CA","cityCode":"ALMED"},{"countryCode":"JO","provinceCode":"MN","cityCode":"*"}]},{"parentDistributorName":"Distributor1","distributorName":"Distributor2","includeData":[{"countryCode":"ES","provinceCode":"EX","cityCode":"*"},{"countryCode":"JO","provinceCode":"AM","cityCode":"*"},{"countryCode":"US","provinceCode":"CA","cityCode":"AVEJO"}],"excludeData":[{"countryCode":"ES","provinceCode":"EX","cityCode":"VLENA"},{"countryCode":"JO","provinceCode":"AM","cityCode":"WASIR"}]}]'

Response-1: {"ServerResponse":["Created Distributor: Distributor1 And Sucessfully Updated Permissions; ","Created Distributor: Distributor2 And Sucessfully Updated Permissions; "]}

Request-2: curl -XPOST "http://localhost:7770/verifyDist/" -d '{"distributorName":"Distributor2","locations":[{"countryCode":"JO","provinceCode":"AM","cityCode":"WASIR"},{"countryCode":"US","provinceCode":"CA","cityCode":"AVEJO"}]}'

Response-2: {"ServerResponse":["NO For : WASIR-AM-JO; ","YES For : AVEJO-CA-US; "]}

Request-3: curl "http://localhost:7770/listDist/"

Response-3: {"ServerResponse":["Distributor1","Distributor2"]}


## Author:

   NAGA SAI AAKARSHIT BATCHU (aakarshit.batchu@gmail.com)
