package model

// Region - Structure to send the data through the channel
type Region struct {
	City     string
	Province string
	Country  string
}

// DataStore - to hold the cities, province and country
func DataStore(dataChannel chan Region, data map[string][]string, cities map[string][]string) {

	for {
		if rec, ok := <-dataChannel; ok {

			if province, ok := data[rec.Country]; ok {
				isPresent := false
				for _, v := range province {
					if rec.Province == v {
						isPresent = true
					}
				}
				if !isPresent {
					data[rec.Country] = append(province, rec.Province)
					cities[rec.Province] = []string{rec.City}
				} else {
					exisiting := cities[rec.Province]
					cities[rec.Province] = append(exisiting, rec.City)
				}

			} else {
				data[rec.Country] = []string{rec.Province}
				cities[rec.Province] = []string{rec.City}
			}

		} else {
			break
		}
	}

}
