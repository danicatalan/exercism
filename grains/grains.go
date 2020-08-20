/*
Package grains provides utilities for solving the grains/chessboard problem.

	Uses geometric sequence rule `a1 * pow(r, n - 1)` to avoid recursive calculations.
	As the common ratio is 2, then `pow(2, n-1)` can be optimized using left bitshift operator `1 << (n - 1)`.
	For the total, this packages uses the sum of a series rule `(a1 * (1 - pow(r,n))) / (1 - r)`.
	a1 is ommited from all formulas as we always start from position 1.
	More details about geometric sequence here: http://www.mathguide.com/lessons/SequenceGeometric.html
*/
package grains

import (
	"errors"
)

// Square takes the position of a square in a chessboard and returns the number of grains of wheat in that square,
// given that the number on each square doubles. Returns an error if the position provided is out of range.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("position is out of range")
	}
	return 1 << (n - 1), nil
}

// Total returns the amount of grains on the chessboard.
func Total() uint64 {
	r, n := 2, 64
	return uint64((1 - (1 << n)) / (1 - r))
}
