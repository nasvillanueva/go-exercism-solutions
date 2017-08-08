package secret

import (
	"fmt"
)

const testVersion = 2

var messages = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(code uint) []string {
	secret := []string{}
	binaryCode := []rune(toReversedBinary(code))
	for index, val := range binaryCode {
		if val == 49 {
			if index >= 0 && index <= 3 {
				secret = append(secret, messages[index])
			} else if index == 4 {
				secret = reverseSecret(secret)
			}
		}
	}
	return secret
}

func toReversedBinary(n uint) string {
	if n > 0 {
		a := n % 2
		// Originally this was supposed to return "%s%d"
		// But for the sake of traversing it in reverse, I'll return "%d%s"
		return fmt.Sprintf("%d%s", a, toReversedBinary(n/2))
	}
	return ""
}

func reverseSecret(secret []string) []string {
	for i, j := 0, len(secret)-1; i < j; i, j = i+1, j-1 {
		secret[i], secret[j] = secret[j], secret[i]
	}
	return secret
}
