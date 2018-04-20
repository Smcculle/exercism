package secret

var maskF = []uint{1, 2, 4, 8}
var maskR = []uint{8, 4, 2, 1}
var code = map[uint]string{1: "wink", 2: "double blink", 4: "close your eyes", 8: "jump"}

func Handshake(n uint) []string {

	result := make([]string, 0, 4)
	var mask []uint

	if n&16 > 0 {
		mask = maskR
	} else {
		mask = maskF
	}

	for _, m := range mask {
		if n&m > 0 {
			result = append(result, code[m])
		}
	}

	return result
}
