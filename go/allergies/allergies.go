package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// AllergicTo takes an allergy score and a substance and determines if the person is allergic to that substance.
func AllergicTo(n uint, s string) bool {
	var score uint
	for p, v := range allergens {
		if v == s {
			score = 1 << p
			break
		}
	}
	return n&score > 0
}

// Allergies takes an alergy score and returns the full list of allergies.
func Allergies(n uint) []string {
	var list []string
	for i := range allergens {
		if n&1 == 1 {
			list = append(list, allergens[i])
		}
		n >>= 1
	}
	return list
}
