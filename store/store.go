package store

import (
	"strings"

	"github.com/challenge2016/models"
	"github.com/gin-gonic/gin"
)

type store struct{
	dMap *models.DistributionMaps
}

func NewStore(dMap *models.DistributionMaps) *store{
	return &store{
		dMap: dMap,
	}
}

func (s *store) AddDistributor(ctx *gin.Context,reqBody *models.Distributor) *models.Distributor{
	// storing a distributor in a map
	s.dMap.Distributor[reqBody.Name] = reqBody

	// get call

	return s.GetDistributorByName(ctx,reqBody.Name)
}

func (s *store) GetDistributorByName(ctx *gin.Context, distributorName string) *models.Distributor{
	distributorName = strings.ToUpper(distributorName)

	distributor ,ok := s.dMap.Distributor[distributorName]
	if !ok{
		return nil
	}

	return distributor
}

func (s *store) GetLocationDetailsByCity(cityName string) *models.Location{
	cityName = strings.ToUpper(cityName)

	loc , ok := s.dMap.CityMap[cityName]
	if !ok {
		return nil
	}

	return loc
}

func (s *store) GetLocationDetailsByProvince(province string) *models.Location{
	province = strings.ToUpper(province)

	loc , ok := s.dMap.ProvinceMap[province]
	if !ok {
		return nil
	}

	return loc
}

func (s *store) GetLocationDetailsByCountry(countryName string) *models.Location{
	countryName = strings.ToUpper(countryName)

	loc , ok := s.dMap.CountryMap[countryName]
	if !ok {
		return nil
	}

	return loc
}