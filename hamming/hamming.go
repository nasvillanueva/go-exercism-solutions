package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	aRuneArr := []rune(a)
	bRuneArr := []rune(b)

	if len(aRuneArr) != len(bRuneArr) {
		return -1, errors.New("Lenght mismatch")
	}

	count := 0
	for c := 0; c < len(aRuneArr); c++ {
		if aRuneArr[c] != bRuneArr[c] {
			count++
		}
	}

	return count, nil

}
