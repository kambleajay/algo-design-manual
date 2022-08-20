package reconstruct_queue

import "sort"

func reconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		height1, kIndex1 := people[i][0], people[i][1]
		height2, kIndex2 := people[j][0], people[j][1]
		if height1 != height2 {
			return height1 > height2
		}
		return kIndex1 <= kIndex2
	})
	res := make([][]int, len(people))
	for _, person := range people {
		if len(res) == 0 {
			res[0] = person
		} else {
			kIndex := person[1]
			copy(res[kIndex+1:], res[kIndex:])
			res[kIndex] = person
		}
	}
	return res
}

func reconstructQueue(people [][]int) [][]int {
	return reconstructQueue2(people)
}
