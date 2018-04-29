package protein

// codon to amino acid mapping
var codons = map[string]string{
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

// FromCodon returns the amino acid corresponding to the given codon
func FromCodon(codon string) string {
	return codons[codon]
}

// FromRNA returns a slice of amino acids retrieved from each codon in rna until
// STOP is detected
func FromRNA(rna string) (aminos []string) {
	for i := 0; i+3 <= len(rna); i += 3 {
		amino := codons[rna[i:i+3]]
		if amino == "STOP" {
			break
		}
		aminos = append(aminos, amino)
	}

	return
}
