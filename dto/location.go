package dto

type City struct {
	Name string `json:"name"`
}

type State struct {
	Name   string `json:"state"`
	Cities []City `json:"cities"`
}

type Country struct {
	Name   string  `json:"country"`
	States []State `json:"states"`
}
