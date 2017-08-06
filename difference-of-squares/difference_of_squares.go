package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	sum := 0
	for c := 1; c <= n; c++ {
		sum += c
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for c := 1; c <= n; c++ {
		sum += c * c
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
