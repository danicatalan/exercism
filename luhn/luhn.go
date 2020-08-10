package luhn

import (
	"strconv"
	"strings"
)

//Valid takes a string and validates it against the Luhn algorithm.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	d := []rune(s)
	for i := len(d) - 2; i >= 0; i -= 2 {
		c := d[i] - '0'
		if c < 5 {
			c = c * 2
		} else {
			c = c*2 - 9
		}
		d[i] = c + '0'
	}

	sum := 0
	for _, n := range d {
		sum += int(n - '0')
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
