package secret

import (
	"strconv"
)

var (
	dictionary = []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
		"REVERSE",
	}
)

//Handshake takes a decimal number and converts it to the appropriate sequence
//of events for a secret handshake
func Handshake(n uint) []string {
	var result []string
	b := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(b); i++ {
		start, end := len(b)-i-1, len(b)-i
		if b[start:end] == "1" {
			if dictionary[i] == "REVERSE" {
				result = reverse(result)
			} else {
				result = append(result, dictionary[i])
			}
		}
	}
	return result
}

func reverse(a []string) []string {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return a
}
