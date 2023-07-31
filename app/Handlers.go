package app

import (
	"encoding/json"
	"golang/dto"
	Service "golang/service"
	"net/http"
)

// The Handlers struct is declared & used to store an instance of the service layer.
type Handlers struct {
	Service Service.Service
}

// In this function, the request is decoded, validated & sent to service layer and
// the response from service layer is validated and returned as HTTP response.
func (h *Handlers) DistributorPermissions(w http.ResponseWriter, req *http.Request) {
	errorsArr := make([]dto.ValidateResponse, 0)
	distributorsReq := make([]dto.DistributorRequest, 0)
	errorReq := json.NewDecoder(req.Body).Decode(&distributorsReq)
	if errorReq != nil {
		errorsArr = append(errorsArr, dto.ValidateResponse{Code: "400", Message: errorReq.Error()})
		writeResponse(w, http.StatusBadRequest, errorsArr)
		return
	}

	//Call to the service layer is made.
	response, err := h.Service.Distributors(distributorsReq)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	} else {
		writeResponse(w, http.StatusOK, response)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
