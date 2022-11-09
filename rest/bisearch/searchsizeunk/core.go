package searchsizeunk

import "math"

type ArrayReader struct {
	items []int
}

func (this *ArrayReader) get(index int) int {
	if index < 0 || index > len(this.items)-1 {
		return math.MaxInt32
	}
	return this.items[index]
}

func NewArrayReader(items []int) ArrayReader {
	return ArrayReader{items}
}

func search(reader ArrayReader, target int) int {
	if reader.get(0) == target {
		return 0
	}
	var x int
	lo, hi := 0, 1
	for {
		x = reader.get(hi)
		if x >= target {
			break
		}
		hi = hi * 2
	}
	if x == target {
		return hi
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		x := reader.get(mid)
		if target < x {
			hi = mid - 1
		} else if target > x {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
