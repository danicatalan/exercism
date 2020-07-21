package isogram

import "unicode"

//IsIsogram takes a word or phrase and checks if all is letters are unique.
//Spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
	letters := make(map[rune]int)
	for _, l := range word {
		if !unicode.IsLetter(l) {
			continue
		}
		if _, exists := letters[unicode.ToLower(l)]; exists {
			return false
		}
		letters[unicode.ToLower(l)]++
	}
	return true
}
