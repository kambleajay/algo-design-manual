package range_sum_query

type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	N := len(nums)
	na := NumArray{tree: make([]int, N*2), n: N}
	//fill leaf nodes
	for i, j := N, 0; i < 2*N; i, j = i+1, j+1 {
		na.tree[i] = nums[j]
	}
	//fill parent nodes
	for i := N - 1; i > 0; i-- {
		na.tree[i] = na.tree[i*2] + na.tree[(i*2)+1]
	}
	return na
}

func (nums *NumArray) Update(index int, val int) {
	k := index + nums.n
	nums.tree[k] = val
	for k > 1 {
		k = k / 2
		c1, c2 := k*2, (k*2)+1
		nums.tree[k] = nums.tree[c1] + nums.tree[c2]
	}
}

func (nums *NumArray) SumRange(left int, right int) int {
	l := left + nums.n
	r := right + nums.n
	var sum int
	for l <= r {
		if l%2 == 1 {
			sum += nums.tree[l]
			l++
		}
		if r%2 == 0 {
			sum += nums.tree[r]
			r--
		}
		l /= 2
		r /= 2
	}
	return sum
}
