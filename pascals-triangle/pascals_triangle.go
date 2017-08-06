package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	for x := 0; x < n; x++ {
		triangle[x] = make([]int, x+1)
		for y := 0; y < len(triangle[x]); y++ {
			if x-1 > 0 && y > 0 && y < len(triangle[x-1]) {
				triangle[x][y] = triangle[x-1][y-1] + triangle[x-1][y]
			} else {
				triangle[x][y] = 1
			}
		}
	}

	return triangle
}
