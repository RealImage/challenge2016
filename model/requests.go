package model

// Permission - Contains permissions given to the distributor
type Permission struct {
	For         string     `json:"for"`
	From        string     `json:"from"`
	Includes    []string   `json:"includes"`
	Excludes    []Excluded `json:"excludes"`
	SubIncludes []Included `json:"sub_includes"`
}

// IsAuthorized - Contains info to check authorization of the distributor
type IsAuthorized struct {
	For      string `json:"for"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}

//Excluded - Contains excluded city, province and country for the distributor
type Excluded struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Country  string `json:"country,omitempty"`
}

//Included - Contains excluded city, province and country for the distributor
type Included struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Country  string `json:"country,omitempty"`
}

//SubPermission - Contains permissions for sub distributor
type SubPermission struct {
	For      string     `json:"for"`
	From     string     `json:"from"`
	Includes []Included `json:"includes"`
	Excludes []Excluded `json:"excludes"`
}
