package utils

type Permission struct {
	Included []Region `json:"included"`
	Excluded []Region `json:"excluded"`
}

// Region represents a geographic region.
type Region struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

type NewDistributer struct {
	Name    string
	Include []string
	Exclude []string
	Check   string
	Sub     []NewDistributer
}
