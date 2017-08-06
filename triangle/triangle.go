package triangle

import (
	"math"
)

const testVersion = 3

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
	Deg        // degenerate
)

func KindFromSides(a, b, c float64) Kind {
	hasInf := math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
	hasZero := a == 0 || b == 0 || c == 0
	if hasInf || hasZero {
		return NaT
	}
	if (a+b) > c && (a+c) > b && (b+c) > a {
		if a == b && b == c {
			return Equ
		} else if (a == b && b != c) || (a != b && b == c) || (a == c && a != b) {
			return Iso
		} else if a != b && a != c && b != c {
			return Sca
		}
	} else if (a+b) == c || (a+c) == b || (b+c) == a {
		return Deg
	}

	return NaT
}
