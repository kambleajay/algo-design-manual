package counting_bits

func countBits(n int) []int {
	var answer []int
	for i := 0; i < n+1; i++ {
		if i == 0 {
			answer = append(answer, 0)
		} else if i == 1 {
			answer = append(answer, 1)
		} else {
			quot, rem := i/2, i%2
			if rem == 0 {
				answer = append(answer, answer[quot])
			} else {
				answer = append(answer, answer[quot]+1)
			}
		}
	}
	return answer
}
