// Package space calculates age on different planets
package space

import (
	"fmt"
	"strconv"
)

const ref = 31557600

type Planet string

var pmap = map[Planet]float64{
	"Earth":   ref,
	"Mercury": 0.2408467 * ref,
	"Venus":   0.61519726 * ref,
	"Mars":    1.8808158 * ref,
	"Jupiter": 11.862615 * ref,
	"Saturn":  29.447498 * ref,
	"Uranus":  84.016846 * ref,
	"Neptune": 164.79132 * ref,
}

// Age returns number of planet-years f seconds is
func Age(f float64, p Planet) float64 {

	return round(f / pmap[p])
}

// Round uses the worst method I could think of to supply the float64 answers to 2 decimal places
// since go does not have a math.round ??
func round(f float64) float64 {

	rounded, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	return rounded
}