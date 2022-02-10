package strand

var (
	transcription = map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
)

// ToRNA takes a DNA strand and returns its RNA complement.
func ToRNA(dna string) (rna string) {
	for _, n := range dna {
		rna += transcription[n]
	}
	return
}
