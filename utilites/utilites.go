package utilites

import (
	"strings"
)

func UpperCaseNoSpace(input string) string {
	return strings.ToUpper(strings.Replace(input, " ", "", -1))
}
