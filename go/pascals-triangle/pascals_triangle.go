package pascal

func Triangle(n int) [][]int {

	var triangle = make([][]int, n)
	if n <= 0 {
		return triangle
	}

	triangle[0] = []int{1}

	for i := 1; i < n; i++ {
		triangle[i] = increment(triangle[i-1])
	}

	return triangle[:n]
}

func increment(row []int) []int {

	next := make([]int, len(row)+1)
	next[0] = 1
	i := 1

	for ; i < len(next)-1; i++ {
		next[i] = row[i] + row[i-1]
	}
	next[i] = 1

	return next
}
