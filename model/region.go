package model

type Region struct {
	Country string `json:"country"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
}

type State struct {
	Cities []string
}

type Country struct {
	States map[string]*State
}
