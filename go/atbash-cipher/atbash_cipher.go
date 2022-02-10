package atbash

import (
	"strings"
	"unicode"
)

//Atbash takes a text and returns it encrypted using the atbash cipher.
func Atbash(text string) string {
	var o strings.Builder
	for _, c := range strings.ToLower(text) {
		if unicode.IsLetter(c) {
			o.WriteRune('z' + 'a' - c)
		}
		if unicode.IsNumber(c) {
			o.WriteRune(c)
		}
		if len(o.String())%6 == 5 {
			o.WriteRune(' ')
		}
	}

	return strings.TrimSpace(o.String())
}
