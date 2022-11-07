package guessnum

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return -1
}

func guessNumber(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		g := guess(mid)
		if g == 1 {
			lo = mid + 1
		} else if g == -1 {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
