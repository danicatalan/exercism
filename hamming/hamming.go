package hamming

import (
	"errors"
)

// Distance takes two DNA strands and returns the Hamming distance between them.
// The strands should be the same length, or the function will throw an error.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("Can't compare strands with different length")
	}

	for i, l := range a {
		if byte(l) != b[i] {
			distance++
		}
	}

	return
}
