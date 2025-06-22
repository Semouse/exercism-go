package strand

var conversion = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, value := range dna {
		rna[i] = conversion[value]
	}

	return string(rna)
}
