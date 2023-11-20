package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/RealImage/challenge2016/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

func CreateDistributorHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		fmt.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, "invalid param value")
		return
	}

	// reading data from body
	var distributor Distributor
	err := json.NewDecoder(r.Body).Decode(&distributor)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	distributor.Id = uuid.New().String()
	distributor.Name = name

	// saving the data to text file as we do not have to use DB
	err = saveDistributorData(distributor)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponseJson(w, http.StatusOK, distributor)
	return
}

func CreateSubDistributorHandler(w http.ResponseWriter, r *http.Request) {
	parentId := chi.URLParam(r, "parent_id")
	name := chi.URLParam(r, "name")
	if name == "" || parentId == "" {
		fmt.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, "invalid param value")
		return
	}

	// reading data from body
	var subDistributor Distributor
	err := json.NewDecoder(r.Body).Decode(&subDistributor)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	distributor, err := FetchDistributorByUid(parentId)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}
	if !CheckPermissions(subDistributor.Permissions, distributor.Permissions) {
		utils.WriteResponseJson(w, http.StatusBadRequest, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
		return
	}

	subDistributor.Id = uuid.New().String()
	subDistributor.Name = name
	subDistributor.Parent = parentId
	subDistributor.Permissions.Exclude = append(subDistributor.Permissions.Exclude, distributor.Permissions.Exclude...)

	// saving the data to text file as we do not have to use DB
	err = saveDistributorData(subDistributor)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponseJson(w, http.StatusOK, subDistributor)
	return
}

func CheckDistributorPermissionHandler(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	if uid == "" {
		fmt.Println("invalid param value")
		utils.WriteResponseJson(w, http.StatusBadRequest, "invalid param value")
		return
	}

	var permissions Permissions
	err := json.NewDecoder(r.Body).Decode(&permissions)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	distributor, err := FetchDistributorByUid(uid)
	if err != nil {
		fmt.Println(err)
		utils.WriteResponseJson(w, http.StatusBadRequest, err)
		return
	}

	if distributor.Parent != "" {
		if !CheckParentDistributorPermissions(permissions, distributor.Parent) {
			utils.WriteResponseJson(w, http.StatusBadRequest, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
			return
		}
	}

	if CheckPermissions(permissions, distributor.Permissions) {
		utils.WriteResponseJson(w, http.StatusOK, "YES: DISTRIBUTOR have mentioned permission to distribute")
		return
	} else {
		utils.WriteResponseJson(w, http.StatusBadRequest, "NO: DISTRIBUTOR does not have mentioned permission to distribute")
		return
	}
}

func CheckParentDistributorPermissions(permissions Permissions, parentId string) bool {
	distributor, err := FetchDistributorByUid(parentId)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return CheckPermissions(permissions, distributor.Permissions)
}

func CheckPermissions(subSet Permissions, superSet Permissions) bool {
	if isSubset(subSet.Include, superSet.Exclude) {
		return false
	}
	if isSubset(subSet.Include, superSet.Include) {
		return true
	}
	return false
}

func isSubset(subset, superset []string) bool {
	i, j := 0, 0
	for i < len(subset) && j < len(superset) {
		if subset[i] == superset[j] || strings.HasSuffix(subset[i], superset[j]) {
			i++
		}
		j++
	}

	return i == len(subset)
}

func FetchDistributorByUid(uid string) (distributor Distributor, err error) {
	file, err := os.Open(DISTRIBUTOR_DB)
	if err != nil {
		fmt.Println(err)
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

func saveDistributorData(distributor Distributor) (err error) {
	file, err := os.OpenFile(DISTRIBUTOR_DB, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(distributor)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}
