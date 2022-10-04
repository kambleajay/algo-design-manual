package singlenum

import "reflect"

func singleNumber1(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			delete(seen, num)
		} else {
			seen[num] = true
		}
	}
	firstKey := reflect.ValueOf(seen).MapKeys()[0]
	return int(firstKey.Int())
}

func singleNumber2(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}

func singleNumber(nums []int) int {
	return singleNumber2(nums)
}
