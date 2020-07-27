package wordcount

import (
	"strings"
	"unicode"
)

// Frequency abstracts how often a word appears in a phrase.
type Frequency map[string]int

//WordCount returns the frequency of words for a given phrase.
func WordCount(phrase string) Frequency {
	output := make(Frequency)

	split := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '\''
	}
	words := strings.FieldsFunc(phrase, split)
	for _, w := range words {
		w = strings.Trim(strings.ToLower(w), "'")
		output[w]++
	}

	return output
}
