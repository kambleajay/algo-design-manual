package sqrtx

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	lo, hi := 2, x/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq < x {
			lo = mid + 1
		} else if sq > x {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return hi
}
