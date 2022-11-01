package services

import (
	"encoding/json"
	"errors"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"mutant-project/helpers"
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

type RequestBody struct {
	DNA []string `json:"dna"`
}

// MutantHelperInterface ...
// Detect ...
type MutantHelperInterface interface {
	Detect(dna []string) bool
	ValidateDNA(dnaRow string) bool
}

// DetectMutantsHandler ...
type DetectMutantsHandler struct {
	mutantHelperInterface MutantHelperInterface
}

// detectMutantsService init route for each service.
func (s *Service) detectMutantsService() {
	s.Router.HandleFunc("/mutant", func(w http.ResponseWriter, r *http.Request) {
		control := DetectMutantsHandler{
			mutantHelperInterface: &helpers.MutantHelper{},
		}
		request, err := control.getRequestBody(*r)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		response := control.mutantHelperInterface.Detect(request.DNA)
		if !response {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}).Methods(http.MethodPost)
}

func (h *DetectMutantsHandler) getRequestBody(r http.Request) (RequestBody, error) {
	var request RequestBody

	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)

	err = json.Unmarshal([]byte(buf.String()), &request)
	if err != nil {
		return RequestBody{}, err
	}

	dataToValidate := gojsonschema.NewGoLoader(request)
	bodyResultValidate, err := gojsonschema.Validate(validationSchema, dataToValidate)
	if err != nil {
		return RequestBody{}, err
	}

	if !bodyResultValidate.Valid() {
		return RequestBody{}, errors.New("")
	}

	for _, item := range request.DNA {
		if !h.mutantHelperInterface.ValidateDNA(item) {
			return RequestBody{}, errors.New("adn paili")
		}
	}

	return request, nil
}
