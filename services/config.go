package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Service struct with the arguments necessary to the api.
type Service struct {
	Router *mux.Router
}

// Initialize initializes the api and the routes.
func (s *Service) Initialize() {
	// create a new router.
	s.Router = mux.NewRouter()

	// set the defined routes.
	//s.setRoutes()

	s.detectMutantsService()
}

// setRoutes init route for each service.
/*func (s *Service) setRoutes() {
	s.Router.HandleFunc("/mutant", func(w http.ResponseWriter, r *http.Request) {
		service := Handler{
			mutantHelperInterface: &helpers.MutantAlgorithmHelper{},
		}
		err := service.detectMutantsHandler([]string{})

		if !err {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			respondWithSuccess(RequestBody{}, w)
		}
	}).Methods(http.MethodPost)
}*/

// Run the app on it's router
func (s *Service) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
