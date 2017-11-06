package raindrops

import "fmt"

func Convert(i int) (o string) {
	if i % 3 == 0 {
		o += "Pling"
	}
	if i % 5 == 0 {
		o += "Plang"
	}
	if i % 7 == 0 {
		o += "Plong"
	}

	if o == "" {
		o = fmt.Sprintf("%v", i)
	}

	return
}
