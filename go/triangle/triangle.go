// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Kind abstracts different types of triangles.
type Kind int

const (
	// NaT Not a triangle.
	NaT = iota
	// Equ An equilateral triangle has all three sides the same length.
	Equ
	// Iso An isosceles triangle has at least two sides the same length.
	Iso
	// Sca A scalene triangle has all sides of different lengths.
	Sca
)

// KindFromSides returns the kind of triangle built by the parameters provided,
// or if the parameters don't build a triangle at all.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}

	for _, s := range sides {
		if s <= 0 || math.IsNaN(s) || math.IsInf(s, 0) {
			return NaT
		}
		if sides[0]+sides[1]+sides[2]-s < s {
			return NaT
		}
	}

	if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	}
	if sides[0] == sides[1] || sides[0] == sides[2] || sides[1] == sides[2] {
		return Iso
	}
	return Sca
}
