package auxilary

import (
	"fmt"

	"github.com/souvikhaldar/challenge2016/file"
)

func FillSlice(size int, list *[]file.Csv) {
	var d file.Csv
	for i := 0; i < size; i++ {
		fmt.Printf("Seriel No.: %d \n", i+1)

		fmt.Println("Enter the city name:")
		fmt.Scanf("%s", &d.CityName)

		fmt.Println("Enter the province name:")
		fmt.Scanf("%s", &d.ProvinceName)

		fmt.Println("Enter the country name:")
		fmt.Scanf("%s", &d.CountryName)
		fmt.Println()

		*list = append(*list, d)
	}
}
