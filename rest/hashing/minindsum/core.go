package minindsum

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	keyToIndexList1 := make(map[string]int)
	for i, s := range list1 {
		keyToIndexList1[s] = i
	}
	minSum := math.MaxInt64
	var result []string
	for i, s := range list2 {
		j, ok := keyToIndexList1[s]
		if ok {
			sum := i + j
			if sum < minSum {
				result = []string{s}
				minSum = sum
			} else if sum == minSum {
				result = append(result, s)
			}
		}
	}
	return result
}
