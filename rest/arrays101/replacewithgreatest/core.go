package replacewithgreatest

func replaceElements(arr []int) []int {
	N := len(arr)
	if N == 0 {
		return nil
	}
	max := arr[N-1]
	for i := N - 1; i >= 0; i-- {
		oldVal := arr[i]
		if i == N-1 {
			arr[i] = -1
		} else {
			arr[i] = max
		}
		if oldVal > max {
			max = oldVal
		}
	}
	return arr
}
