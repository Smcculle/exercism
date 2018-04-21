package house

import "strings"

var rhyme = []string{
	"the house that Jack built.",
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
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

func Verse(n int) string {
	v := n - 1

	if verses[v] != "" {
		return verses[v]
	}

	if v == 0 {
		verses[v] = "This is " + rhyme[0]
		return verses[v]
	}

	verses[v] = "This is " + rhyme[v] + Verse(v)[7:]
	return verses[v]
}
