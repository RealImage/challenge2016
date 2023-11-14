package handlers

import (
	"challengeQube/internal/globals"
	"challengeQube/internal/services/authorisation"
	"context"
	"encoding/json"
	"net/http"

	"challengeQube/dtos"

	"github.com/FenixAra/go-prom/prom"
	"github.com/julienschmidt/httprouter"
)

// setting fresh empty context
var ctx = context.Background()

func setAuthorisationRoutes(router *httprouter.Router) {
	router.POST("/v1/authorize-distributor", prom.Track(Authorize, "Authorize"))
	router.GET("/v1/get-auth-status", prom.Track(GetAuthorisationStatus, "GetAuthorisationStatus"))
}

func Authorize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	rd := &RequestData{
		w: w,
		r: r,
	}
	a := authorisation.New(globals.AllowBool)
	decoder := json.NewDecoder(r.Body)
	req := &dtos.AuthorisationReq{}

	err := decoder.Decode(req)
	if err != nil {
		return writeJSONMessage(err.Error(), globals.TypeErrorMsg, http.StatusBadRequest, rd)
	}
	err = a.GiveAuthorisation(ctx, req)
	if err != nil {
		return writeJSONMessage(err.Error(), globals.TypeErrorMsg, http.StatusBadRequest, rd)
	}
	return writeJSONMessage(globals.TypeSuccess, globals.TypeMsg, http.StatusOK, rd)
}

func GetAuthorisationStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	rd := &RequestData{
		w: w,
		r: r,
	}
	a := authorisation.New(globals.AllowBool)
	//set get req
	region := dtos.Location{
		City:    r.FormValue("city"),
		State:   r.FormValue("state"),
		Country: r.FormValue("country"),
	}
	req := &dtos.GetStatusReq{
		DistributorName: r.FormValue("name"),
		Region:          region,
	}
	allow, err := a.GetAuthorisationStatus(ctx, req)
	if err != nil {
		return writeJSONMessage(err.Error(), globals.TypeErrorMsg, http.StatusBadRequest, rd)
	}
	if allow {
		return writeJSONMessage("YES", globals.TypeMsg, http.StatusOK, rd)
	} else {
		return writeJSONMessage("NO", globals.TypeMsg, http.StatusOK, rd)
	}
}
