package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in Marshalling data: ", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}

func respondError(w http.ResponseWriter, status int, err string) {
	respondJSON(w, status, errorMessage{err})
}

func getPasswordHash(password string) ([]byte, error) {

	if len(password) == 0 {
		return nil, errors.New(emptyPassword)
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return hashBytes, nil

}

func isAlreadyLoggedIn(req *http.Request) (bool, credential, error) {

	if len(req.Cookies()) == 0 {
		return false, credential{}, nil
	}
	authToken, err := getCookieValue(req, authenticationToken)
	if err != nil {
		return false, credential{}, err
	}

	creds, ok := sessionObject.getFromSessionMap(authToken)
	return ok, creds, nil
}

func getCookieValue(req *http.Request, cookieName string) (string, error) {

	authCookie, err := req.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	return authCookie.Value, nil
}

func isValidCreds(inputCreds *credential) bool {
	if creds, ok := credentialsObject.getFromCredentialMap(inputCreds.Username); ok {
		inputCreds.EncryptedPassword = creds.EncryptedPassword
		return true
	}
	return false
}

func getUserFromUsers(username string) *user {
	for _, u := range users {
		if u.Name == username {
			return u
		}
	}

	for _, user := range users {
		if outUser := getUserFromUsersHelper(user, username); outUser != nil {
			return outUser
		}
	}

	return nil
}

func getUserFromUsersHelper(inUser *user, username string) *user {

	for _, child := range inUser.Children {
		if child.Name == username {
			return child
		}
		if len(child.Children) == 0 {
			return nil
		}
		if outUser := getUserFromUsersHelper(child, username); outUser != nil {
			return outUser
		}

	}

	return nil

}

func validateLocation(loc location) error {
	if loc.CountryName == "" {
		return errors.New(countryNameCannotBeEmpty)
	}
	if loc.ProvinceName == "" && loc.CityName != "" {
		return errors.New(cannotGiveCityWithoutProvince)
	}

	if loc.CityName == "" {
		return errors.New(countryProvinceCityCannotBeEmpty)
	}
	return nil
}

func isValidLocation(inCheckUser *user, inLoc location) bool {

	coun, err := getCountry(inLoc.CountryName, inLoc.ProvinceName, inLoc.CityName)
	if err != nil {
		return false
	}

	if inCheckUser == nil {
		return true
	}

	checkUserCountries, err := getUserCountries(inCheckUser)
	if err != nil {
		return false
	}

	// return !appendToUserCountries(checkUserCountries, c, false)

	_, c, cok := getCountryFromCountries(checkUserCountries, coun.Name)
	if !cok {
		return false
	}
	if len(coun.Provinces) == 1 {
		_, p, pok := getProvinceFromCountry(c, coun.Provinces[0].Name)
		if !pok {
			return false
		}

		if len(coun.Provinces[0].Cities) == 1 {
			_, _, ciok := getCityFromProvince(p, coun.Provinces[0].Cities[0].Name)
			if !ciok {
				return false

			}
		}
	}
	return true

}

func getUserCountries(u *user) ([]*country, error) {

	userCountries := make([]*country, 1)

	for _, loc := range u.Includes {
		coun, err := getCountry(loc.CountryName, loc.ProvinceName, loc.CityName)
		if err != nil {
			return nil, err
		}

		// appendToUserCountries(userCountries, c, true)

		_, c, cok := getCountryFromCountries(userCountries, coun.Name)
		if !cok {
			userCountries = append(userCountries, c)
			continue
		}
		if len(coun.Provinces) == 1 {
			_, p, pok := getProvinceFromCountry(c, coun.Provinces[0].Name)
			if !pok {
				c.Provinces = append(c.Provinces, coun.Provinces[0])
				continue
			}

			if len(coun.Provinces[0].Cities) == 1 {
				_, _, ciok := getCityFromProvince(p, coun.Provinces[0].Cities[0].Name)
				if !ciok {
					p.Cities = append(p.Cities, coun.Provinces[0].Cities[0])

				}
			}
		}

	}

	for _, loc := range u.Excludes {
		c, err := getCountry(loc.CountryName, loc.ProvinceName, loc.CityName)
		if err != nil {
			return nil, err
		}

		removeFromUserCountries(userCountries, c)

	}

	return userCountries, nil

}

// func appendToUserCountries(inUserCountries []*country, inCountry country, append bool) bool {

// 	_, c, cok := getCountryFromCountries(inUserCountries, inCountry.Name)
// 	if !cok {
// 		if append {

// 			inUserCountries = append(inUserCountries, inCountry)
// 		}
// 		return true
// 	}

// 	if len(inCountry.Provinces) == 1 {
// 		_, p, pok := getProvinceFromCountry(c, inCountry.Provinces[0].Name)
// 		if !pok {
// 			if append {
// 				c.Provinces = append(c.Provinces, inCountry.Provinces[0])
// 			}
// 			return true
// 		}

// 		if len(inCountry.Provinces[0].Cities) == 1 {

// 			_, _, ciok := getCityFromProvince(p, inCountry.Provinces[0].Cities[0].Name)
// 			if !ciok {
// 				if append {
// 					p.Cities = append(p.Cities, inCountry.Provinces[0].Cities[0])
// 				}
// 				return true
// 			}
// 		}
// 	}
// 	return false

// }

func removeFromUserCountries(inUserCountries []*country, inCountry country) {

	i, c, cok := getCountryFromCountries(inUserCountries, inCountry.Name)
	if cok {
		if len(inCountry.Provinces) == 0 {
			inUserCountries = append(inUserCountries[:i], inUserCountries[:i+1]...)
			return
		}

		j, p, pok := getProvinceFromCountry(c, inCountry.Provinces[0].Name)
		if pok {
			if len(inCountry.Provinces[0].Cities) == 0 {
				c.Provinces = append(c.Provinces[:j], c.Provinces[:j+1]...)
				return
			}

			k, _, ciok := getCityFromProvince(p, inCountry.Provinces[0].Cities[0].Name)
			if ciok {
				p.Cities = append(p.Cities[:k], p.Cities[:k+1]...)
			}

		}
	}
	return
}

func getCountryFromCountries(inCountries []*country, countryName string) (int, *country, bool) {

	for i, c := range inCountries {
		if c.Name == countryName {
			return i, c, true
		}
	}

	return 0, nil, false

}

func getProvinceFromCountry(inCountry *country, provinceName string) (int, *province, bool) {
	if inCountry != nil {

		for i, p := range inCountry.Provinces {
			if p.Name == provinceName {
				return i, p, true
			}
		}
	}

	return 0, nil, false
}

func getCityFromProvince(inProvince *province, cityName string) (int, *city, bool) {

	if inProvince == nil {

		for i, ci := range inProvince.Cities {
			if ci.Name == cityName {
				return i, ci, true
			}
		}
	}

	return 0, nil, false
}
