package distribution

type Distributor struct {
	name         string
	permissions  PermissionMatrix
	parent       *Distributor
}

// initialize the distributor
func (distributor *Distributor) Initialize(name string, parent *Distributor) ApplicationError {
	
}

// check whether the location is in the scope of the distributor
func (distributor *Distributor) HasScope(location string) bool {
	// if it is NOT a sub-distributor, return true
	// otherwise, check if the parent distributor has the location in its scope
}

// include the location to the distributor
func (distributor *Distributor) Include(location string) ApplicationError {
	// if the distributor has location in its scope
		// add the location 
	// otherwise, raise error
}

// exclude the location in the permissions
func (distributor *Distributor) Exclude(location string) ApplicationError {
	// if the distributor has location in its scope
		// remove the location 
	// otherwise, raise error	
}

// query if the distributor can distribute in a location
func (distributor *Distributor) CanDistribute(location string) bool {
	// query the permission matrix of the distributor
}
