//Package dupzeros provides functionality to duplicate zeros in an array.
package dupzeros

//findLast takes an array arr and its length N, and returns the two positions - slow and fast.
//slow conveys how many elements can fit into the array after putting two zeros for every zero.
//fast conveys what will be the effective length after doubling zeros.
func findLast(arr []int, N int) (int, int) {
	var slow, fast int
	for i := 0; i < N && fast < N; i++ {
		if arr[i] == 0 {
			fast += 2
			slow++
		} else {
			fast++
			slow++
		}
	}
	return slow, fast
}

//isSafeToExpandZero takes a destination index dst, length N and effective length fast,
//and tells whether a zero should be doubled or not. We do not expand a zero if
//it is last (right-most) and when expanded goes beyond the length N.
func isSafeToExpandZero(dst int, N int, fast int) bool {
	return dst < N-1 || (dst == N-1 && fast <= N)
}

//duplicateZeros takes an array arr and modifies it in-place, such that two zeros are
//added for every zero. The length is not changed, so this causes elements on the
//right end to drop off.
func duplicateZeros(arr []int) {
	N := len(arr)
	slow, fast := findLast(arr, N)
	dst, src := N-1, slow-1
	for dst >= 0 && src >= 0 {
		if arr[src] == 0 && isSafeToExpandZero(dst, N, fast) { //if zero is the last char (when dst == N-1) in the source array, we do not want to expand it as it is beyond the effective length
			arr[dst], arr[dst-1] = 0, 0
			dst -= 2
			src--
		} else {
			arr[dst] = arr[src]
			dst--
			src--
		}
	}
}
