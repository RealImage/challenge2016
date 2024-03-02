package utils

import "fmt"

func Println(v ...interface{}) {
	fmt.Printf("[DEBUG]:%+v\n", v)
}

func Error(v ...interface{}) {
	fmt.Printf("[ERROR]:%+v\n", v)
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
