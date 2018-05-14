package strand

var complement [86]byte

func init() {
	complement['G'] = 'C'
	complement['C'] = 'G'
	complement['T'] = 'A'
	complement['A'] = 'U'
}

// ToRNA takes the given string and returns the nucleotide complements.
func ToRNA(dna string) string {

	var buf = []byte(dna)
	for i, b := range buf {
		buf[i] = complement[b]
	}
	return string(buf)
}
