package models

// Structure to store the distributor and its permissions
type Distributor struct {
	ID          string
	Permissions *DistributorPermissions
	TrieRoot    *TrieNode
}

// Structure to represent the distributor's permissions
type DistributorPermissions struct {
	Include []string `json:"include" validate:"required"`
	Exclude []string `json:"exclude" binding:"required"`
}
