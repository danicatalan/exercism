/*
Package pythagorean can be solved by force brute with three nested loops on a, b, c,
but time complexity can be reduced using the Euclid's formula (TO-DO)
*/
package pythagorean

import "math"

// Triplet is a set of natural numbers a, b and c that fits the rule: a^2 + b^2 == c^2
type Triplet [3]int

// Range takes a minimun and maximun number and finds all the triplets inside the range.
func Range(min, max int) (triplets []Triplet) {
	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum takes a value (perimeter of a triangle) and finds all the triplets that form a triangle with that perimeter
func Sum(p int) (triplets []Triplet) {
	for _, t := range Range(1, p) {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return
}
