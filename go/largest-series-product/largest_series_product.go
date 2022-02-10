package lsproduct

import (
	"errors"
)

//LargestSeriesProduct given a string of digits, calculates the largest product
//for a contiguous substring of digits of length n.
func LargestSeriesProduct(s string, n int) (largest int, err error) {
	if n == 0 {
		return 1, nil // empty product is always 1
	}
	if n < 0 || n > len(s) {
		return -1, errors.New("n must be a positive number smaller than the string length")
	}

	for i := 0; i <= len(s)-n; i++ {
		substring := s[i : i+n]
		value := 1
		for _, char := range substring {
			if char < '0' || char > '9' {
				return -1, errors.New("string cannot contain a non-digit character")
			}
			value *= int(char - '0')
		}
		if value > largest {
			largest = value
		}
	}

	return
}
