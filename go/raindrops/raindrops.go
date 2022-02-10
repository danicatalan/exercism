package raindrops

import "strconv"

// Sound abstracts the rules of raindrops.
type Sound struct {
	Factor int
	Value  string
}

var (
	sounds = []Sound{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
)

// Convert takes an integer and returns the sound in raindrops.
func Convert(number int) (raindrops string) {
	for _, s := range sounds {
		if number%s.Factor == 0 {
			raindrops += s.Value
		}
	}
	if raindrops == "" {
		raindrops = strconv.Itoa(number)
	}
	return
}
