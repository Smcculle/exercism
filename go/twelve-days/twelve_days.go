package twelve

import (
	"fmt"
	"bytes"
)

const fmtstring = "On the %s day of Christmas my true love gave to me, %s"

var ordinals = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree.",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

func Song() string {
	var song bytes.Buffer
	for day := 1; day < 13; day++ {
		song.WriteString(Verse(day) + "\n")
	}
	return song.String()
}

func Verse(i int) string {
	var total bytes.Buffer
	for j := i; j > 1; j-- {
		total.WriteString(gifts[j])
		total.WriteString(", ")
	}
	if total.Len() == 0 {
		total.WriteString(gifts[1])
	} else {
		total.WriteString(fmt.Sprintf("and %s", gifts[1]))
	}

	return fmt.Sprintf(fmtstring, ordinals[i], total.String())
}
