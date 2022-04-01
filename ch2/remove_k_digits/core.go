package remove_k_digits

import "strings"

func removeKdigits(num string, k int) string {
	var s []uint8
	for i := 0; i < len(num); i++ {
		for j := len(s) - 1; j >= 0 && s[j] > num[i] && k > 0; j-- {
			s = s[:len(s)-1]
			k--
		}
		s = append(s, num[i])
	}
	if k > 0 {
		s = s[:len(s)-k]
	}
	answer := string(s)
	answer = strings.TrimLeft(answer, "0")
	if len(answer) == 0 {
		return "0"
	}
	return answer
}
