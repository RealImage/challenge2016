package permissions

import (
	"strings"
	"sync"
)

const separator = "-"

func newRegions() *regions {
	return &regions{
		m:     &sync.Mutex{},
		table: map[string]*regions{},
	}
}

type regions struct {
	m     *sync.Mutex
	table map[string]*regions
}

func (r *regions) Add(fullRegion string) {
	if fullRegion == "" {
		return
	}

	splitRegions := strings.Split(fullRegion, separator)
	region := splitRegions[len(splitRegions)-1]
	subRegions, ok := r.table[region]
	if !ok {
		subRegions = newRegions()
		r.table[region] = subRegions
	}

	if len(splitRegions) == 1 {
		return
	}

	str := strings.Join(splitRegions[:len(splitRegions)-1], separator)
	subRegions.Add(str)
	return
}

func (r *regions) Contains(fullRegion string, strict bool) bool {
	splitRegions := strings.Split(fullRegion, separator)
	sub, ok := r.table[splitRegions[len(splitRegions)-1]]
	if !ok {
		return ok
	}

	if len(splitRegions) == 1 || len(sub.table) == 0 {
		if strict {
			return len(splitRegions) == 1 && len(sub.table) == 0
		}

		return true
	}

	subRegions := strings.Join(splitRegions[:len(splitRegions)-1], separator)
	return sub.Contains(subRegions, strict)

	//if !strict || !ok {
	//	return ok
	//}
	//
	//if len(splitRegions) == 1 && len(sub.table) == 0 {
	//	return true
	//}
	//
	//subRegions := strings.Join(splitRegions[:len(splitRegions)-1], separator)
	//return sub.Contains(subRegions, strict)
}

func (r *regions) Remove(fullRegion string) {
	if fullRegion == "" {
		return
	}

	r.m.Lock()
	defer r.m.Unlock()

	splitRegions := strings.Split(fullRegion, separator)
	regName := splitRegions[len(splitRegions)-1]
	if len(splitRegions) == 1 {
		delete(r.table, regName)
		return
	}

	reg, ok := r.table[regName]
	if !ok {
		return
	}

	subRegions := strings.Join(splitRegions[:len(splitRegions)-1], separator)
	reg.Remove(subRegions)

	//if there are no more subRegions then delete region
	if len(reg.table) == 0 {
		delete(r.table, regName)
	}

	return
}

func (r *regions) String() string {
	if len(r.table) == 0 {
		return ""
	}

	regions := make([]string, 0, len(r.table))
	for name, subRegions := range r.table {
		subs := subRegions.String()
		if subs == "" {
			regions = append(regions, name)
			continue
		}

		splitSubs := strings.Split(subs, "\n")
		for _, s := range splitSubs {
			regions = append(regions, s+separator+name)
		}
	}

	return strings.Join(regions, "\n")
}
