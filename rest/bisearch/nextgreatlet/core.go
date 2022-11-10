package nextgreatlet

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return letters[lo%len(letters)]
}
