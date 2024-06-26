package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
// type DNA struct {
// CountsFunc func() (Histogram, error)
// strand string
// }

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	// var h Histogram
	h := make(Histogram)
	for _, n := range "ACGT" {
		h[n] = 0
	}

	for _, v := range d {
		if v != 'A' && v != 'C' && v != 'G' && v != 'T' {
			return h, fmt.Errorf("invalid nucleotide %c", v)
		}
		h[v]++
	}
	return h, nil
}
