package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Readfile(path string, csvSlice *[]Csv) error {
	file, er := os.Open(path)
	if er != nil {
		log.Println(er)
		return er
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		det := strings.Split(scanner.Text(), ",")
		c := Csv{
			det[0],
			det[1],
			det[2],
			det[3],
			det[4],
			det[5],
		}
		*csvSlice = append(*csvSlice, c)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func GetParent(parentName string, distList []Distributor) Distributor {
	for _, dist := range distList {
		if dist.Name == parentName {
			return dist
		}
	}
	return Distributor{}
}
func (child *Distributor) appendExlist(parent Distributor) {
	for _, list := range parent.Exlist {
		child.Exlist = append(child.Exlist, list)
	}
}
func permErr(a, b string) string {
	return fmt.Sprintf("%s does not have permission for %s", a, b)
}

// CheckInclusion :child can't include regions that is not included in the parent
func (d Distributor) CheckInclusion(list []Distributor) string {
	parent := GetParent(d.ParentName, list)
	for _, ch := range d.InList {
		for _, pt := range parent.InList {
			if ch.CountryName == pt.CountryName {
				if ch.ProvinceName == pt.ProvinceName && pt.ProvinceName != "" {
					if ch.CityName == pt.CityName && pt.CityName != "" {
						return "Fine"
					} else if ch.CityName == pt.CityName && pt.CityName == "" {
						return "Fine"
					} else {
						return permErr(d.Name, ch.CityName)
					}
				} else if ch.ProvinceName == pt.ProvinceName && pt.ProvinceName == "" {
					return "Fine"
				} else {
					return permErr(d.Name, ch.ProvinceName)
				}
			} else {
				return permErr(d.Name, ch.CountryName)
			}
		}
	}
	return "Something Different"
}

// CheckExclusion: child can't include regions that is excluded in parent
func (d Distributor) CheckExclusion(list []Distributor) string {
	parent := GetParent(d.ParentName, list)
	for _, ch := range d.InList {
		for _, pt := range parent.Exlist {
			if ch.CountryName == pt.CountryName && ch.CountryName != "" {
				if ch.ProvinceName == pt.ProvinceName && ch.ProvinceName != "" {
					if ch.CityName == pt.CityName && ch.CityName != "" {
						return permErr(parent.Name, pt.CityName)
					} else if ch.CityName == pt.CityName && pt.CityName == "" {
						return permErr(parent.Name, pt.ProvinceName)
					} else {
						return "Fine"
					}
				} else if ch.ProvinceName == pt.ProvinceName && ch.ProvinceName == "" {
					return permErr(parent.Name, pt.CountryName)
				}
			} else if ch.CountryName == pt.CountryName && ch.CountryName == "" {
				return "Fine"
			} else {
				return "Fine"
			}
		}
	}
	return "Something Different"
}
