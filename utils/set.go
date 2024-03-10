package utils

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) AddItems(items []string) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set) Union(set2 Set) {
	for item := range set2 {
		s.Add(item)
	}
}

func (s Set) Intersection(set2 Set) Set {
	result := NewSet()
	for item := range s {
		if set2.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Contains(item string) bool {
	_, ok := s[item]
	return ok
}
