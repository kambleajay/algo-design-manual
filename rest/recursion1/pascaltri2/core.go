package pascaltri2

func pascalVal(i, j int) int {
	if j == 0 || j == i {
		return 1
	}
	return pascalVal(i-1, j-1) + pascalVal(i-1, j)
}

func getRow(rowIndex int) []int {
	d := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		d[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				d[i][j] = 1
			} else {
				d[i][j] = d[i-1][j-1] + d[i-1][j]
			}
		}
	}
	return d[rowIndex]
}
