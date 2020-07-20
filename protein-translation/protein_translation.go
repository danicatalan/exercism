package protein

import "errors"

var (
	//ErrStop terminates the function execution
	ErrStop = errors.New("STOP")

	//ErrInvalidBase invalid codon
	ErrInvalidBase = errors.New("Invalid codon")
)

var (
	//Translations dictionary with codon to protein translations
	Translations = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
)

//FromRNA takes a RNA strand and returns a secuence of proteins.
func FromRNA(input string) (proteins []string, err error) {
	proteins = make([]string, 0)
	for c := 0; c < len(input); c += 3 {
		protein, err := FromCodon(input[c : c+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return proteins, ErrInvalidBase
		}
		proteins = append(proteins, protein)
	}
	return
}

//FromCodon takes a codon and returns the translated protein.
func FromCodon(input string) (protein string, err error) {
	protein, ok := Translations[input]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return
}
