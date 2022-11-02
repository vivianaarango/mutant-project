package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"mutant-project/helpers"
	"mutant-project/repositories"
	"net/http"
	"strings"
)

// validationSchema struct to validate body request.
var validationSchema = gojsonschema.NewStringLoader(`{
		"type": "object",
			"properties": {
				"dna": {
					"type": "array",
					"items": {
						"type": "string"
					},
					"minItems": 1
				}
			},
			"required": [
				"dna"
			]
		}`)

// RequestBody data necessary from request.
type RequestBody struct {
	DNA []string `json:"dna"`
}

// MutantHelperInterface interface for helpers.MutantHelper
// Detect this method detect if the human dna given is of mutant or not.
// ValidateDNA this method check if the human dna given is valid.
type MutantHelperInterface interface {
	Detect(dna []string) bool
	ValidateDNA(dnaRow string) bool
}

// DetectMutantsHandler arguments necessary for detect mutant service.
type DetectMutantsHandler struct {
	mutantHelperInterface     MutantHelperInterface
	mutantRepositoryInterface repositories.MutantRepositoryInterface
}

// detectMutantsService service for detect if the human dna given is of mutant or not.
func (s *Service) detectMutantsService() {
	s.Router.HandleFunc("/mutant", func(w http.ResponseWriter, r *http.Request) {
		// init arguments for service.
		control := initializeDetectMutantService()

		// obtain and validate request body.
		request, err := control.getRequestBody(*r)
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		// check if the human dna is mutant or not.
		isMutant := control.mutantHelperInterface.Detect(request.DNA)

		// save dna info.
		err = control.mutantRepositoryInterface.Save(request.DNA, isMutant)
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		// response if the human is mutant or not.
		if !isMutant {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}).Methods(http.MethodPost)
}

// getRequestBody obtain and validate request body.
func (h *DetectMutantsHandler) getRequestBody(r http.Request) (RequestBody, error) {
	var request RequestBody

	// obtain data from request.
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)

	err = json.Unmarshal([]byte(buf.String()), &request)
	if err != nil {
		return RequestBody{}, err
	}

	// validate body data.
	dataToValidate := gojsonschema.NewGoLoader(request)
	bodyResultValidate, err := gojsonschema.Validate(validationSchema, dataToValidate)
	if err != nil {
		return RequestBody{}, err
	}

	if !bodyResultValidate.Valid() {
		return RequestBody{}, errors.New("")
	}

	for _, item := range request.DNA {
		// check if the dna obtain from request is valid.
		if !h.mutantHelperInterface.ValidateDNA(item) {
			return RequestBody{}, errors.New("adn paili")
		}
	}

	return request, nil
}

// initialize arguments for detect mutant service.
func initializeDetectMutantService() *DetectMutantsHandler {
	return &DetectMutantsHandler{
		mutantHelperInterface:     &helpers.MutantHelper{},
		mutantRepositoryInterface: repositories.NewMutantRepository(),
	}
}
