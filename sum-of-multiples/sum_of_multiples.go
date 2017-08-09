package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) (total int) {
	for x := 1; x < limit; x++ {
		for _, divisor := range divisors {
			num := limit - x
			if num%divisor == 0 {
				total += num
				break
			}
		}
	}
	return total
}
