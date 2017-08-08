package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

func CanQueenAttack(w, b string) (attack bool, err error) {
	wx, wy, err := parseNotation(w)
	if err != nil {
		return
	}
	bx, by, err := parseNotation(b)
	if err != nil {
		return
	}
	if wx == bx && wy == by {
		err = errors.New("Piece can't be on the same tile")
	} else if checkInvalidPos([]int{wx, wy, bx, by}) {
		err = errors.New("Pieces can't be out of the board")
	} else if (wx == bx && wy != by) || (wx != bx && wy == by) {
		attack = true
	} else if math.Abs(float64(wx-bx)) == math.Abs(float64(wy-by)) {
		attack = true
	}

	return
}

func parseNotation(pos string) (int, int, error) {
	if len(pos) != 2 {
		return -1, -1, errors.New("Invalid notation")
	}
	// Get the difference of the ASCII decimal value of 'h' (105) and the ascii decmal value of
	// the letter in the Algebraic notation then subtract it with 9 to get something between 1-8
	//
	// note: I could have used the ASCII Decimal value of 'i' (104) and then add 1 then subtract it
	// with 9. but to reduce the computation, i used 'h' (105)
	x := (9 - (105 - int(pos[0])))
	y := int(pos[1] - '0')
	return x, y, nil
}

func checkInvalidPos(pos []int) bool {
	for _, val := range pos {
		if val <= 0 || val > 8 {
			return true
		}
	}
	return false
}
