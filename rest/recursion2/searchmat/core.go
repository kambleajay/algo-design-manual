package searchmat

func helper(matrix [][]int, target int, rowLo, rowHi, colLo, colHi int) bool {
	if rowLo > rowHi || colLo > colHi {
		return false
	}
	rowMid := rowLo + (rowHi-rowLo)/2
	colMid := colLo + (colHi-colLo)/2
	pivot := matrix[rowMid][colMid]
	if pivot == target {
		return true
	}
	if target < pivot {
		return helper(matrix, target, rowLo, rowMid-1, colLo, colMid-1) ||
			helper(matrix, target, rowLo, rowMid-1, colMid, colHi) ||
			helper(matrix, target, rowMid, rowHi, colLo, colMid-1)
	} else if target > pivot {
		return helper(matrix, target, rowMid+1, rowHi, colMid+1, colHi) ||
			helper(matrix, target, rowMid+1, rowHi, colLo, colMid) ||
			helper(matrix, target, rowLo, rowMid, colMid+1, colHi)
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return helper(matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}
