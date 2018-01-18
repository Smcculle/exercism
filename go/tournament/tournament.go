// Package tournament can be used to keep score in a 1v1 team tournament, creates a formatted table as output.
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Point values for games
const (
	WinP  = 3
	DrawP = 1
	LossP = 0
)

const tableFormat = "%-30v |%3v |%3v |%3v |%3v |%3v\n"

// Team holds a single team's tournament results
type Team struct {
	name           string
	mp, w, d, l, p int
}

func (t *Team) AddGame(p int) {
	t.mp++
	t.p += p
	switch p {
	case WinP:
		t.w++
	case LossP:
		t.l++
	case DrawP:
		t.d++
	default:
		panic("can't add game")
	}
}

func (t *Team) String() string {
	return fmt.Sprintf(tableFormat, t.name, t.mp, t.w, t.d, t.l, t.p)
}

// TeamSlice implements sort.Interface which sorts ascending by points and then name
type TeamSlice []Team

func (t TeamSlice) Len() int      { return len(t) }
func (t TeamSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TeamSlice) Less(i, j int) bool {
	if t[i].p == t[j].p {
		return t[i].name < t[j].name
	}
	return t[i].p > t[j].p
}

// Tally converts input from io.Reader, requiring format of (team; team; outcome) for each line. Returns error if
// format is incorrect or outcome is not win/loss/draw.  Formatted table is written to the supplied io.Writer.
func Tally(reader io.Reader, writer io.Writer) error {

	results := make(map[string]*Team)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		game := scanner.Text()
		if len(game) == 0 || game[0] == '#' {
			continue
		}

		t1, t2, outcome, err := splitGame(game, &results)

		if err != nil {
			return err
		}

		if outcome == "win" {
			t1.AddGame(WinP)
			t2.AddGame(LossP)
		} else if outcome == "loss" {
			t1.AddGame(LossP)
			t2.AddGame(WinP)
		} else if outcome == "draw" {
			t1.AddGame(DrawP)
			t2.AddGame(DrawP)
		} else {
			return errors.New("incorrect format")
		}
	}

	WriteResults(results, writer)

	return nil
}

func WriteResults(results map[string]*Team, w io.Writer) {
	rs := MakeSlice(results)
	sort.Sort(rs)
	w.Write([]byte(fmt.Sprintf(tableFormat, "Team", "MP", "W", "D", "L", "P")))
	for _, v := range rs {
		w.Write([]byte(fmt.Sprintf(tableFormat, v.name, v.mp, v.w, v.d, v.l, v.p)))
	}
}

func MakeSlice(results map[string]*Team) TeamSlice {
	n := len(results)
	slice := make([]Team, n)
	i := 0
	for _, v := range results {
		slice[i] = *v
		i++
	}

	return slice
}

func NewResult(name string) *Team {
	r := new(Team)
	r.name = name
	return r
}

func addOrRetrieveResult(s string, results *map[string]*Team) *Team {
	r, ok := (*results)[s]

	if !ok {
		r = NewResult(s)
		(*results)[s] = r
	}
	return r
}

// splitResults returns two results from the given string and the game outcome.
func splitGame(game string, results *map[string]*Team) (t1 *Team, t2 *Team, outcome string, err error) {
	err = nil
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()
	split := strings.Split(game, ";")
	if len(split) != 3 {
		panic("incorrect format")
	}
	t1, t2 = addOrRetrieveResult(split[0], results), addOrRetrieveResult(split[1], results)
	outcome = split[2]
	return
}
