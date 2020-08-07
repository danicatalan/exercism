package isbn

import (
	"strings"
)

//IsValidISBN given a string returns if the provided string is a valid ISBN-10.
func IsValidISBN(s string) bool {
	n := strings.ReplaceAll(s, "-", "")

	if len(n) != 10 {
		return false
	}

	sum := 0
	for i, d := range n {
		var di int

		if d == 'X' && i == len(n)-1 {
			di = 10
		} else {
			di = int(d - '0')
		}

		r := (len(n) - i) * di
		sum += r
	}

	if sum%11 == 0 {
		return true
	}

	return false
}
