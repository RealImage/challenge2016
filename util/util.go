package util

import (
	"strings"

	"github.com/challenge2016/model"
)

func IsAuthorized(distributorPermissions model.Permissions, region []string) bool {
	isFound := false
	for _, ex := range distributorPermissions.Exclude {
		if match(region, strings.Split(ex, "-")) {
			isFound = false
			break
		}
	}

	for _, inc := range distributorPermissions.Include {
		isFound = matchInclude(region, strings.Split(inc, "-"))
	}

	return isFound
}

func match(region, pattern []string) bool {
	if len(region) != len(pattern) {
		return false
	}

	for i := range region {
		if pattern[i] != region[i] {
			return false
		}
	}

	return true
}

func matchInclude(region, pattern []string) bool {
	i := len(region)
	for {
		if len(region) == len(pattern) {
			for i := range region {
				if pattern[i] != region[i] {
					return false
				}
			}
		}
		region = region[1:]
		i = i - 1
		if i > 0 {
			break
		}
	}
	return true
}
