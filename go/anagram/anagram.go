package anagram

import (
	"reflect"
	"sort"
	"strings"
)

//Detect takes a word and a list of candidates and returns all the candidates
//that are an anagram of the initial word (same letters, rearranged).
//Note: A word is not an anagram of itself.
func Detect(word string, candidates []string) (anagrams []string) {
	w := strings.ToLower(word)
	wr := []rune(w)
	sort.Slice(wr, func(i, j int) bool { return wr[i] < wr[j] })

	for _, candidate := range candidates {
		c := strings.ToLower(candidate)
		cr := []rune(c)
		sort.Slice(cr, func(i, j int) bool { return cr[i] < cr[j] })

		if c != w && reflect.DeepEqual(cr, wr) {
			anagrams = append(anagrams, candidate)
		}
	}
	return
}
