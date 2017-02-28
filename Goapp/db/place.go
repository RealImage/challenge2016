package db

import (    
    "encoding/csv"
    "os"
    "fmt"
    "io"
    "strings"
    "github.com/bradfitz/slice"
)

type Place struct {
    Id int
    CityCode string
    ProvinceCode  string
    CountryCode  string 
    City string
    Province string
    Country string
    FormattedName string //Used only for search, not serialized
    FormattedCode string //Used only for search, not serialized
    MatchIndex int //Used only for search, not serialized
}

//Acts as DB
var PlaceDB []Place

//Serializers to convert into jsonapi.org conventions
func SerializePlace(p Place) map[string]interface{} {
    return map[string]interface{}{
            "id" : p.Id,
            "type" : "places",
            "attributes": map[string]interface{}{
                            "cityCode" : p.CityCode,
                            "provinceCode" : p.ProvinceCode,
                            "countryCode" : p.CountryCode,
                            "city" : p.City,
                            "province" : p.Province,
                            "country" : p.Country,
                            "formattedName": p.FormattedName,
                            "formattedCode": p.FormattedCode,
                        },
        }            
}

//Serialize array of places
func SerialzeAllPlace(places []Place) map[string]interface{} {    
    var serializedPlaces []map[string]interface{}

    for _,p := range places {
        serializedPlaces = append(serializedPlaces, SerializePlace(p))
    }

    return map[string]interface{}{"data": serializedPlaces}
}

//Check if element exist in array
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

//Find type of place city, province or country
func placeType(placeCode string) string{
    var placeTypeParts = strings.Split(placeCode, "-")

    if(len(placeTypeParts) == 3){
        return "city"
    }else if(len(placeTypeParts) == 2){
        return "province"
    }else{
        return "country"
    }
}

//Returns serialized search results
func SearchPlace(query string, query_type string) map[string]interface{} {
    var matchingPlaces []Place
    var placeCodes []string

    for _,place := range PlaceDB {
        var search_term string
        var formattedCode string

        if(strings.ToLower(query_type) == "country"){
            search_term = strings.Join(strings.Split(place.FormattedName, ", ")[2:3], ", ")
            formattedCode = strings.Join(strings.Split(place.FormattedCode, "-")[2:3], "-")
        }else if(strings.ToLower(query_type) == "province"){
            search_term = strings.Join(strings.Split(place.FormattedName, ", ")[1:3], ", ")
            formattedCode = strings.Join(strings.Split(place.FormattedCode, "-")[1:3], "-")
        }else{
            search_term = place.FormattedName
            formattedCode = place.FormattedCode
        }

        var matchingIndex int = strings.Index(strings.ToLower(search_term), strings.ToLower(query))

        if(matchingIndex != -1 && !contains(placeCodes, formattedCode)){
            placeCodes = append(placeCodes, formattedCode)
            matchingPlaces = append(matchingPlaces, Place{place.Id, place.CityCode, place.ProvinceCode, place.CountryCode, place.City, place.Province, place.Country, search_term, formattedCode, matchingIndex})
        }
    }

    slice.Sort(matchingPlaces[:], func(i, j int) bool {
        return matchingPlaces[i].MatchIndex < matchingPlaces[j].MatchIndex
    })

    if(len(matchingPlaces) >= 5){
        matchingPlaces = matchingPlaces[0:5]
    }else{
        matchingPlaces = matchingPlaces[0:len(matchingPlaces)]
    }

    return SerialzeAllPlace(matchingPlaces)
}

func getProvince(place string) string{
    var placeArray []string = strings.Split(place, "-")
    if(len(placeArray) == 3){
        return strings.Join(placeArray[1:], "-")
    }else{
        return place
    }
}

func getCountry(place string) string{
    var placeArray []string = strings.Split(place, "-")
    if(len(placeArray) == 3){
        return strings.Join(placeArray[2:], "-")
    }else if (len(placeArray) == 2){
        return strings.Join(placeArray[1:], "-")
    }else{
        return place
    }
}

//Read the CSV file(Errors are not handles well since I know this is a valid CSV)
func SeedPlaceDB(){
    //Read as CSV
    var file, error = os.Open("cities.csv")
    if error != nil {
        fmt.Println("Error")
    }
    var Reader = csv.NewReader(file)

    lineCount := 0
    for {
        record, err := Reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
        }
        
        if(lineCount > 0){
            PlaceDB = append(PlaceDB, Place{lineCount, record[0], record[1], record[2], record[3], record[4], record[5], (record[3] + ", " + record[4] + ", " + record[5]), (record[0] + "-" + record[1] + "-" + record[2]), 0})
        }

        lineCount += 1
    }
}