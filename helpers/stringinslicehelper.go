package helpers

/*Func to check whether a string is in a slice*/
func StringInSlice(str string, list [][]string) bool {
	for i := 0; i < len(list); i++ {
		for _, v := range list[i] {
			if v == str {
				return true
			}
		}
	}

	return false
}

func StringInSingleSlice(str string, list []string) bool {
		for _, v := range list {
			if v == str {
				return true
			}
		}

	return false
}
