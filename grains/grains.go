package grains

import (
	"errors"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid input")
	}

	return 1 << uint64(n-1), nil
}

func Total() (result uint64) {
	for x := 0; x <= 64; x++ {
		grains, _ := Square(x)
		result += grains
	}
	return result
}
