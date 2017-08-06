package pangram

import (
	"bytes"
	"strings"
)

const testVersion = 1

const letters = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(input string) bool {
	inputBytes := []byte(strings.ToLower(input))
	for index := range letters {
		if bytes.IndexByte(inputBytes, letters[index]) == -1 {
			return false
		}
	}
	return true
}
