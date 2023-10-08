# DESCRIPTION

This is a solution for the challenge provided in https://github.com/RealImage/challenge2016 (Refere QUESTION_README.md for question)

Developed a simple API interface using golang to interact with the inner system. The api is based on the "github.com/gin-gonic/gin" (gin in short) api framework.

For reference purpose, an example postman collection is given in the repo, please check it in the postman (make changes if required)

## PREREQUISITE

1. Must have go installed on the system

2. Run the below command in this directory before running the program to download the required libraries
```
go mod download
```

## HOW TO RUN THE PROGRAM

```
$ go run . 
```
This will start the API service and can be accesible through 8000 port (You can change the port in main.go if you have conlfict running in port 8000)

## ENDPOINTS FOR REFERENCE

### GET : /list-cities

Endpoint to list the available cities, province and country data loaded from the given csv file

### POST : /add-distributors

Endpoint to add a distributor to the system. Example json body format,

##### Note: You can omit the `parent` key if the distributor don't have a parent

```
{
	"id" : "distributor2",
	"parent" : "distributor1",
	"included" : ["IN"],
	"excluded" : ["TN-IN"]
}
```
id - Ditributor unqiue id/username (mandatory)

parent - Id of the parent distributor (optional)

included - List of regions included (optional)

excluded - List of inner regions which need to be excluded (optional)


### GET : /list-distributors

Endpoint to list the available distributors in the system

### GET : /check-distributor-region/`<distributor>`/`<region>`

Endpoint to check the permission of a distributor in a corresponding region

Example: If distributor1 included in India(IN), but excluded in Tamilnadu in India(TN-IN)

GET : /check-region/distributor1/KA-IN will return 200 as response code and `{"data": "YES"}` as response

GET : /check-region/distributor1/CENAI-TN-IN will return 403 as response code and `{"data": "NO"}` as response

