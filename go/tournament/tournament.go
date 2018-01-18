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

// Team holds a single team's tournament results.
type Team struct {
	name           string
	mp, w, d, l, p int
}

func (t *Team) AddWin() {
	t.mp++
	t.w++
	t.p += WinP
}

func (t *Team) AddLoss() {
	t.mp++
	t.l++
	t.p += LossP
}

func (t *Team) AddDraw() {
	t.mp++
	t.d++
	t.p += DrawP
}

func (t *Team) String() string {
	return fmt.Sprintf(tableFormat, t.name, t.mp, t.w, t.d, t.l, t.p)
}

// TeamSlice implements sort.Interface using points and then name as the sort key.
type TeamSlice []Team

func (t TeamSlice) Len() int      { return len(t) }
func (t TeamSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TeamSlice) Less(i, j int) bool {
	if t[i].p == t[j].p {
		return t[i].name < t[j].name
	}
	return t[i].p > t[j].p
}

type Results map[string]*Team

// WriteTo sorts results and sends it to the specified io.Writer.
func (results Results) WriteTo(w io.Writer) (int64, error) {
	var err error
	var n int
	rs := results.MakeSlice()
	sort.Sort(rs)
	written, _ := io.WriteString(w, fmt.Sprintf(tableFormat, "Team", "MP", "W", "D", "L", "P"))
	for _, v := range rs {
		written, err = io.WriteString(w, fmt.Sprintf(tableFormat, v.name, v.mp, v.w, v.d, v.l, v.p))
		n += written

		if err != nil {
			break
		}
	}

	return int64(n), err
}

func (results Results) MakeSlice() TeamSlice {
	n := len(results)
	slice := make([]Team, n)
	i := 0
	for _, v := range results {
		slice[i] = *v
		i++
	}

	return slice
}

func (results Results) addOrRetrieve(s string) *Team {
	r, ok := results[s]

	if !ok {
		r = NewResult(s)
		results[s] = r
	}
	return r
}

func (results Results) Add(game string) (err error) {
	team1, team2, outcome, err := splitGame(game)
	t1, t2 := results.addOrRetrieve(team1), results.addOrRetrieve(team2)
	if err != nil {
		return err
	}

	if outcome == "win" {
		t1.AddWin()
		t2.AddLoss()
	} else if outcome == "loss" {
		t1.AddLoss()
		t2.AddWin()
	} else if outcome == "draw" {
		t1.AddDraw()
		t2.AddDraw()
	} else {
		return errors.New("incorrect format")
	}

	return nil
}

// Tally converts input from io.Reader, requiring format of (team; team; outcome) for each line. Returns error if
// format is incorrect or outcome is not win/loss/draw.  Formatted table is written to the supplied io.Writer.
func Tally(reader io.Reader, writer io.Writer) error {

	results := make(Results)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		game := scanner.Text()
		if len(game) == 0 || game[0] == '#' {
			continue
		}

		err := results.Add(game)
		if err != nil {
			return err
		}

	}

	results.WriteTo(writer)

	return nil
}

// splitResults returns two results from the given string and the game outcome.
func splitGame(game string) (t1 string, t2 string, outcome string, err error) {
	err = nil
	split := strings.Split(game, ";")
	if len(split) != 3 {
		err = fmt.Errorf("incorrect format")
		return
	}
	t1, t2, outcome = split[0], split[1], split[2]
	return
}

func NewResult(name string) *Team {
	r := new(Team)
	r.name = name
	return r
}
