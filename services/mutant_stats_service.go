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

// ResponseBody data for body response.
type ResponseBody struct {
	CountMutantDNA int     `json:"count_mutant_dna"`
	CountHumanDNA  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

// mutantStatsService service for obtain stats of human dna.
func (s *Service) mutantStatsService() {
	s.Router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		// init arguments for service.
		control := initializeMutantStatsService()

		// get total humans and total humans with dna mutant.
		totalMutants, totalHumans, err := control.mutantRepositoryInterface.GetDNAStats()
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		// generate response body with the result.
		response := ResponseBody{
			CountMutantDNA: totalMutants,
			CountHumanDNA:  totalHumans,
			Ratio:          float64(totalMutants / totalHumans),
		}

		// convert response to json data.
		body, err := json.Marshal(response)
		if err != nil {
			fmt.Fprintf(w, "Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(body)
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}

// initializeMutantStatisticsService arguments for detect mutant service.
func initializeMutantStatsService() *MutantStatsHandler {
	return &MutantStatsHandler{
		mutantRepositoryInterface: repositories.NewMutantRepository(),
	}
}
