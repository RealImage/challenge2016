package util

import (
	"fmt"
)

func NewAuthority(Name string) Authority {
	return Authority{Name: Name}
}

func (a *Authority) RegisterDistributor(d Distributor) {
	if a.Distributors == nil {
		a.Distributors = make(map[string]*Distributor)
	}
	if _, ok := a.Distributors[d.Name]; !ok {
		a.Distributors[d.Name] = &Distributor{Name: d.Name}
		fmt.Printf("%v is registered\n", d.Name)
		return

	}
	fmt.Printf("%v is already registered\n", d.Name)
}

func (a *Authority) GetDistributorByName(name string) (*Distributor, error) {
	if a.Distributors == nil {
		return nil, fmt.Errorf("No distributor exists [%s]\n", name)
	}
	if _, ok := a.Distributors[name]; ok {
		return a.Distributors[name], nil
	} else {
		return nil, fmt.Errorf("No distributor is found with the name [%s]\n", name)
	}

}

// Before calling this check compatibility with old distributor
func (a Authority) RegisterDistributorByDist(new_distributor, old_distributor *Distributor) error {

	if _, ok := a.Distributors[old_distributor.Name]; !ok {
		return fmt.Errorf("%v is not a authorized distributor\n", old_distributor)
	}
	a.Distributors[new_distributor.Name] = &Distributor{}
	return nil
}

func (d *Distributor) CheckPermissionForEntireCountry(country Country) bool {
	if d.Permission.CountryPermission.Allowed == nil {
		return false
	}
	if provincePermission, ok := d.Permission.CountryPermission.Allowed[country]; ok {
		// if there is any single province in the province map then return false
		if provincePermission.NotAllowed == nil || len(provincePermission.NotAllowed) == 0 {
			if provincePermission.Allowed == nil || len(provincePermission.Allowed) == 0 {
				return true
			} else {
				//range all the province and if any of the country is not allowed
				for _, cityPermission := range provincePermission.Allowed {
					if cityPermission.NotAllowed != nil && len(cityPermission.NotAllowed) > 0 {
						return false
					}
				}
				return true
			}
		}
		return false
	} else {
		return false
	}

}

func (d *Distributor) CheckPermissionForEntireProvince(country Country, province Province) bool {
	countryPermission := d.Permission.CountryPermission.Allowed
	//country itself is not added
	if countryPermission == nil {
		return false
	}

	provincePermission := d.Permission.CountryPermission.Allowed[country]
	// if country is there and that is not having allowed and not allowed
	// province which means this stats is also having permission as whole counter is permitted
	if provincePermission.Allowed == nil && provincePermission.NotAllowed == nil {
		return true
	}

	// if stats itself is part of not allowed for that country
	if provincePermission.NotAllowed != nil {
		if _, ok := provincePermission.NotAllowed[province]; ok {
			return false
		}
	}

	// if one city is there which is not a part of allowed and also which is not part of not allowed
	// then thats city is also said to allowed
	if provincePermission.Allowed != nil {
		for _, cityPermission := range provincePermission.Allowed {
			if cityPermission.NotAllowed != nil && len(cityPermission.NotAllowed) > 0 {
				return false
			}
		}
	}
	return true
}

func (d *Distributor) CheckPermissionForCity(country Country, province Province, city City) bool {
	if d.Permission.CountryPermission.Allowed == nil || len(d.Permission.CountryPermission.Allowed) == 0 {
		return false
	}
	// fmt.Printf("province %v", d.Permission.CountryPermission.Allowed[country])
	// _, ok := d.Permission.CountryPermission.Allowed[country]
	// fmt.Printf("province exists [%v]", ok)
	if provincePermission, ok := d.Permission.CountryPermission.Allowed[country]; ok {
		if provincePermission.NotAllowed != nil {
			if _, ok := provincePermission.NotAllowed[province]; ok {
				return false
			}
		}

		// it means it allowed to the all province
		if provincePermission.Allowed == nil {
			return true
		}

		if cityPermission, ok := provincePermission.Allowed[province]; ok {
			if cityPermission.NotAllowed != nil {
				if _, ok := cityPermission.NotAllowed[city]; ok {
					return false
				}
			}
			// it means all cities of this province is allowed
			return true
			// if cityPermission.Allowed == nil && cityPermission.NotAllowed == nil {
			// 	fmt.Println("debug5")
			// 	return true
			// }
			// if _, ok := cityPermission.Allowed[city]; ok {
			// 	fmt.Println("debug6")
			// 	return true
			// } else {
			// 	fmt.Println("debug7")
			// 	return false
			// }

		} else {
			return false
		}

	} else {
		return false
	}
}

