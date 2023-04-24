package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func GetMainMenu() {
	fmt.Println("1. Add Distributor with Permission")
	fmt.Println("2. List all Distributors")
	fmt.Println("3. Check Distributor from the distributor list")
	fmt.Println("4. Check Permission for a Distributor")
	fmt.Println("5. Main Menu")
	fmt.Println("6. Exit")
}

func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}

	return strings.ToLower(string(rr))
}
