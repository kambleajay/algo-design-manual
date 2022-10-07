package sudoku

func box(row, col int) int {
	return (row/3)*3 + col/3
}

func mapOfSets() map[int]map[byte]bool {
	mofs := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		mofs[i] = make(map[byte]bool)
	}
	return mofs
}

func isValidSudoku(board [][]byte) bool {
	rowNums := mapOfSets()
	colNums := mapOfSets()
	boxNums := mapOfSets()
	for i, row := range board {
		for j, val := range row {
			if val != '.' {
				if rowNums[i][val] {
					return false
				} else {
					rowNums[i][val] = true
				}
				if colNums[j][val] {
					return false
				} else {
					colNums[j][val] = true
				}
				boxNum := box(i, j)
				if boxNums[boxNum][val] {
					return false
				} else {
					boxNums[boxNum][val] = true
				}
			}
		}
	}
	return true
}
