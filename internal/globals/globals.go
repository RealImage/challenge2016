package globals

import "challengeQube/dtos"

var MasterData map[string]*dtos.Country
var DistributorData map[string]dtos.Distributor

const (
	TypeErrorMsg = "ERROR_MESSAGE"
	TypeMsg      = "MESSAGE"
	TypeSuccess  = "SUCCESS"
	TypeFailed   = "FAILED"
	TypeInclude  = "include"
	TypeExclude  = "exclude"
)

var (
	CsvFileName = "cities.csv"
	AllowBool   = true
)
