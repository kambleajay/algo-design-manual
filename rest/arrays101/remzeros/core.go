package remzeros

//removeElement removes val from nums array, and it returns
//the new length of the array after subtracting count of val.
func removeElement(nums []int, val int) int {
	var reader, writer int
	N := len(nums)
	for reader = 0; reader < N; reader++ {
		if nums[reader] != val {
			nums[writer] = nums[reader]
			writer++
		}
	}
	return writer
}
