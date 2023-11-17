package controller

import (
	"net/http"
	datacsv "qube-cinemas-challenge/data-csv"
	"qube-cinemas-challenge/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"status":true,"Locations": datacsv.Cities})
}

//Distributor Management
func GetDistributor(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"status":true, "distributors":datacsv.Distributor})
}

func AddDistributor(c *gin.Context){
	newDistributor := &models.Distributor{}

	type DistributorData struct{
		Parent string `json:"parent"`
	}
	var newData DistributorData
	if err := c.ShouldBind(&newData);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"error":err.Error()})
		return
	}
	var exist bool
	if newData.Parent != ""{
		for _, distributor:= range datacsv.Distributor{
			exist = distributor.ID==newData.Parent
			if distributor.ID==newData.Parent{
				newDistributor.Parent = distributor
				break
			} 
		}
		if !exist{
			c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"Distributor id didn't exist"})
			return
		}
	}
	newDistributor.ID = strconv.Itoa(len(datacsv.Distributor)+1)
	datacsv.Distributor = append(datacsv.Distributor, newDistributor)

	c.JSON(http.StatusAccepted, gin.H{"status":true, "message":"New distributor created with id "+ newDistributor.ID})
}

func GetParentDetails(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	for _,distributor := range datacsv.Distributor{
		if distributor.ID == dist.Id {
			c.JSON(http.StatusAccepted, gin.H{"status":true, "Parent":distributor.Parent})
			return
		}
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{"status":false, "message":"Distributor id doesnot exits"})

}

func GetSubDistributors(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var subDistributors []*models.Distributor
	for _,distributor := range datacsv.Distributor {
		if distributor != nil && distributor.Parent != nil && distributor.Parent.ID == dist.Id {
			subDistributors = append(subDistributors, distributor)
		}
	}
	c.JSON(http.StatusAccepted, gin.H{"status":true, "Sub-Distributors":subDistributors})
}

//Permission management
func GetIncludedRegion(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var includedCity []models.City
	for _,rule:=range datacsv.Rules {
		if rule.Dist_Id == dist.Id{
			for region,val:= range rule.Included{
				if val {
					includedCity = append(includedCity, *region)
				}
			}
		}
	}
	c.JSON(http.StatusAccepted, gin.H{"status":true, "Included-City":includedCity})
}

func AddIncludedCountry(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CountryCode string `json:"country-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Province.Country.Code == dist.CountryCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = true
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

func AddIncludedProvince(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		ProvinceCode string `json:"province-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Province.Code == dist.ProvinceCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = true
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

func AddIncludedCity(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CityCode string `json:"city-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Code == dist.CityCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = true
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

func RemoveIncludedCountry(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CountryCode string `json:"country-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Province.Country.Code == dist.CountryCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = false
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

func RemoveIncludedProvince(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		ProvinceCode string `json:"province-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Province.Code == dist.ProvinceCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = false
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

func RemoveIncludedCity(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CityCode string `json:"city-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var dist_rule *models.Rule
	var exist bool
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			exist = true
			dist_rule = rule
		}
	}
	for _,distributor:= range datacsv.Distributor{
		if distributor.ID ==dist.Id{
			for _, city := range datacsv.Cities{
				if city.Code == dist.CityCode{
					includeNode := make(map[*models.City]bool)
					includeNode[city] = false
					if !exist{
						dist_rule = &models.Rule{Dist_Id: dist.Id, Included: includeNode}
						datacsv.Rules = append(datacsv.Rules, dist_rule)
					} else {
						dist_rule.Included = includeNode
					}
					
				}
			}
		}
	}
}

//Checking Permission
func CityLevelPermission(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CityCode string `json:"city-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var city *models.City
	var exist bool
	for _,getCity := range datacsv.Cities{
		if getCity.Code == dist.CityCode{
			exist = true
			city = getCity
		}
	}
	if !exist{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"The city code doesn't exist"})
		return
	}
	for _,rule:= range datacsv.Rules{
		if rule.Dist_Id == dist.Id{
			if rule.Included[city]{
				c.JSON(http.StatusAccepted, gin.H{"status":true, "message":city.Name+" is permitted for Distributor "+dist.Id})
				return
			} else {
				c.JSON(http.StatusAccepted, gin.H{"status":true, "message":city.Name+" is not permitted for Distributor "+dist.Id})
				return
			}
		}
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"Distributor Id doesn't exist"})
}

func ProvinceLevelPermission(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		ProvinceCode string `json:"province-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var province *models.Province
	var exist bool
	var included, excluded int
	for _,getCity := range datacsv.Cities{
		if getCity.Province.Code == dist.ProvinceCode{
			province = getCity.Province
			exist = true
			for _,rule:= range datacsv.Rules{
				if rule.Dist_Id == dist.Id{
					if rule.Included[getCity]{
						included++
					} else {
						excluded++
					}
				}
			}
		}
	}
	if !exist{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"The province code doesn't exist"})
		return
	}

	if excluded == 0 {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":province.Name+" is permitted for Distributor "+dist.Id})
		return
	} else if included !=0 {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":province.Name+" is permitted for Distributor "+dist.Id+" but there are "+strconv.Itoa(excluded)+" cities excluded for this distributor"})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":province.Name+" is not permitted for Distributor "+dist.Id})
		return
	}
}

func CountryLevelPermission(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
		CountryCode string `json:"city-code"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var country *models.Country
	var exist bool
	var included, excluded int
	for _,getCity := range datacsv.Cities{
		if getCity.Province.Country.Code == dist.CountryCode{
			country = getCity.Province.Country
			exist = true
			for _,rule:= range datacsv.Rules{
				if rule.Dist_Id == dist.Id{
					if rule.Included[getCity]{
						included++
					} else {
						excluded++
					}
				}
			}
		}
	}
	if !exist{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"The country code doesn't exist"})
		return
	}

	if excluded == 0 {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":country.Name+" is permitted for Distributor "+dist.Id})
		return
	} else if included !=0 {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":country.Name+" is permitted for Distributor "+dist.Id+" but there are "+strconv.Itoa(excluded)+" cities excluded for this distributor"})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status":true, "message":country.Name+" is not permitted for Distributor "+dist.Id})
		return
	}
}