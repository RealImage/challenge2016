package restservice /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	dist "../distributions"
	"encoding/json"
	"flag"
	"net/http"
)

var distributors map[string]*dist.Distributor

var port string

func init() {
	portPtr := flag.String("port", "7770", "REST Service Port")
	port = ":" + *portPtr
	distributors = make(map[string]*dist.Distributor)
}

// DistributorData Struct stores the Incoming Distributor JSON Data
type DistributorData struct {
	ParentDistributorName string          `json:"parentDistributorName"`
	DistributorName       string          `json:"distributorName"`
	IncludeData           []dist.Location `json:"includeData"`
	ExcludeData           []dist.Location `json:"excludeData"`
}

type verifyDistributorData struct {
	DistributorName string          `json:"distributorName"`
	Locations       []dist.Location `json:"locations"`
}

// ResponseMessage Struct stores the outgoing Response Message
type ResponseMessage struct {
	ServerResponse []string
}

func duplicateDistributors(name string) bool {
	for distname := range distributors {
		if name == distname {
			return true
		}
	}
	return false
}

func addDistributor(w http.ResponseWriter, req *http.Request) {
	distData := []DistributorData{}
	msg := []string{}
	decodeErr := json.NewDecoder(req.Body).Decode(&distData)
	if decodeErr != nil {
		dist.ErrorLog("Failed to Decode JSON Input Stream", decodeErr)
	}
	for _, distributorselect := range distData {
		var parentDistributor *dist.Distributor
		if distributorselect.ParentDistributorName == "none" || distributorselect.ParentDistributorName == "" {
			parentDistributor = nil
		} else if _, parentcheck := distributors[distributorselect.ParentDistributorName]; parentcheck {
			parentDistributor = distributors[distributorselect.ParentDistributorName]
		} else {
			errormsg := "Cannot Create Distributor: " + distributorselect.DistributorName + " As No Parent Distributor: " + distributorselect.ParentDistributorName + "; "
			msg = append(msg, errormsg)
			dist.CustomErrorLog(errormsg)
			continue
		}
		duplicate := duplicateDistributors(distributorselect.DistributorName)
		if distributorselect.DistributorName != "" && !duplicate {
			distributor := &dist.Distributor{}
			distributor.Initialize(distributorselect.DistributorName, parentDistributor)
			distributors[distributorselect.DistributorName] = distributor
			dist.InfoLog("Created Distributor: " + distributorselect.DistributorName)
			includeerr := distributor.Include(distributorselect.IncludeData)
			if includeerr != nil {
				errormsg := "Cannot Include Permissions for Distributor: " + distributorselect.DistributorName + "; "
				dist.ErrorLog(errormsg, includeerr)
				msg = append(msg, errormsg)
				continue
			}
			excludeerr := distributor.Exclude(distributorselect.ExcludeData)
			if excludeerr != nil {
				errormsg := "Cannot Exclude Permissions for Distributor: " + distributorselect.DistributorName + "; "
				dist.ErrorLog(errormsg, excludeerr)
				msg = append(msg, errormsg)
				continue
			}
			successmsg := "Created Distributor: " + distributorselect.DistributorName + " And Sucessfully Updated Permissions; "
			dist.InfoLog(successmsg)
			msg = append(msg, successmsg)
		} else {
			dist.CustomErrorLog("Invalid or Duplicate Distributor Name")
			errormsg := "Cannot Create Distributor: " + distributorselect.DistributorName
			msg = append(msg, errormsg)
		}
	}
	resp := ResponseMessage{msg}
	dist.CustomLog(resp)
	response, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		dist.ErrorLog("Failed to Marshal to JSON", marshalErr)
	}
	w.Write(response)
}

func listDistributor(w http.ResponseWriter, req *http.Request) {
	distlist := []string{}
	for dist := range distributors {
		distlist = append(distlist, dist)
	}
	resp := ResponseMessage{distlist}
	dist.CustomLog(resp)
	response, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		dist.ErrorLog("Failed to Marshal to JSON", marshalErr)
	}
	w.Write(response)
}

func verifyDistributor(w http.ResponseWriter, req *http.Request) {
	var verifydist verifyDistributorData
	var msg []string
	decodeErr := json.NewDecoder(req.Body).Decode(&verifydist)
	if decodeErr != nil {
		dist.ErrorLog("Failed to Decode JSON Input Stream", decodeErr)
	}
	if verifydist.DistributorName != "" {
		dist := distributors[verifydist.DistributorName]
		for _, locationselect := range verifydist.Locations {
			if dist.VerifyPermissions(locationselect) {
				successmsg := "YES" + " For : " + locationselect.City + "-" + locationselect.Province + "-" + locationselect.Country + "; "
				msg = append(msg, successmsg)
			} else {
				errormsg := "NO" + " For : " + locationselect.City + "-" + locationselect.Province + "-" + locationselect.Country + "; "
				msg = append(msg, errormsg)
			}
		}
	} else {
		errmsg := "Distributor Name Missing; "
		msg = append(msg, errmsg)
		dist.CustomErrorLog(errmsg)
	}
	resp := ResponseMessage{msg}
	dist.CustomLog(resp)
	response, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		dist.ErrorLog("Failed to Marshal to JSON", marshalErr)
	}
	w.Write(response)
}

// StartService starts this REST Service
func StartService() {
	http.HandleFunc("/addDist/", addDistributor)
	http.HandleFunc("/listDist/", listDistributor)
	http.HandleFunc("/verifyDist/", verifyDistributor)
	dist.InfoLog("Started Service on port: " + port)
	http.ListenAndServe(port, nil)
}
