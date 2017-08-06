package raindrops

import "strconv"

const testVersion = 3

func Convert(num int) string {
	raindrop := ""

	if num%3 == 0 {
		raindrop += "Pling"
	}
	if num%5 == 0 {
		raindrop += "Plang"
	}
	if num%7 == 0 {
		raindrop += "Plong"
	}

	if len([]rune(raindrop)) == 0 {
		return strconv.Itoa(num)
	}

	return raindrop

}
