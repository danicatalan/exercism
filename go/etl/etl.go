package etl

import (
	"strings"
)

// Transform takes a list of scores and the strings assigned to that value
// and returns a list of strings with the score assigned to that string.
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return
}
