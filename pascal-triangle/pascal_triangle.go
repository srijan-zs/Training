package pascal

// pascal function creates a slice of slices of type int which stores the different rows of pascal's triangle
func pascal(num int) [][]int {
	if num <= 0 {
		return nil
	}

	result := make([][]int, num)

	for i := 0; i < num; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1

		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
	}

	return result
}
