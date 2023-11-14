package dtos

type Country struct {
	States map[string]*State
}

type State struct {
	Cities map[string]bool
}
