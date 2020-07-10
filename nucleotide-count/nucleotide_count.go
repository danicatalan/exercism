package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

const (
	validNucleotides = "ACGT"
)

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
	h = make(Histogram)

	for _, n := range validNucleotides {
		h[n] = 0
	}

	for _, c := range d {
		if _, exists := h[c]; !exists {
			return h, errors.New("Found invalid nucleotide")
		}
		h[c]++
	}
	return h, nil
}
