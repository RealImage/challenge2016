package distribution

import (
	"fmt"
)

// sample User collection, can add more if required
var UserDataMap = map[string]User{
	"1": User{
		ID: "1",
		Permissions: []string{
			"INCLUDE: INDIA",
			"INCLUDE: UNITEDSTATES",
			"EXCLUDE: KARNATAKA-INDIA",
			"EXCLUDE: CHENNAI-TAMILNADU-INDIA",
		},
	},
	"2": User{
		ID:       "2",
		ParentID: "1",
		Permissions: []string{
			"INCLUDE: INDIA",
			"EXCLUDE: TAMILNADU-INDIA",
		},
	},
	"3": User{
		ID:       "3",
		ParentID: "2",
		Permissions: []string{
			"INCLUDE: SOMWARPET-KARNATAKA-INDIA",
		},
	},
	"4": User{
		ID:       "4",
		ParentID: "1",
		Permissions: []string{
			"INCLUDE: UNITEDSTATES",
			"INCLUDE: INDIA",
		},
	},
}

// getUser will return user by user id
func getUser(userID string) (User, error) {
	user, ok := UserDataMap[userID]
	if !ok {
		return user, fmt.Errorf("Invalid UserID:(%s)\n", userID)
	}
	return user, nil
}
