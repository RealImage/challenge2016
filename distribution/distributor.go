package distribution

type Distributor struct {
	name        string
	permissions PermissionMatrix
	parent      *Distributor
}

// initialize the distributor
func (distributor *Distributor) Initialize(name string, parent *Distributor) ApplicationError {
	distributor.name = name
	distributor.parent = parent
	distributor.permissions = basePermissions
	return nil
}

// check whether the location is in the scope of the distributor
func (distributor *Distributor) HasScope(location string) bool {
	// if it is NOT a sub-distributor, then it has scope to all locations
	if distributor.parent == nil {
		return true
	} else { // otherwise, the scope is limited to scope of the parent distributor
		return distributor.parent.permissions.IsAllowed(location)
	}
}

// include the location to the distributor permissions
func (distributor *Distributor) Include(location string) ApplicationError {
	// if the distributor has location in its scope, include the location
	if distributor.HasScope(location) {
		return distributor.permissions.Include(location)
	} else { // otherwise, raise error
		return DistributionScopeError(location)
	}
}

// exclude the location to the distributor permissions
func (distributor *Distributor) Exclude(location string) ApplicationError {
	// no need to check the scope for exclude
	return distributor.permissions.Exclude(location)
}

// query if the distributor can distribute in a location
func (distributor *Distributor) CanDistribute(location string) bool {
	// query the location in permission matrix
	return distributor.permissions.IsAllowed(location)
}

// TODO(ilayaraja): Do not expose
func (distributor *Distributor) Permissions() PermissionMatrix {
	return distributor.permissions
}