func (a *Authority) AddNewPermissionForEntireCountry(d *Distributor, country Country) error {
	if _, ok := a.Distributors[d.Name]; !ok {
		return fmt.Errorf("Distributor [%v] is not authorised", d.Name)
	}
	if d.Permission.CountryPermission.Allowed == nil {
		d.Permission.CountryPermission.Allowed = make(map[Country]ProvincePermission)
		d.Permission.CountryPermission.Allowed[country] = ProvincePermission{}
		fmt.Printf("Permission for %v added for Country [%v] \n", d.Name, country.Name)
		return nil
	}

	if _, ok := d.Permission.CountryPermission.Allowed[country]; ok {
		fmt.Printf("Warning: Country [%v] is already added for %v \n", country.Name, d.Name)
		return nil
	}
	d.Permission.CountryPermission.Allowed[country] = ProvincePermission{}
	fmt.Printf("Permission for %v added for Country [%v] \n", d.Name, country.Name)
	return nil
}

func (a *Authority) AddNewPermissionForEntireProvince(d *Distributor, province Province, country Country, permissionType string) error {
	// err := a.AddNewPermissionForEntireCountry(d, country)
	// If country itself is not having any permission
	// then just come back
	if d.Permission.CountryPermission.Allowed == nil {
		return fmt.Errorf("Country [%v] is not in included list,  First include it\n", country.Name)
	}

	if _, ok := d.Permission.CountryPermission.Allowed[country]; !ok {
		return fmt.Errorf("Country [%v] is not in included list,  First include it\n", country.Name)
	}

	//
	p, _ := d.Permission.CountryPermission.Allowed[country]

	if permissionType == "INCLUDE" {
		//Already present in block
		//then first remove it from blocked list
		if p.NotAllowed != nil {
			if _, ok := p.NotAllowed[province]; ok {
				delete(p.NotAllowed, province)
			}
		}
		if p.Allowed == nil {
			p.Allowed = make(map[Province]CityPermission)
			p.Allowed[province] = CityPermission{}
			d.Permission.CountryPermission.Allowed[country] = p
			fmt.Printf("New Permission added for province [%v].\n", province.Name)
			return nil
		}
		//Already present
		if _, ok := p.Allowed[province]; ok {
			fmt.Printf("Already permission present for province [%v].\n", province.Name)
			return nil
		}
		// if not present
		p.Allowed[province] = CityPermission{}
		fmt.Printf("New Permission added for province [%v].\n", province.Name)
		return nil
	} else {
		//Already present in allow
		//then first remove it from allowed list
		if p.Allowed != nil {
			if _, ok := p.Allowed[province]; ok {
				delete(p.Allowed, province)
			}
		}
		if p.NotAllowed == nil {
			p.NotAllowed = make(map[Province]int)
			p.NotAllowed[province] = 1
			d.Permission.CountryPermission.Allowed[country] = p
			fmt.Printf("New Permission added for province [%v].\n", province.Name)
			return nil
		}
		//Already blocked
		if _, ok := p.NotAllowed[province]; ok {
			fmt.Printf("Province [%s] already blocked.\n", province.Name)
			return nil
		}
		// if not blocked, now block
		p.NotAllowed[province] = 1
		fmt.Printf("Province [%v] is now blocked. \n", province.Name)

		return nil
	}
}

