package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m *Matrix) Set(row, col, val int) (ok bool) {
	defer func() { recover() }()
	(*m)[row][col] = val
	ok = true
	return
}

func hasUnevenRows(s string) bool {
	return strings.Count(s, " ")%2 != 0
}

func New(s string) (Matrix, error) {
	if hasUnevenRows(s) {
		return nil, errors.New("uneven rows")
	}
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))
	for i, row := range rows {
		cols := strings.Split(row, " ")
		m[i] = make([]int, len(cols))
		for j, col := range cols {
			c, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			m[i][j] = c
		}

	}
	return m, nil
}

func (m Matrix) Rows() [][]int {
	var rows = make(Matrix, len(m))
	for i := range m {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	numrows, numcols := len(m[0]), len(m)
	var transpose = make(Matrix, numrows)
	for i := range transpose {
		transpose[i] = make([]int, numcols)
		for j := range transpose[i] {
			transpose[i][j] = m[j][i]
		}
	}
	return transpose
}
