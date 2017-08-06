package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(longName string) string {
	words := regexp.MustCompile("( |-)").Split(longName, -1)
	abbr := ""
	for _, value := range words {
		abbr += string(value[0])
	}
	return strings.ToUpper(abbr)
}
