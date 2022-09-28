package pascaltri

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	size := 1
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if j == 0 || j == size-1 {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		size++
	}
	return result
}
