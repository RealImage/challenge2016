package main

type Permission struct {
	Includes map[string]bool
	Excludes map[string]bool
}

type Distributor struct {
	Name        string
	Permissions map[string]Permission
}

func NewDistributor(name string) *Distributor {
	return &Distributor{
		Name:        name,
		Permissions: make(map[string]Permission),
	}
}

func (d *Distributor) AddPermission(region, action string) {
	permission := d.Permissions[d.Name]
	if action == "INCLUDE" {
		permission.Includes[region] = true
	} else if action == "EXCLUDE" {
		permission.Excludes[region] = true
	}
	d.Permissions[d.Name] = permission
}

func (d *Distributor) HasPermission(region, action string) bool {
	permission, ok := d.Permissions[d.Name]
	if !ok {
		return false
	}

	if action == "INCLUDE" {
		return permission.Includes[region]
	} else if action == "EXCLUDE" {
		return permission.Excludes[region]
	}

	return false
}
