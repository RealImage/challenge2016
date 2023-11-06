# Distributor Permissions API

The Distributor Permissions API is a Go web application with Gin that allows you to manage and check distributor permissions for specific geographic locations. This API enables you to add distributors, retrieve a list of all distributors, and check if a distributor has permission in a given area.

## Getting Started

Follow these instructions to set up and run the Distributor Permissions API on your local machine.

### Prerequisites

Ensure you have the following tools installed:

- [Go (Golang)](https://golang.org/dl/)

### Installation

 Clone the repository:

   git clone https://github.com/kamlesh1807/challenge2016.git
   
   cd challenge2016
   
Build and run the application:

go run main.go
The API should be running on http://localhost:8080.

API Endpoints
Get All Distributors: Retrieve a JSON list of all distributors.
URL: GET /get-all-distributors

Check Permissions: Check if a distributor has permission in a given area.
URL: GET /check-permissions

Query Parameters:
distributor_name: The name of the distributor.
location: The geographic location to check permissions for.
Example:
GET /check-permissions?distributor_name=Aman001&location=IN


Add Distributor: Add a new distributor to the system.
URL: POST /add-distributor
Request Body: JSON object with distributor details.
Example:
json
{
    Name:      "Aman001",
		Cities:    []string{"PUNCH-JK-IN", "KLRAI-TN-IN" , "PLMYR-NY-US"},
		States:    []string{"TN-IN","AZ-US"},
		Countries: []string{"IN","US"},
		ExCities: []string{},
		ExStates: []string{},
		ExCountries: []string{},
		Addedby: "Admin",

}
    


