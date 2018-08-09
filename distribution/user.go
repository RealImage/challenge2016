package distribution

import (
	"fmt"
	"strings"
)

// User Struct to store user permissions
type User struct {
	ID          string
	Permissions []string
	ParentID    string
}

var UniversalAreaList []Area

// parses user permissions and returns a permission map
func (u *User) ParsePermission() (map[string]bool, error) {

	// the below maps are permission maps based on the area level,
	// state level , and country level
	var areaLevel = make(map[string]bool)
	var stateLevel = make(map[string]bool)
	var countryLevel = make(map[string]bool)
	var parentPermissions = make(map[string]bool)

	var allPermsissions = make(map[string]bool)

	// check parent user permissions to follow the permission inheritence
	if parentUser, err := GetUser(u.ParentID); err == nil {
		parentPermissions, err = parentUser.ParsePermission()
		if err != nil {
			return nil, err
		}
	} else {
		parentPermissions = nil
	}

	for _, p := range u.Permissions {
		splitPerm := strings.Split(p, ":")
		switch splitPerm[0] {
		// Allowed permissions
		case "INCLUDE":
			area, state, country := splitLocationText(splitPerm[1])
			if state != "" && country != "" && area != "" {
				areaLevel[fmt.Sprintf("%s-%s-%s", area, state, country)] = true
			}
			if state != "" && country != "" {
				stateLevel[fmt.Sprintf("%s-%s", state, country)] = true
			}
			if country != "" {
				countryLevel[country] = true
			}
		// Permission denied locations
		case "EXCLUDE":
			area, state, country := splitLocationText(splitPerm[1])

			if state != "" && country != "" && area != "" {
				areaLevel[fmt.Sprintf("%s-%s-%s", area, state, country)] = false
				continue
			}
			if state != "" && country != "" {
				stateLevel[fmt.Sprintf("%s-%s", state, country)] = false
				continue
			}
			if country != "" {
				countryLevel[country] = false
			}

		default:
			fmt.Printf("Invalid authorisations")
		}
	}

	// Check child permissions with perent user
	// if the permissions clash, return error denoting invalid user
	// test only for permissions not allowed in parent user
	if parentPermissions != nil {

		for kArea, vArea := range areaLevel {
			v, _ := parentPermissions[kArea]
			if !v && vArea {
				return nil, fmt.Errorf("Invalid Parent (%+s) for user (%+s)", u.ParentID, u.ID)
			}
		}

		for kState, vState := range stateLevel {
			v, _ := parentPermissions[kState]
			if !v && vState {
				return nil, fmt.Errorf("Invalid permissions in child (%+s) for Parent (%+s)", u.ID, u.ParentID)
			}
		}

		for kCountry, vCountry := range countryLevel {
			v, _ := parentPermissions[kCountry]

			if !v && vCountry {
				return nil, fmt.Errorf("Invalid permissions in child (%+s) for Parent (%+s)", u.ID, u.ParentID)
			}
		}
	}
	// create permission map for user
	for _, a := range UniversalAreaList {
		valueCountry, okCountry := countryLevel[a.Country]

		if okCountry && valueCountry {
			allPermsissions[a.Country] = true
			allPermsissions[fmt.Sprintf("%s-%s", a.State, a.Country)] = true
			allPermsissions[fmt.Sprintf("%s-%s-%s", a.Area, a.State, a.Country)] = true
		}
		valueState, okState := stateLevel[fmt.Sprintf("%s-%s", a.State, a.Country)]
		if okState && !valueState {
			allPermsissions[fmt.Sprintf("%s-%s-%s", a.Area, a.State, a.Country)] = false
			allPermsissions[fmt.Sprintf("%s-%s", a.State, a.Country)] = false
		}
		valueArea, okArea := stateLevel[fmt.Sprintf("%s-%s-%s", a.Area, a.State, a.Country)]
		if okArea && !valueArea {
			allPermsissions[fmt.Sprintf("%s-%s-%s", a.Area, a.State, a.Country)] = false
		}
	}
	// revoke permissions where parent user has no access
	if parentPermissions != nil {
		for k, v := range allPermsissions {
			vp, _ := parentPermissions[k]
			if v && !vp {
				allPermsissions[k] = vp
			}
		}
	}

	return allPermsissions, nil
}

// Splits the string permissions into area , state and country
func splitLocationText(s string) (string, string, string) {
	asc := strings.Split(strings.ToLower(strings.TrimSpace(s)), "-")
	switch len(asc) {
	case 1:
		return "", "", asc[0]
	case 2:
		return "", asc[0], asc[1]
	case 3:
		return asc[0], asc[1], asc[2]
	default:
		return "", "", ""
	}
}
