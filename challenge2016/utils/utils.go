package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func GetMainMenu() {
	fmt.Println("                                      ")
	fmt.Println("1. Add Permission for Distributor")
	fmt.Println("2. List of all Valid distributors")
	fmt.Println("3. Validate the permission of distributor")
	fmt.Println("4. Create a Network of Distributors")
	fmt.Println("5. Main Menu")
	fmt.Println("6. Abort")
	fmt.Println("                                      ")
}

// RemoveSpace removes the spaces from the string and returns the string in lower case
func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}

	return strings.ToLower(string(rr))
}

// Contains checks if the given string is present in the given slice
func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}
