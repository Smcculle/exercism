package twelve

import (
	"bytes"
	"fmt"
)

const fmtstring = "On the %s day of Christmas my true love gave to me, %s"

var ordinals = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = [12]string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	var song bytes.Buffer
	for day := 1; day <= 12; day++ {
		song.WriteString(Verse(day) + "\n")
	}
	return song.String()
}

func Verse(i int) string {
	i--  // convert to 0-index
	var total bytes.Buffer
	for j := i; j > 0; j-- {
		total.WriteString(gifts[j])
		total.WriteString(", ")
	}
	if total.Len() == 0 {
		total.WriteString(gifts[0])
	} else {
		total.WriteString(fmt.Sprintf("and %s", gifts[0]))
	}

	return fmt.Sprintf(fmtstring, ordinals[i], total.String())
}
