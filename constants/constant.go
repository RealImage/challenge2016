package constants

const (
	DEFAULT_PORT = "8080"
	ContextPath  = "/api/v1"

	CreateDistributorURI             = ContextPath + "/distributors"
	GetDistributorLocationDetailsURI = ContextPath + "/distributors/{distributor}/location"
	GetDistributorDetailsURI         = ContextPath + "/distributors/{distributor}"

	GET  = "GET"
	POST = "POST"

	DistributorCreated = "Distributor has been created"
	DistributorFailed  = "Distributor creation has been failed"
)
