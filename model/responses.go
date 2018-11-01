package model

// AssignResponse - Contains response for assign distributor request
type AssignResponse struct {
	Status string `json:"status,omitempty"`
}

//CheckDistributionResponse - Contains response for checking distributor permission for the region
type CheckDistributionResponse struct {
	IsAuthorized string `json:"is_authorized,omitempty"`
}
