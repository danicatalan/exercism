// Package proverb Given a list of inputs, generate the relevant proverb.
package proverb

import "fmt"

const (
	paragraph     = "For want of a %s the %s was lost."
	lastParagraph = "And all for the want of a %s."
)

// Proverb takes a list of strings and generates a proverb.
func Proverb(rhyme []string) (proverb []string) {
	for i := range rhyme {
		if i == len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf(lastParagraph, rhyme[0]))
		} else {
			proverb = append(proverb, fmt.Sprintf(paragraph, rhyme[i], rhyme[i+1]))
		}
	}
	return
}
