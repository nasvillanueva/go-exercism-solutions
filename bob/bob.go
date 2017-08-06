package bob

import (
	"strings"
)

const testVersion = 3

func Hey(input string) string {
	in := strings.TrimSpace(input)
	if in == "" {
		return "Fine. Be that way!"
	} else if containsAlpha(in) && in == strings.ToUpper(in) {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(in, "?") {
		return "Sure."
	}
	return "Whatever."
}

func containsAlpha(str string) bool {
	for _, value := range []byte(str) {
		charCode := int32(value)
		// Char code from ASCII Table http://ascii.cl
		// 'A' == 65, 'Z' == 90, 'a' == 97, 'z' == 122
		if (charCode >= 65 && charCode <= 90) || (charCode >= 97 && charCode <= 122) {
			return true
		}
	}

	return false
}
