package db

import (    
    "fmt"
)

type Distributor struct {
    Id int 
    Name string 
    ParentDistributorId int 
    Includes  []string 
    FormattedIncludes []string
    Excludes  []string     
    FormattedExcludes []string
}

var DistributorDB map[string][]Distributor
var DistributorCount int

func InitDB(){    
    DistributorDB = map[string][]Distributor{"SampleDb": SampleDB()}
}

func NextId() int{
    DistributorCount = DistributorCount + 1
    return DistributorCount
}

func FindDistributor(distributors []Distributor, id int) Distributor{
    var distributor Distributor
    for _,d := range distributors {
        if(d.Id == id){
            return d
        }
    }

    return distributor
}

func isAllowed(mainRegion string, subRegion string) bool{
    if(mainRegion == subRegion){
        return true
    }
    fmt.Println(mainRegion)
    fmt.Println(subRegion)
    if(placeType(mainRegion) == "country"){
        //return true if state available in mainRegion
        if(getCountry(subRegion) == mainRegion){
            fmt.Println("true")
            return true
        }
    }

    if(placeType(mainRegion) == "province"){
        //return true if country available in mainRegion
        if(getProvince(subRegion) == mainRegion){
            fmt.Println("true")
            return true
        }
    }

    fmt.Println("false")
    return false
}

func AddToDb(sessionId string, placeCode string, id int, formattedName string, listType string) {
    var db []Distributor
    for _,d := range DistributorDB[sessionId] {
        if(d.Id == id){
            if(listType == "include"){
                d.Includes = append(d.Includes, placeCode)
                d.FormattedIncludes = append([]string{formattedName}, d.FormattedIncludes...)
            }else{
                d.Excludes = append(d.Excludes, placeCode)
                d.FormattedExcludes = append([]string{formattedName}, d.FormattedExcludes...)
            }         
        }
        db = append(db, d)
    }
    
    DistributorDB[sessionId] = db    
}

//listType may be include or exclude
func AddPlace(sessionId string, distributorId int, placeCode string, listType string, formattedName string) bool{ 
    var distributor, parentDistributor Distributor
    //find distributor
    distributor = FindDistributor(DistributorDB[sessionId], distributorId)
    var addFlag bool = false    

    //Insert without conditions if it is main distributor
    for _,excludeCode := range distributor.Excludes {
        if(isAllowed(excludeCode, placeCode)){
            fmt.Println("1 false")
            return false
        }
    }

    if(listType == "exclude"){
        if(contains(distributor.Excludes, placeCode)){
            return false
        }
        for _,includeCode := range distributor.Includes {
            if(isAllowed(placeCode, includeCode)){
                fmt.Println("2 false")
                return false
            }
        }
    }else{
        if(contains(distributor.Includes, placeCode)){
            return false
        }
    }
    if(distributor.ParentDistributorId != 0){
        var parentDistributorId int = distributor.ParentDistributorId
        var mainParentDistributor Distributor = FindDistributor(DistributorDB[sessionId], distributor.ParentDistributorId)

        for(parentDistributorId != 0){
            parentDistributor = FindDistributor(DistributorDB[sessionId], parentDistributorId)
            for _,excludeCode := range parentDistributor.Excludes {
                if(isAllowed(excludeCode, placeCode)){
                    fmt.Println("3 false")
                    return false
                }
            }
            parentDistributorId = parentDistributor.ParentDistributorId
        }

        for _,includeCode := range mainParentDistributor.Includes {
            if(isAllowed(includeCode, placeCode)){       
                fmt.Println("4 true")         
                addFlag = true
            }
        }
    }else{
        fmt.Println("5 true")         
        addFlag = true
    }
    

    if(addFlag){
        AddToDb(sessionId, placeCode, distributorId, formattedName, listType)
        return true
    }

    return false

}

func SampleDB() []Distributor{
    var userDB []Distributor
    userDB = append(userDB, Distributor{NextId(), "Distributor 1", 0, []string{"IN", "US"}, []string{"India", "United States"}, []string{"CENAI-TN-IN", "KA-IN"}, []string{"Chennai, Tamil Nadu, India", "Karnataka, India"}})    
    return userDB
}

//Serializers to convert into jsonapi.org conventions
func SerializeDistributor(d Distributor) map[string]interface{} {
    return map[string]interface{}{
            "id" : d.Id,
            "type" : "distributors",
            "attributes": map[string]interface{}{
                            "name" : d.Name,
                            "parentDistributorId" : d.ParentDistributorId,
                            "includes" : d.Includes,
                            "excludes" : d.Excludes,
                            "formattedIncludes": d.FormattedIncludes,
                            "formattedExcludes": d.FormattedExcludes,
                        },
        }            
}

func SerialzeAllDistributor(distributors []Distributor) map[string]interface{} {    
    var serializedDistributors []map[string]interface{}

    for _,d := range distributors {
        serializedDistributors = append(serializedDistributors, SerializeDistributor(d))
    }

    return map[string]interface{}{"data": serializedDistributors}
}

//Read the CSV file(Errors are not handles well since I know this is a valid CSV)
func SeedDBForUser(sessionId string){
    var userDB = DistributorDB[sessionId]

    if userDB == nil{
        DistributorDB[sessionId] = SampleDB()
    }    
}