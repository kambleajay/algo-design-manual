package countvowels

const N = int(1e9) + 7

func sum(counts []int) int {
	var s int
	for _, c := range counts {
		s += c
	}
	return s
}

func countVowelPermutation1(n int) int {
	children := [][]int{{1}, {0, 2}, {0, 1, 3, 4}, {2, 4}, {0}}
	prevCounts := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		nextCounts := make([]int, 5)
		for i := 0; i < 5; i++ {
			for _, child := range children[i] {
				multiplier := prevCounts[i]
				nextCounts[child] += multiplier
			}
		}
		for i := 0; i < len(nextCounts); i++ {
			nextCounts[i] = nextCounts[i] % N
		}
		prevCounts = nextCounts
	}
	return sum(prevCounts) % N
}

func countVowelPermutation(n int) int {
	aCount, eCount, iCount, oCount, uCount := 1, 1, 1, 1, 1
	var aCountNew, eCountNew, iCountNew, oCountNew, uCountNew int
	for i := 1; i < n; i++ {
		aCountNew = (eCount + iCount + uCount) % N
		eCountNew = (aCount + iCount) % N
		iCountNew = (eCount + oCount) % N
		oCountNew = iCount % N
		uCountNew = (iCount + oCount) % N
		aCount, eCount, iCount, oCount, uCount = aCountNew, eCountNew, iCountNew, oCountNew, uCountNew
	}
	return (aCount + eCount + iCount + oCount + uCount) % N
}
