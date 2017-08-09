package etl

const testVersion = 1

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}

	for score, letters := range in {
		for _, letter := range letters {
			out[string(letter[0] + 32)] = score
		}
	}

	return out
}