func (a *Authority) AddNewPermissionForCity(d *Distributor, city City, province Province, country Country, permissionType string) error {

	// If country itself is not having any permission
	// then just come back
	if d.Permission.CountryPermission.Allowed == nil {
		return fmt.Errorf("Country [%v] is not in included list,  First include it\n", country.Name)
	}
	//if country is not added
	if _, ok := d.Permission.CountryPermission.Allowed[country]; !ok {
		return fmt.Errorf("Country [%v] is not in included list,  First include it\n", country.Name)
	}

	provincePermission := d.Permission.CountryPermission.Allowed[country]
	// Stats itself not added
	if provincePermission.Allowed == nil {
		return fmt.Errorf("province [%v] is not in included list,  First include it\n", province.Name)
	}
	// stats is not added in the allowed list
	if _, ok := provincePermission.Allowed[province]; !ok {
		return fmt.Errorf("province [%v] is not in included list,  First include it\n", province.Name)
	}

	if provincePermission.NotAllowed != nil {
		if _, ok := provincePermission.NotAllowed[province]; ok {
			return fmt.Errorf("province [%v] is already blocked entirely \n", province.Name)
		}
	}
	c := d.Permission.CountryPermission.Allowed[country].Allowed[province]

	if permissionType == "INCLUDE" {
		// if this city is already excluded
		if c.NotAllowed != nil {
			if _, ok := c.NotAllowed[city]; ok {
				fmt.Printf("Deleting province [%v] from excluded list. \n", province.Name)
				delete(c.NotAllowed, city)
			}
		}
		if c.Allowed == nil {
			c.Allowed = make(map[City]int)
			c.Allowed[city] = 1
			d.Permission.CountryPermission.Allowed[country].Allowed[province] = c
			fmt.Printf("New Permission added for City [%v] \n", city.Name)
			return nil
		}
		//Already present
		if _, ok := c.Allowed[city]; ok {
			fmt.Printf("Permission already added for city [%v]\n", city.Name)
			return nil
		}
		// if not present
		c.Allowed[city] = 1
		fmt.Printf("New Permission added for City [%v] \n", city.Name)
		return nil
	} else {
		//Already present in allow
		//then first remove it from allowed list
		if c.Allowed != nil {
			if _, ok := c.Allowed[city]; ok {
				delete(c.Allowed, city)
			}
		}
		if c.NotAllowed == nil {
			c.NotAllowed = make(map[City]int)
			c.NotAllowed[city] = 1
			d.Permission.CountryPermission.Allowed[country].Allowed[province] = c
			fmt.Printf("City [%v] is blocked succesfuly\n", city.Name)
			return nil
		}
		//Already blocked
		if _, ok := c.NotAllowed[city]; ok {
			fmt.Printf("City [%v]  is already blocked \n", city.Name)
			return nil
		}
		// if not present
		c.NotAllowed[city] = 1
		fmt.Printf("City [%v] is blocked succesfuly\n\n", city.Name)
		return nil
	}
}

func (a *Authority) CheckCompatibilityForCountry(d_child, d_parent *Distributor, country Country) error {
	if _, ok := a.Distributors[d_parent.Name]; !ok {
		return fmt.Errorf("Distributor [%v] is not authorised\n", d_parent)
	}
	if _, ok := a.Distributors[d_child.Name]; !ok {
		return fmt.Errorf("Distributor [%v] is not authorised\n", d_child.Name)
	}

	c := d_parent.Permission.CountryPermission

	if c.Allowed == nil || len(c.Allowed) == 0 {
		return fmt.Errorf("parent_distributor [%v] is not having any permissions\n", d_parent.Name)
	}
	if _, ok := c.Allowed[country]; !ok {
		return fmt.Errorf("parent distributor [%v] is not having permissions for country[%v]\n", d_parent.Name, country.Name)

	}
	return nil
}

