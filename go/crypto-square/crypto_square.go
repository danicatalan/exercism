package cryptosquare

import (
	"math"
	"strings"
)

//Encode takes an English text and returns the cryptosquare version of the text.
func Encode(text string) string {
	var nb strings.Builder
	for _, c := range strings.ToLower(text) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			nb.WriteRune(c)
		}
	}

	n := nb.String()
	sqrt := math.Sqrt(float64(len(n)))
	c, r := int(math.Ceil(sqrt)), int(math.Round(sqrt))

	var encoded strings.Builder
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			p := i + j*c
			if p < len(n) {
				encoded.WriteByte(n[p])
			} else {
				encoded.WriteString(" ")
			}
			if i != c-1 && j == r-1 {
				encoded.WriteString(" ")
			}
		}
	}

	return encoded.String()
}
