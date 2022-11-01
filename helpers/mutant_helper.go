package helpers

import (
	"strings"
)

// MutantHelper struct for mutant helper methods.
type MutantHelper struct{}

// Detect this method detect if the human dna given is of mutant or not.
// This only contains nitrogenous base of dna (A,T,C,G).
// A human is mutant if exist more than one sequence of four letters equal, obliquely,
// horizontally or vertically that are not repeated.
func (h *MutantHelper) Detect(adn []string) bool {
	// we create a matrix to storage each nitrogenous base of dna.
	mutantDNA := make([][]string, len(adn))

	// looping through the slice to declare
	// slice of slices of length nxn.
	for i := 0; i < len(adn); i++ {
		// separate chain of dna.
		separatedDNA := strings.Split(adn[i], "")
		mutantDNA[i] = make([]string, len(separatedDNA))

		// assigning values to each  slice of a slice
		for j := 0; j < len(separatedDNA); j++ {
			// storage the value in upper case to validate coincidences.
			mutantDNA[i][j] = strings.ToUpper(separatedDNA[j])
		}
	}

	if horizontal(mutantDNA) > 1 || vertical(mutantDNA) > 1 {
		return false
	}

	if horizontal(mutantDNA) == 1 || vertical(mutantDNA) == 1 {
		return true
	}

	return false
}

// ValidateDNA this method check if the human dna given is valid
// for evaluated if this is mutant.
func (h *MutantHelper) ValidateDNA(dnaRow string) bool {
	// separate dna for chars.
	separatedDNA := strings.Split(dnaRow, "")

	// validate if the dna contains only nitrogenous base.
	for _, dna := range separatedDNA {
		dna = strings.ToUpper(dna)
		if dna == "A" || dna == "T" || dna == "C" || dna == "G" {
			continue
		}
		return false
	}

	return true
}

// horizontal check the number of sequence of four letters equals horizontally
// that exist in the dna given.
func horizontal(mutantADN [][]string) int {
	// count for number of sequences found.
	countValid := 0
	for r := 0; r < len(mutantADN); r++ {
		// init coincidences number.
		coincidences := 1
		// value to compare values.
		currentValue := ""

		for c := 0; c < len(mutantADN[r])-1; c++ {
			// set value to compare.
			if currentValue == "" {
				currentValue = mutantADN[r][c]
			}

			// validate if the letters are equals.
			if currentValue == mutantADN[r][c+1] {
				coincidences++
			} else {
				currentValue = ""
			}

			// check if the sequence has four letters and adds to count.
			if coincidences == 4 {
				countValid++
				coincidences = 1
			}
		}
	}

	return countValid
}

// vertical check the number of sequence of four letters equals vertically
// that exist in the dna given.
func vertical(mutantADN [][]string) int {
	// count for number of sequences found.
	countValid := 0
	for i := 0; i < len(mutantADN); i++ {
		// init coincidences number.
		coincidences := 1
		// value to compare values.
		currentValue := ""

		for j := 0; j < len(mutantADN)-1; j++ {
			// set value to compare.
			if currentValue == "" {
				currentValue = mutantADN[j][i]
			}

			// validate if the letters are equals.
			if currentValue == mutantADN[j+1][i] {
				coincidences++
			} else {
				currentValue = ""
			}

			// check if the sequence has four letters and adds to count.
			if coincidences == 4 {
				countValid++
				coincidences = 1
			}
		}
	}

	return countValid
}
