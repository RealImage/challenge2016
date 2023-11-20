package models

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var validate = validator.New()

type Distributor struct {
	ID     uuid.UUID `json:"ID"`
	Access `json:"Access"`
}

type Access struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

// Create Distributor Request Type

type CreateDistributorRequest struct {
	Access `json:"Access" validate:"required"`
}

func (r *CreateDistributorRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func (r *CreateDistributorRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

type CreateDistributorAPIResponse struct {
	Message *Distributor
}

func (cr *CreateDistributorAPIResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}

func NewCreateDistributorAPIResponse(dis *Distributor) *CreateDistributorAPIResponse {
	return &CreateDistributorAPIResponse{
		Message: dis,
	}
}

func CreateRequestToDistributor(req CreateDistributorRequest) *Distributor {
	return &Distributor{
		ID:     uuid.New(),
		Access: req.Access,
	}
}

// Create Split Distributor Request Type

type CreateSplitDistributorRequest struct {
	ParentId uuid.UUID `json:"ParentId" validate:"required"`
	ID       uuid.UUID `json:"ID"`
	Access   `json:"Access"`
}

func (r *CreateSplitDistributorRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func (r *CreateSplitDistributorRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

func SplitRequestToDistributor(req CreateSplitDistributorRequest) *Distributor {
	dist := &Distributor{
		Access: req.Access,
	}

	if req.ID != uuid.Nil {
		dist.ID = req.ID
	} else {
		dist.ID = uuid.New()
	}

	return dist
}

// Utility Types

type apiError struct {
	status int
	body   string
}

func (e *apiError) Write(w http.ResponseWriter) error {
	// Implement serialization and writing logic for the User API response
	// Serialize the struct r and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.status)
	return json.NewEncoder(w).Encode(e.body)
}

func NewAPIError(s int, b string) *apiError {
	return &apiError{
		status: s,
		body:   b,
	}
}
