package helpers

import (
	"fmt"
	"strings"
)

// MutantHelper struct for mutant helper methods.
type MutantHelper struct{}

// Detect this method detect if the human dna given is of mutant or not.
// This only contains nitrogenous base of dna (A,T,C,G).
// A human is mutant if exist more than one sequence of four letters equal, obliquely,
// horizontally or vertically that are not repeated.
func (mh *MutantHelper) Detect(dna []string) bool {
	// we create a matrix to storage each nitrogenous base of dna.
	mutantDNA := make([][]string, len(dna))

	// looping through the slice to declare
	// slice of slices of length nxn.
	for i := 0; i < len(dna); i++ {
		// separate chain of dna.
		separatedDNA := strings.Split(dna[i], "")
		mutantDNA[i] = make([]string, len(separatedDNA))

		// assigning values to each  slice of a slice
		for j := 0; j < len(separatedDNA); j++ {
			// storage the value in upper case to validate coincidences.
			mutantDNA[i][j] = strings.ToUpper(separatedDNA[j])
		}

		fmt.Println(mutantDNA[i])
	}

	v, h, tr, tl := searchMutant(mutantDNA)
	return isMutant(v, h, tr, tl)
}

// ValidateDNA this method check if the human dna given is valid
// for evaluated if this is mutant.
func (mh *MutantHelper) ValidateDNA(dnaRow string) bool {
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

func ones(size int) []int {
	onesSlice := make([]int, size)
	for i := 0; i < len(onesSlice); i++ {
		onesSlice[i] = 1
	}
	return onesSlice
}

func searchMutant(a [][]string) ([]int, []int, []int, []int) {
	v := ones(len(a))
	h := ones(len(a))
	tr := ones(len(a)*2 - 1)
	tl := ones(len(a)*2 - 1)

	for i := range a {
		for j := range a[i] {

			// VERTICAL
			if i > 0 {
				if a[i][j] == a[i-1][j] {
					v[j]++
				}
			}

			// TOP RIGHT
			if j < len(a)-1 && i > 0 {
				if a[i][j] == a[i-1][j+1] {
					tr[i+j]++
				}
			}

			// HORIZONTAL
			if j > 0 {
				if a[i][j] == a[i][j-1] {
					h[i]++
				}
			}

			// TOP LEFT
			if j > 0 && i > 0 {
				if a[i][j] == a[i-1][j-1] {
					_j := len(a) - j - 1
					tl[_j+i]++
				}
			}
		}
	}

	return v, h, tr, tl
}

func isMutant(v []int, h []int, tr []int, tl []int) bool {
	is := false
	all := append(v, h...)
	all = append(all, tr...)
	all = append(all, tl...)

	for _, l := range all {
		if l == 4 {
			is = true
		}
	}
	return is
}
