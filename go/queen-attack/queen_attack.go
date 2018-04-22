package queenattack

import (
	"errors"
)

type coord struct {
	x int
	y int
}

func (c coord) inColumn(co coord) bool {
	return c.y == co.y
}

func (c coord) inRow(co coord) bool {
	return c.x == co.x
}

func (c coord) inDiag(co coord) bool {
	return abs(c.x-co.x) == abs(c.y-co.y)
}

func (c coord) canAttack(co coord) bool {
	return c.inColumn(co) || c.inRow(co) || c.inDiag(co)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CanQueenAttack(p1, p2 string) (attack bool, ok error) {

	if p1 == p2 {
		return false, errors.New("queens on same square")
	}

	c1, c2 := getCoord(p1), getCoord(p2)

	if c1 == nil || c2 == nil {
		return false, errors.New("invalid positions")
	}

	canAttack := c1.canAttack(*c2)
	return canAttack, nil

}

func getCoord(s string) *coord {
	xval := int(s[0] - 96)
	yval := int(s[1] - 48)

	if validPos(xval) && validPos(yval) {
		return &coord{xval, yval}
	}

	return nil
}

func validPos(x int) bool {
	return 1 <= x && x <= 8
}
