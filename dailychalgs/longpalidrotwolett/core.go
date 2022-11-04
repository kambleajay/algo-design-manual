package longpalidrotwolett

func reverse(s string) string {
	return string(s[1]) + string(s[0])
}

func isSameLetters(s string) bool {
	return s[0] == s[1]
}

func isEven(n int) bool {
	return (n & 1) == 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestPalindrome(words []string) int {
	freqs := make(map[string]int)
	for _, word := range words {
		freqs[word]++
	}
	var noOfWords int
	var central bool
	for word, count := range freqs {
		if isSameLetters(word) {
			if isEven(count) {
				noOfWords += count
			} else {
				noOfWords += count - 1
				central = true
			}
		} else if word[0] < word[1] {
			reverseWord := reverse(word)
			noOfWords += 2 * min(count, freqs[reverseWord])
		}
	}
	if central {
		noOfWords++
	}
	return 2 * noOfWords
}
