package romannumerals

import (
	"errors"
	"sort"
)

// Roman abstracts a number in Roman numerals.
type Roman struct {
	Value  int
	Letter string
}

var (
	dictionary = []Roman{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
	}
)

// ToRomanNumeral takes an integer and returns a string with a Roman numeral.
func ToRomanNumeral(number int) (roman string, err error) {
	if number < 1 || number > 3000 {
		return "", errors.New("Number is outside range")
	}

	sort.Slice(dictionary, func(i, j int) bool { return dictionary[i].Value > dictionary[j].Value })

	for _, d := range dictionary {
		for number >= d.Value {
			roman += d.Letter
			number -= d.Value
		}
	}

	return
}
