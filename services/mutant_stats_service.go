package services

import (
	"encoding/json"
	"fmt"
	"mutant-project/repositories"
	"net/http"
)

// MutantStatsHandler ... arguments necessary for detect mutant service.
type MutantStatsHandler struct {
	mutantRepositoryInterface repositories.MutantRepositoryInterface
}

type ResponseBody struct {
	CountMutantDNA int     `json:"count_mutant_dna"`
	CountHumanDNA  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

// mutantStatsService ... service for detect if the human dna given is of mutant or not.
func (s *Service) mutantStatsService() {
	s.Router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		// init arguments for service.
		control := initializeMutantStatsService()
		totalMutants, totalHumans, err := control.mutantRepositoryInterface.GetDNAStats()
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		response := ResponseBody{
			CountMutantDNA: totalMutants,
			CountHumanDNA:  totalHumans,
			Ratio:          0,
		}

		// We apply json.Marshal to send message to SQS.
		body, err := json.Marshal(response)
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(body)
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)
}

// initializeMutantStatisticsService arguments for detect mutant service.
func initializeMutantStatsService() *MutantStatsHandler {
	return &MutantStatsHandler{
		mutantRepositoryInterface: repositories.NewMutantRepository(),
	}
}
