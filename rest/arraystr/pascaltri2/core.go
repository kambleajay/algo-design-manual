package pascaltri2

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 0; i < rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
		row[i+1] = 1
	}
	return row
}
