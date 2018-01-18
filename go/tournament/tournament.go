package tournament

import (
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

// Struct to hold a team's results
type Result struct {
	name           string
	mp, w, d, l, p int
}

// ResultSlice implements sort.Interface which sorts ascending by points and then name
type ResultSlice []Result

func (r ResultSlice) Len() int      { return len(r) }
func (r ResultSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ResultSlice) Less(i, j int) bool {
	if r[i].p == r[j].p {
		return r[i].name < r[j].name
	}
	return r[i].p > r[j].p
}

// Tally converts input from io.Reader, requiring format of (team; team; outcome) for each line. Returns error if
// format is incorrect or outcome is not win/loss/draw.  Formatted table is written to the supplied io.Writer
func Tally(reader io.Reader, writer io.Writer) error {

	results := make(map[string]*Result)
	buff := make([]byte, 1024)
	reader.Read(buff)
	games := strings.Split(string(buff), "\n")
	for _, game := range games {

		game = strings.Trim(game, " \x00")
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
func WriteResults(results map[string]*Result, w io.Writer) {
	rs := MakeSlice(results)
	sort.Sort(rs)
	w.Write([]byte(fmt.Sprintf(tableFormat, "Team", "MP", "W", "D", "L", "P")))
	for _, v := range rs {
		w.Write([]byte(fmt.Sprintf(tableFormat, v.name, v.mp, v.w, v.d, v.l, v.p)))
	}
}

func MakeSlice(results map[string]*Result) ResultSlice {
	n := len(results)
	slice := make([]Result, n)
	i := 0
	for _, v := range results {
		slice[i] = *v
		i++
	}

	return slice
}

func NewResult(name string) *Result {
	r := new(Result)
	r.name = name
	return r
}

func (r *Result) AddGame(p int) {
	r.mp++
	r.p += p
	switch p {
	case WinP:
		r.w++
	case LossP:
		r.l++
	case DrawP:
		r.d++
	default:
		panic("can't add game")
	}
}

func addOrRetrieveResult(s string, results *map[string]*Result) *Result {
	r, ok := (*results)[s]

	if !ok {
		r = NewResult(s)
		(*results)[s] = r
	}
	return r
}

// splitResults returns two results from the given string and the game outcome.
func splitGame(game string, results *map[string]*Result) (t1 *Result, t2 *Result, outcome string, err error) {
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
