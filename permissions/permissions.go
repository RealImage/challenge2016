package permissions

import (
	"errors"
	"strings"
	"sync"
)

var ErrCannotExclude = errors.New("cannot exclude")
var ErrCannotInclude = errors.New("cannot include")
var ErrPermissionsAlreadyExist = errors.New("permissions already exist")
var ErrPermissionsNotFound = errors.New("permissions not found")

type (
	Distributor struct {
		Name string

		m        *sync.Mutex
		children []*Distributor
	}

	Permissions struct {
		parentPerms *Permissions

		m                  *sync.Mutex
		includes, excludes *regions
	}

	permissionTable struct {
		m *sync.Mutex
		// the copies of perms to avoid changes
		table map[string]Permissions
	}
)

func NewPermissions(parent *Permissions, included, excluded []string) (*Permissions, error) {
	perms := &Permissions{
		parentPerms: parent,
		m:           &sync.Mutex{},
		includes:    newRegions(),
		excludes:    newRegions(),
	}

	for _, incl := range included {
		if incl == "" {
			continue
		}

		err := perms.Include(incl)
		if err != nil {
			return nil, err
		}
	}

	for _, excl := range excluded {
		if excl == "" {
			continue
		}

		err := perms.Exclude(excl)
		if err != nil {
			return nil, err
		}
	}

	return perms, nil
}

// IsAllowed checks if region is allowed
func (p *Permissions) IsAllowed(regionDelta string) bool {
	p.m.Lock()
	defer p.m.Unlock()

	return p.isAllowedByParent(regionDelta) &&
		p.includes.Contains(regionDelta, false) &&
		!p.excludes.Contains(regionDelta, true)

}

// Include takes regionDelta. If the region exists, it will be deleted, if it does not exist, it will be created.
func (p *Permissions) Include(region string) error {
	p.m.Lock()
	defer p.m.Unlock()

	if region == "" {
		return nil
	}

	if !p.isAllowedByParent(region) || p.excludes.Contains(region, true) {
		return ErrCannotInclude
	}

	if p.includes.Contains(region, true) {
		p.includes.Remove(region)
		return nil
	}

	p.includes.Add(region)
	return nil
}

// Exclude takes regionDelta. If the region exists, it will be deleted, if it does not exist, it will be created.
func (p *Permissions) Exclude(regionDelta string) error {
	p.m.Lock()
	defer p.m.Unlock()

	// no sense exclude the region because it's already not allowed by parent
	if regionDelta == "" || !p.isAllowedByParent(regionDelta) {
		return nil
	}

	if p.includes.Contains(regionDelta, true) {
		return ErrCannotExclude
	}

	if p.excludes.Contains(regionDelta, true) {
		p.excludes.Remove(regionDelta)
		return nil
	}

	p.excludes.Add(regionDelta)
	return nil
}

func (p *Permissions) String() string {
	str := ""
	if p.includes != nil {
		str += stringPrettifier(p.includes.String(), "INCLUDE: ")
	}

	if p.excludes != nil {
		str += stringPrettifier(p.excludes.String(), "EXCLUDE: ")
	}

	return str
}

func (p *Permissions) isAllowedByParent(region string) bool {
	if p.parentPerms == nil {
		return true
	}

	return p.parentPerms.IsAllowed(region)
}

// Add adds perms by distributor.Name
func (t permissionTable) Add(distributor *Distributor, permissions *Permissions) error {
	t.m.Lock()
	defer t.m.Unlock()

	_, ok := t.table[distributor.Name]
	if ok {
		return ErrPermissionsAlreadyExist
	}

	t.table[distributor.Name] = *permissions
	return nil
}

// Get copies perms and returns the reference to copy
func (t permissionTable) Get(distributorName string) (*Permissions, error) {
	t.m.Lock()
	defer t.m.Unlock()

	perm, ok := t.table[distributorName]
	if !ok {
		return nil, ErrPermissionsNotFound
	}

	return &perm, nil
}

// Update updates perms
func (t permissionTable) Update(distributorName string, newPerms *Permissions) error {
	t.m.Lock()
	defer t.m.Unlock()

	_, ok := t.table[distributorName]
	if !ok {
		return ErrPermissionsNotFound
	}

	t.table[distributorName] = *newPerms
	return nil
}

// Delete deletes perms by distributor
func (t permissionTable) Delete(distributor string) {
	t.m.Lock()
	delete(t.table, distributor)
	t.m.Unlock()

	return
}

func (d *Distributor) addSubDistributor(distributor *Distributor) {
	d.m.Lock()
	d.children = append(d.children, distributor)
	d.m.Unlock()

	return
}

func stringPrettifier(str, prefix string) string {
	if str == "" {
		return ""
	}

	res := ""
	excl := strings.Split(str, "\n")
	for _, s := range excl {
		res += prefix + s + "\n"
	}

	return res
}