func (a *Authority) CheckCompatibilityForProvince(d_child, d_parent *Distributor, province Province, country Country, pType string) error {

	err := a.CheckCompatibilityForCountry(d_child, d_parent, country)

	if err != nil {
		return err
	}

	c := d_parent.Permission.CountryPermission
	p := c.Allowed[country]
	if pType == "INCLUDE" {
		// for that country there is no province added that means entire country is available
		if p.Allowed == nil && p.NotAllowed == nil {
			return nil
		}
		// if no allowed province
		if p.Allowed == nil {
			// if that province belonged to one of the province
			if _, ok := p.NotAllowed[province]; ok {
				return fmt.Errorf("Unable to give Include Permission to a Exclude province [%v]\n", province.Name)
			}
		}
		// if it doesnt belongs from the allowed province but there in not allowed map then return with error
		if c, ok := p.Allowed[province]; !ok {
			if p.NotAllowed != nil {
				if _, ok := p.NotAllowed[province]; ok {
					return fmt.Errorf("Unable to give Include Permission to a Exclude province [%v]\n", province.Name)
				}
			}
			return nil
		} else {
			// if it belongs from the allowed stats but have some blocked city
			if c.NotAllowed != nil && len(c.NotAllowed) > 0 {
				return fmt.Errorf("Unable to give Include Permission to a province [%v], which is having some exclude cities", province.Name)
			}
			// if no blocked city
			return nil
		}
	} else {
		// in case of blocking we can always block.
		return nil
	}
}

func (a *Authority) CheckCompatibilityForCity(d_child, d_parent *Distributor, city City, province Province, country Country, pType string) error {

	// if entire province there then we can always include that
	if pType == "INCLUDE" {
		if err := a.CheckCompatibilityForProvince(d_child, d_parent, province, country, pType); err == nil {
			return nil
		} else if err.Error() == fmt.Sprintf("Some cities are blocked in Province[%v]\n", province.Name) {
			if _, ok := d_parent.Permission.CountryPermission.Allowed[country].Allowed[province].NotAllowed[city]; ok {
				return fmt.Errorf("This city[%v] is already blocked In this Province[%v]\n", province.Name, city.Name)
			} else {
				return nil
			}
		} else {
			return err
		}
	} else {
		return a.CheckCompatibilityForCountry(d_child, d_parent, country)
	}
}

func (a *Authority) IsCompatibeForNewPermissionForCountry(dList []string, country string) bool {
	d_child, err := a.GetDistributorByName(dList[0])
	if err != nil {
		fmt.Printf("[%s] is not a valid distributor. [%v]\n", d_child.Name, err.Error())
		return false
	}
	if len(dList) > 1 {
		d_ancestor, err := a.GetDistributorByName(dList[1])
		if err != nil {
			fmt.Printf("[%s] is not a valid distributor. [%v]\n", d_ancestor.Name, err.Error())
			return false
		}
		return d_ancestor.CheckPermissionForEntireCountry(Country{Name: country})
	}
	return true
}

func (a *Authority) IsCompatibeForNewPermissionForProvince(dList []string, province, country, pType string) bool {
	d_child, err := a.GetDistributorByName(dList[0])
	if err != nil {
		fmt.Printf("[%s] is not a valid distributor\n", d_child.Name)
		return false
	}
	for i := 1; i < len(dList); i++ {
		d_ancestor, err := a.GetDistributorByName(dList[i])
		if err != nil {
			fmt.Printf("[%s] is not a valid distributor\n", d_child.Name)
			return false
		}
		err = a.CheckCompatibilityForProvince(d_child, d_ancestor, Province{Name: province}, Country{Name: country}, pType)
		if err != nil {
			return false
		}
	}
	return true
}

func (a *Authority) IsCompatibeForNewPermissionForCity(dList []string, city, province, country, pType string) bool {
	d_child, err := a.GetDistributorByName(dList[0])
	if err != nil {
		fmt.Printf("[%s] is not a valid distributor\n", d_child.Name)
		return false
	}
	for i := 1; i < len(dList); i++ {
		d_ancestor, err := a.GetDistributorByName(dList[i])
		if err != nil {
			fmt.Printf("[%s] is not a valid distributor\n", d_child.Name)
			return false
		}
		err = a.CheckCompatibilityForCity(d_child, d_ancestor, City{Name: city}, Province{Name: province}, Country{Name: country}, pType)
		if err != nil {
			return false
		}
	}
	return true
}
