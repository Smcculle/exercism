package house

import "strings"

var rhyme = []string{
	"house that Jack built.",
	"malt\nthat lay in",
	"rat\nthat ate",
	"cat\nthat killed",
	"dog\nthat worried",
	"cow with the crumpled horn\nthat tossed",
	"maiden all forlorn\nthat milked",
	"man all tattered and torn\nthat kissed",
	"priest all shaven and shorn\nthat married",
	"rooster that crowed in the morn\nthat woke",
	"farmer sowing his corn\nthat kept",
	"horse and the hound and the horn\nthat belonged to",
}

var verses = make([]string, len(rhyme))

func Song() string {

	var builder strings.Builder

	builder.WriteString(Verse(1))

	for i := 2; i <= len(rhyme); i++ {
		builder.WriteString("\n\n")
		builder.WriteString(Verse(i))
	}
	return builder.String()
}

//Verse uses memoization to avoid repeating work
func Verse(n int) string {
	v := n - 1

	if verses[v] != "" {
		return verses[v]
	}

	if v == 0 {
		verses[v] = "This is the " + rhyme[0]
		return verses[v]
	}

	verses[v] = "This is the " + rhyme[v] + Verse(v)[7:]
	return verses[v]
}
