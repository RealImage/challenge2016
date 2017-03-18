package distribution

type PermissionMatrix map[string]map[string]map[string]bool

// load the base permissions from the file
func (permissions *PermissionMatrix) Initialize(file string) ApplicationError {
	// load from CSV
}

// check whether the location is permitted
func (permissions *PermissionMatrix) IsAllowed(location string) bool {
	
}

// include the location in the permissions
func (permissions *PermissionMatrix) Include(location string) ApplicationError {
	
}

// exclude the location in the permissions
func (permissions *PermissionMatrix) Exclude(location string) ApplicationError {
	
}
