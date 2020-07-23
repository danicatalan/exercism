package pangram

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

//IsPangram takes a sentence and returns true if it contains all letters from the alphabet.
func IsPangram(sentence string) bool {
	s := strings.ToLower(sentence)
	for _, l := range alphabet {
		if !strings.ContainsRune(s, l) {
			return false
		}
	}
	return true
}
