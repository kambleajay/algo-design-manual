package heightchkr

func heightChecker(heights []int) int {
	N := len(heights)
	if N == 0 || N == 1 {
		return 0
	}
	heightFreq := make([]int, 101)
	for _, h := range heights {
		heightFreq[h]++
	}
	height, mismatches := 1, 0
	for _, h := range heights {
		for ; height < len(heightFreq) && heightFreq[height] <= 0; height++ {
		}
		if height != h { //height is in increasing order (1 to 100) - so it should match the next height from array heights, if not we have a mismatch
			mismatches++
		}
		heightFreq[height]-- //we have counted current height once, so decrement it
	}
	return mismatches
}
