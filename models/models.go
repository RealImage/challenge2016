package models

type Permissions interface {
	GetName() string
	GetPermittedPlaces() []string
	GetRestrictedPlaces() []string
	GetAuthorizedRegions() map[string]bool
	GetParent() Permissions
}

type Distributor struct {
	Name              string
	PermittedPlaces   []string
	RestrictedPlaces  []string
	AuthorizedRegions map[string]bool
	Parent            Permissions // Reference to the parent distributor
}

func (d Distributor) GetName() string {
	return d.Name
}

func (d Distributor) GetPermittedPlaces() []string {
	return d.PermittedPlaces
}

func (d Distributor) GetRestrictedPlaces() []string {
	return d.RestrictedPlaces
}

func (d Distributor) GetAuthorizedRegions() map[string]bool {
	return d.AuthorizedRegions
}

func (d Distributor) GetParent() Permissions {
	return d.Parent
}
