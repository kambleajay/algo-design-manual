package diatrav

func saveDiagonals(mat [][]int) (map[int][]int, int) {
	diagonals := make(map[int][]int)
	var maxSum int
	for row := 0; row < len(mat); row++ {
		for column := 0; column < len(mat[row]); column++ {
			currSum := row + column
			_, ok := diagonals[currSum]
			if !ok {
				diagonals[currSum] = []int{mat[row][column]}
			} else {
				diagonals[currSum] = append(diagonals[currSum], mat[row][column])
			}
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}
	return diagonals, maxSum
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func findDiagonalOrder(mat [][]int) []int {
	var result []int
	diagonals, maxSum := saveDiagonals(mat)
	for i := 0; i <= maxSum; i++ {
		currDiagonal := diagonals[i]
		if i%2 == 0 {
			reverse(currDiagonal)
		}
		result = append(result, currDiagonal...)
	}
	return result
}
