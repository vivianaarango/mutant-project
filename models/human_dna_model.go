package models

// HumanDNA model for save dna human info.
type HumanDNA struct {
	PK       string `json:"pk"`        // dna base separated for "-", e.x: ATGCGA-CAGTGC-TTATGT-AGAAGG-CCCCTA-TCACTG.
	IsMutant bool   `json:"is_mutant"` // validate if the human is mutand.
}
