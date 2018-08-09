# Solution for **Real Image Challenge 2016**

## Installation

If GoLang already installed skip the below step

[Refer this link](https://golang.org/doc/install) to install GoLang in your machine.

## Running

    Place the csv file in project home directory (or) Specify it using -file flag

    Start the service

    1. If (cities data)CSV file is in project home directory
        go run main.go
    
    2. If you want to add a custom location for (cities data)CSV file
        go run main.go -file csv/cities.csv

## APIs Available:
`/createDistributor`
    - Input:
        - Array of Distributors with parent distributor details, included & excluded locations
    - Output:
        - Created Distributors details as response
    - Check
        - Checking for duplicates
        - Checking for valid parent distributor

`/verifyDistribution`
    - Input:
        - Array of Distributors with locations
    - Output
        - "YES" or "NO" response with city/state/country

## Sample APIs with data

**Request 1**

    curl -XPOST 'http://localhost:4000/createDistributor' -d '[{"parentDistributor":"none","distributorName":"Distributor1","includedLocations":[{"countryCode":"IN"},{"countryCode":"US"}],"excludedLocations":[{"countryCode":"IN","stateCode":"KA"},{"countryCode":"IN","stateCode":"TN","cityCode":"CENAI"}]},{"parentDistributor":"Distributor1","distributorName":"Distributor2","includedLocations":[{"countryCode":"IN"},{"countryCode":"US","stateCode":"CA"}],"excludedLocations":[{"countryCode":"IN","stateCode":"TN"},{"countryCode":"IN","stateCode":"GJ"}]},{"parentDistributor":"Distributor2","distributorName":"Distributor3","includedLocations":[{"countryCode":"IN","stateCode":"TN","cityCode":"CENAI"},{"countryCode":"IN","stateCode":"KA","cityCode":"HBALI"},{"countryCode":"IN","stateCode":"KL"},{"countryCode":"IN","stateCode":"GJ","cityCode":"NAVSR"}]}]'

**Response:**

    {"Distributor1":{"parentDistributor":"none","distributorName":"Distributor1","includedLocations":[{"cityCode":"","stateCode":"","countryCode":"IN"},{"cityCode":"","stateCode":"","countryCode":"US"}],"excludedLocations":[{"cityCode":"","stateCode":"KA","countryCode":"IN"},{"cityCode":"CENAI","stateCode":"TN","countryCode":"IN"}]},"Distributor2":{"parentDistributor":"Distributor1","distributorName":"Distributor2","includedLocations":[{"cityCode":"","stateCode":"","countryCode":"IN"},{"cityCode":"","stateCode":"CA","countryCode":"US"}],"excludedLocations":[{"cityCode":"","stateCode":"TN","countryCode":"IN"},{"cityCode":"","stateCode":"GJ","countryCode":"IN"}]},"Distributor3":{"parentDistributor":"Distributor2","distributorName":"Distributor3","includedLocations":[{"cityCode":"CENAI","stateCode":"TN","countryCode":"IN"},{"cityCode":"HBALI","stateCode":"KA","countryCode":"IN"},{"cityCode":"","stateCode":"KL","countryCode":"IN"},{"cityCode":"NAVSR","stateCode":"GJ","countryCode":"IN"}],"excludedLocations":null}}

**Request 2**

    curl -XPOST 'http://localhost:4000/verifyDistribution' -d '[{"distributorName":"Distributor1","location":[{"countryCode":"US","stateCode":"CA","cityCode":"FILMR"},{"countryCode":"US","stateCode":"CA","cityCode":"AVEJO"},{"countryCode":"IN","stateCode":"TN","cityCode":"CENAI"},{"countryCode":"IN","stateCode":"TN","cityCode":"NMAKL"},{"countryCode":"IN","stateCode":"KA","cityCode":"MYSUR"}]},{"distributorName":"Distributor2","location":[{"countryCode":"US","stateCode":"CA"},{"countryCode":"IN","stateCode":"KA","cityCode":"NGAML"},{"countryCode":"IN","stateCode":"TN","cityCode":"NGAML"},{"countryCode":"IN","stateCode":"TN","cityCode":"SALEM"},{"countryCode":"IN","stateCode":"KA","cityCode":"MYSUR"}]},{"distributorName":"Distributor3","location":[{"countryCode":"US","stateCode":"CA"},{"countryCode":"IN","stateCode":"KA","cityCode":"NGAML"},{"countryCode":"IN","stateCode":"KL"},{"countryCode":"IN","stateCode":"KL", "cityCode":"THRAP"},{"countryCode":"IN","stateCode":"TN","cityCode":"NGAML"},{"countryCode":"IN","stateCode":"TN","cityCode":"SALEM"},{"countryCode":"IN","stateCode":"KA","cityCode":"MYSUR"},{"countryCode":"IN","stateCode":"GJ","cityCode":"NAVSR"}]}]'

**Response:**

    {
        "Distributor1": {
            "AVEJO-CA-US": "YES",
            "CENAI-TN-IN": "NO",
            "FILMR-CA-US": "YES",
            "MYSUR-KA-IN": "NO",
            "NMAKL-TN-IN": "YES"
        },
        "Distributor2": {
            "CA-US": "YES",
            "MYSUR-KA-IN": "NO",
            "NGAML-KA-IN": "NO",
            "NGAML-TN-IN": "NO",
            "SALEM-TN-IN": "NO"
        },
        "Distributor3": {
            "CA-US": "NO",
            "KL-IN": "YES",
            "THRAP-KL-IN":"YES",
            "MYSUR-KA-IN": "NO",
            "NAVSR-GJ-IN": "NO",
            "NGAML-KA-IN": "NO",
            "NGAML-TN-IN": "NO",
            "SALEM-TN-IN": "NO"
        }
    }

## Unit test

For running unit test, you to need to get smartystreets gunit & assertions packages

    go get github.com/smartystreets/gunit

    go get github.com/smartystreets/assertions/should

Then change your current directory to distributor/

    go test