package services

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Service struct with the arguments necessary to init api.
type Service struct {
	Router *mux.Router
}

// Initialize initialize services.
func (s *Service) Initialize() {
	// create a new router.
	s.Router = mux.NewRouter()

	// initialize detect mutant service.
	s.detectMutantsService()
}

// Run the app on it's router
func (s *Service) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
