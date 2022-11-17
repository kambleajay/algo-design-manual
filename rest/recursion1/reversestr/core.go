package reversestr

func reverseStringRecur(i, j int, s []byte) {
	if i > j {
		return
	}
	s[i], s[j] = s[j], s[i]
	reverseStringRecur(i+1, j-1, s)
}

func reverseString(s []byte) {
	reverseStringRecur(0, len(s)-1, s)
}
