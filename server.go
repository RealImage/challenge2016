/* server.go

 Usage(from working directory in command prompt):
  	$ go get "github.com/julienschmidt/httprouter"
   	$ go run server.go
 API(open Browser):
		http://localhost:3000/DISTRIBUTOR1
		http://localhost:3000/DISTRIBUTOR2
		http://localhost:3000/DISTRIBUTOR3
*/
package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Data struct {
	CityName     string `field:"City Name"`
	ProvinceName string `field:"Province Name"`
	CountryName  string `field:"Country Name"`
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Handler for default page
	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Simply write some test data for now
		fmt.Fprint(w, "Welcome to Homepage!\n")
	})

	// Added a handlers for each Distributors
	r.GET("/DISTRIBUTOR1", Distributor1Handler)
	r.GET("/DISTRIBUTOR2", Distributor2Handler)
	r.GET("/DISTRIBUTOR3", Distributor3Handler)

	// Fire up the server
	fmt.Println("Listening to port:3000")
	http.ListenAndServe("localhost:3000", r)
}
func Distributor1Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	// read the csv file
	csvFile, err := os.Open("cities.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Data
	var allRecords []Data

	// create a loop array for distributor1
	/*	Permissions for DISTRIBUTOR1
		INCLUDE: INDIA
		INCLUDE: UNITEDSTATES
		EXCLUDE: KARNATAKA-INDIA
		EXCLUDE: CHENNAI-TAMILNADU-INDIA
	*/
	for _, each := range csvData {
		if each[3] == "City Name" && each[4] == "Province Name" && each[5] == "Country Name" {
			fmt.Println("Explicit headers are neglected")
		} else if each[5] == "India" || each[5] == "United States" {
			if each[4] != "Karnataka" {
				if each[3] != "Chennai" {
					oneRecord.CityName = each[3]
					oneRecord.ProvinceName = each[4]
					oneRecord.CountryName = each[5]
					allRecords = append(allRecords, oneRecord)
				}
			}
		}
	}

	html1 := "<html>" +
		"<body>" +
		"<h1>Permitted regions for Distributors1</h1>" +
		"<h3>Note: Permission denied countries and its repective provinces and cities are excluded and also excluded Chennai in Tamil Nadu and Karnataka provinces in India</h2>" +
		"<table>" +
		"<tr>" +
		"<th>City Name</th>" +
		"<th>Province Name</th>" +
		"<th>Country Name</th>" +
		"</tr>"

	html2 := ""
	for i := 0; i < len(allRecords); i++ {

		html2 += "<tr>" +
			"<td>" + allRecords[i].CityName + "</td>" +
			"<td>" + allRecords[i].ProvinceName + "</td>" +
			"<td>" + allRecords[i].CountryName + "</td>" +
			"</tr>"
	}

	html3 := "</table>" +
		"</body>" +
		"</html>"
	webPage := html1 + html2 + html3

	fmt.Fprint(w, string(webPage))
}
func Distributor2Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	// read the csv file
	csvFile, err := os.Open("cities.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Data
	var allRecords []Data

	// create a loop array for distributor2
	/*	Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
		INCLUDE: INDIA
		INCLUDE: UNITEDSTATES
		EXCLUDE: KARNATAKA-INDIA
		EXCLUDE: TAMILNADU-INDIA
	*/
	for _, each := range csvData {
		if each[3] == "City Name" && each[4] == "Province Name" && each[5] == "Country Name" {
			fmt.Println("Explicit headers are neglected")
		} else if each[5] == "India" || each[5] == "United States" {
			if each[4] != "Karnataka" && each[4] != "Tamil Nadu" {
				oneRecord.CityName = each[3]
				oneRecord.ProvinceName = each[4]
				oneRecord.CountryName = each[5]
				allRecords = append(allRecords, oneRecord)
			}
		}
	}

	html1 := "<html>" +
		"<body>" +
		"<h1>Permitted regions for Distributors2</h1>" +
		"<h3>Note: Permission denied countries and its repective provinces and cities are excluded and also excluded Tamil Nadu and Karnataka provinces in India</h2>" +
		"<table>" +
		"<tr>" +
		"<th>City Name</th>" +
		"<th>Province Name</th>" +
		"<th>Country Name</th>" +
		"</tr>"

	html2 := ""
	for i := 0; i < len(allRecords); i++ {

		html2 += "<tr>" +
			"<td>" + allRecords[i].CityName + "</td>" +
			"<td>" + allRecords[i].ProvinceName + "</td>" +
			"<td>" + allRecords[i].CountryName + "</td>" +
			"</tr>"
	}

	html3 := "</table>" +
		"</body>" +
		"</html>"
	webPage := html1 + html2 + html3

	fmt.Fprint(w, string(webPage))
}
func Distributor3Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	// read the csv file
	csvFile, err := os.Open("cities.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Data
	var includeRecord Data
	var allRecords []Data

	// create a loop array for distributor3
	/*	Permissions for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
		INCLUDE: INDIA
		INCLUDE: UNITEDSTATES
		INCLUDE: HUBLI-KARNATAKA-INDIA
		EXCLUDE: TAMILNADU-INDIA
	*/
	for _, each := range csvData {
		if each[3] == "City Name" && each[4] == "Province Name" && each[5] == "Country Name" {
			fmt.Println("Explicit headers are neglected")
		} else if each[5] == "India" || each[5] == "United States" {

			if each[3] == "Hubballi" && each[4] == "Karnataka" {
				includeRecord.CityName = each[3]
				includeRecord.ProvinceName = each[4]
				includeRecord.CountryName = each[5]
				allRecords = append(allRecords, includeRecord)
			}
			if each[4] != "Karnataka" && each[4] != "Tamil Nadu" {
				oneRecord.CityName = each[3]
				oneRecord.ProvinceName = each[4]
				oneRecord.CountryName = each[5]
				allRecords = append(allRecords, oneRecord)
			}

		}
	}

	html1 := "<html>" +
		"<body>" +
		"<h1>Permitted regions for Distributors3</h1>" +
		"<h3>Note: Permission denied countries and its repective provinces and cities are excluded and also excluded Tamil Nadu and Karnataka(permission for Hubballi granted) provinces  in India</h2>" +
		"<table>" +
		"<tr>" +
		"<th>City Name</th>" +
		"<th>Province Name</th>" +
		"<th>Country Name</th>" +
		"</tr>"

	html2 := ""
	for i := 0; i < len(allRecords); i++ {

		html2 += "<tr>" +
			"<td>" + allRecords[i].CityName + "</td>" +
			"<td>" + allRecords[i].ProvinceName + "</td>" +
			"<td>" + allRecords[i].CountryName + "</td>" +
			"</tr>"
	}

	html3 := "</table>" +
		"</body>" +
		"</html>"
	webPage := html1 + html2 + html3

	fmt.Fprint(w, string(webPage))
}
