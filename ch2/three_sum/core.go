package three_sum

import "sort"

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	var res [][]int
	sort.Sort(IntSlice(nums))
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			twoSum(nums, i, &res)
		}
	}
	return res
}

func twoSum(nums []int, i int, res *[][]int) {
	lo, hi := i+1, len(nums)-1
	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		} else {
			*res = append(*res, []int{nums[i], nums[lo], nums[hi]})
			lo++
			hi--
			for ; lo < hi && nums[lo-1] == nums[lo]; lo++ {
			}
		}
	}
}
