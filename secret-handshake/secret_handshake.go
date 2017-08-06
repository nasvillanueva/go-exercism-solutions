package secret

import (
	"fmt"
)

const testVersion = 2

func Handshake(code uint) []string {
	secret := []string{}
	binaryCode := []rune(toReversedBinary(code))
	for index, val := range binaryCode {
		if index == 0 && isOn(val) {
			secret = append(secret, "wink")
		}
		if index == 1 && isOn(val) {
			secret = append(secret, "double blink")
		}
		if index == 2 && isOn(val) {
			secret = append(secret, "close your eyes")
		}
		if index == 3 && isOn(val) {
			secret = append(secret, "jump")
		}
		if index == 4 && isOn(val) {
			secret = reverseSecret(secret)
		}
	}
	return secret
}

func isOn(bit rune) bool {
	return string(bit) == "1"
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
