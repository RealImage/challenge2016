package service

import (
	"bufio"
	"encoding/json"
	"github.com/RealImage/challenge2016/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strings"
)

const DISTRIBUTOR_DB = "distributor.txt"

type Distributor struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Permissions Permissions `json:"permissions"`
	Parent      string      `json:"parent,omitempty"`
}

type Permissions struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude,omitempty"`
}

// CreateDistributorHandler - handler for creating new distributor
func CreateDistributorHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		log.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "invalid param value")
		return
	}

	// reading data from body
	var distributor Distributor
	err := json.NewDecoder(r.Body).Decode(&distributor)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	distributor.Id = uuid.New().String()
	distributor.Name = name

	// saving the data to text file as we do not have to use DB
	err = saveDistributorData(distributor)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.WriteResponseJson(w, http.StatusOK, distributor, "")
	return
}

// CreateSubDistributorHandler - handler for creating new sub distributor
// we check here if the permissions for new sub distributor matches the distributor
// only then the sub distributor will be permitted to distribute
func CreateSubDistributorHandler(w http.ResponseWriter, r *http.Request) {
	parentId := chi.URLParam(r, "parent_id")
	name := chi.URLParam(r, "name")
	if name == "" || parentId == "" {
		log.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "invalid param value")
		return
	}

	// reading data from body
	var subDistributor Distributor
	err := json.NewDecoder(r.Body).Decode(&subDistributor)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	distributor, err := FetchDistributorByUid(parentId)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	if !CheckPermissions(subDistributor.Permissions, distributor.Permissions) {
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
		return
	}

	subDistributor.Id = uuid.New().String()
	subDistributor.Name = name
	subDistributor.Parent = parentId
	subDistributor.Permissions.Exclude = append(subDistributor.Permissions.Exclude, distributor.Permissions.Exclude...)

	// saving the data to text file as we do not have to use DB
	err = saveDistributorData(subDistributor)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.WriteResponseJson(w, http.StatusOK, subDistributor, "")
	return
}

// CheckDistributorPermissionHandler - handler for checking the distributor has given permission to distribute
// if the distributor is sub distributor then also check the permission with parent distributor
func CheckDistributorPermissionHandler(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	permission := chi.URLParam(r, "permission")
	if uid == "" || permission == "" {
		log.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "invalid param value")
		return
	}

	permissions := Permissions{
		Include: []string{permission},
	}

	distributor, err := FetchDistributorByUid(uid)
	if err != nil {
		log.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	if distributor.Id == "" {
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "DISTRIBUTOR: id does not match with any distributor")
		return
	}

	if distributor.Parent != "" {
		if !CheckParentDistributorPermissions(permissions, distributor.Parent) {
			utils.WriteResponseJson(w, http.StatusBadRequest, nil, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
			return
		}
	}

	if CheckPermissions(permissions, distributor.Permissions) {
		utils.WriteResponseJson(w, http.StatusOK, nil, "YES: DISTRIBUTOR have mentioned permission to distribute")
		return
	} else {
		utils.WriteResponseJson(w, http.StatusBadRequest, nil, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
		return
	}
}

// CheckParentDistributorPermissions - fetching the distributor by uid and checking permission to distribute
func CheckParentDistributorPermissions(permissions Permissions, parentId string) bool {
	distributor, err := FetchDistributorByUid(parentId)
	if err != nil {
		log.Println(err)
		return false
	}

	return CheckPermissions(permissions, distributor.Permissions)
}

// CheckPermissions - check if the permissions are correct and distributor is eligible to distribute
func CheckPermissions(subSet Permissions, superSet Permissions) bool {
	if IsSubset(subSet.Include, superSet.Exclude) {
		return false
	}
	if IsSubset(subSet.Include, superSet.Include) {
		return true
	}
	return false
}

// IsSubset - func to check for subset permission with superset permission for distributor and sub-distributor
func IsSubset(subset, superset []string) bool {
	i, j := 0, 0
	for i < len(subset) && j < len(superset) {
		if subset[i] == superset[j] || strings.HasSuffix(subset[i], superset[j]) {
			i++
		}
		j++
	}

	return i == len(subset)
}

// FetchDistributorByUid - to fetch the distributor by uid
func FetchDistributorByUid(uid string) (distributor Distributor, err error) {
	file, err := os.Open(DISTRIBUTOR_DB)
	if err != nil {
		log.Println(err)
		return distributor, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, uid) {
			err = json.Unmarshal([]byte(line), &distributor)
			if err != nil {
				return distributor, err
			}
			break
		}
	}

	return distributor, err
}

// saveDistributorData - save the distributor data to file for further access
func saveDistributorData(distributor Distributor) (err error) {
	file, err := os.OpenFile(DISTRIBUTOR_DB, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(distributor)
	if err != nil {
		log.Println("Error writing file:", err)
		return err
	}

	return nil
}
