package twelve

import (
	"fmt"
	"strings"
)

const fmtString = "On the %s day of Christmas my true love gave to me, "

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

var gifts = [13]string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
	"and a Partridge in a Pear Tree.",
}

func Song() string {
	var song strings.Builder
	for day := 1; day <= 12; day++ {
		song.WriteString(Verse(day) + "\n")
	}
	return song.String()
}

func Verse(i int) string {
	i-- // convert to 0-index
	var verse strings.Builder
	if i == 0 {
		return fmt.Sprintf(fmtString, ordinals[0]) + gifts[0]
	}

	verse.WriteString(fmt.Sprintf(fmtString, ordinals[i]))
	for ; i > 0; i-- {
		verse.WriteString(gifts[i])
	}
	verse.WriteString(gifts[12])
	return verse.String()
}
