package grains

import (
	"errors"
)

// Square takes the position of a square in a chessboard and returns the number of grains of wheat in that square,
// given that the number on each square doubles. Returns an error if the position provided is out of range.
func Square(p int) (uint64, error) {
	if p < 1 || p > 64 {
		return 0, errors.New("position is out of range")
	}
	t := 1
	for i := 1; i < p; i++ {
		t *= 2
	}
	return uint64(t), nil
}

// Total returns the amount of grains on the chessboard.
func Total() uint64 {
	var t uint64
	for i := 1; i <= 64; i++ {
		r, _ := Square(i)
		t += r
	}
	return t
}
