package permissions

import (
	"errors"
	"log"
	"strings"
	"sync"
)

var ErrDistributorAlreadyExist = errors.New("distributor already exist")
var ErrDistributorNotFound = errors.New("distributor not found")

type (
	Service struct {
		mDistributors *sync.Mutex
		distributors  []*Distributor

		KnownRegions []*RegionInfo
		perms        *permissionTable
	}
)

func NewService() *Service {
	return &Service{
		mDistributors: &sync.Mutex{},
		distributors:  make([]*Distributor, 0),
		//KnownRegions:  regions,
		perms: &permissionTable{
			m:     &sync.Mutex{},
			table: map[string]Permissions{},
		},
	}
}

func (s *Service) Distributors() []*Distributor {
	s.mDistributors.Lock()
	cpDistributor := make([]*Distributor, len(s.distributors))
	copy(cpDistributor, s.distributors)
	s.mDistributors.Unlock()

	return cpDistributor
}

// CreatePermissions creates permissions for `newDistributor`. Pass "" as parentDistributor if there is no parent permissions
func (s *Service) CreatePermissions(parentDistributor, newDistributor string, included, excluded []string) (*Permissions, error) {
	var parentPermissions *Permissions
	var parent *Distributor
	var err error
	if parentDistributor != "" {
		parent, err = s.getDistributor(parentDistributor)
		if err != nil {
			return nil, errors.New("error getting parent distributor: " + err.Error())
		}

		parentPermissions, err = s.perms.Get(parentDistributor)
		if err != nil {
			return nil, errors.New("error getting parent permissions: " + err.Error())
		}
	}

	perms, err := NewPermissions(parentPermissions, included, excluded)
	if err != nil {
		return nil, err
	}

	d, _ := s.getDistributor(newDistributor)
	if d != nil {
		return nil, err
	}

	d = &Distributor{
		Name: newDistributor,
		m:    &sync.Mutex{},
	}

	err = s.perms.Add(d, perms)
	if err != nil {
		return nil, err
	}

	s.addDistributor(d)
	if parent != nil {
		parent.addSubDistributor(d)
	}
	return perms, nil
}

func (s *Service) UpdatePermissions(distributor string, newPermissions *Permissions) error {
	d, err := s.getDistributor(distributor)
	if err != nil {
		return err
	}

	for _, child := range d.children {
		childPerms, err := s.perms.Get(child.Name)
		if err != nil {
			log.Printf("cannot get child permissions: %s", err)
			continue
		}

		newPerms, err := NewPermissions(newPermissions, []string{}, []string{})
		if err != nil {
			log.Printf("cannot get child permissions: %s", err)
			continue
		}

		// include and exclude separately to avoid errors
		includes := strings.Split(childPerms.includes.String(), "\n")
		for _, include := range includes {
			newPerms.Include(include)
		}

		excludes := strings.Split(childPerms.excludes.String(), "\n")
		for _, exclude := range excludes {
			newPerms.Exclude(exclude)
		}

		err = s.UpdatePermissions(child.Name, newPerms)
		if err != nil {
			log.Printf("cannot update child permissions: %s", err)
		}
	}

	return s.perms.Update(distributor, newPermissions)
}

func (s *Service) GetPermissions(distributor string) (*Permissions, error) {
	return s.perms.Get(distributor)
}

func (s *Service) HasPermissions(distributor string, region string) (bool, error) {
	perms, err := s.perms.Get(distributor)
	if err != nil {
		return false, err
	}

	return perms.IsAllowed(region), nil
}

func (s *Service) RemovePermissions(distributor string) error {
	dist, err := s.getDistributor(distributor)
	if err != nil {
		return err
	}

	// delete children perms as well
	for _, child := range dist.children {
		err := s.RemovePermissions(child.Name)
		if err != nil {
			log.Printf("cannot delete child permissions: %s", err)
		}
	}

	s.perms.Delete(distributor)
	return s.removeDistributor(distributor)
}

func (s *Service) addDistributor(distributor *Distributor) {
	s.mDistributors.Lock()
	s.distributors = append(s.distributors, distributor)
	s.mDistributors.Unlock()
}

func (s *Service) getDistributor(distributor string) (*Distributor, error) {
	s.mDistributors.Lock()
	defer s.mDistributors.Unlock()

	for _, d := range s.distributors {
		if distributor == d.Name {
			return d, nil
		}
	}

	return nil, ErrDistributorNotFound
}

func (s *Service) removeDistributor(distributor string) error {
	s.mDistributors.Lock()
	defer s.mDistributors.Unlock()
	for i, d := range s.distributors {
		if distributor == d.Name {
			s.distributors = append(s.distributors[:i], s.distributors[i+1:]...)
			return nil
		}
	}

	return ErrDistributorNotFound
}
