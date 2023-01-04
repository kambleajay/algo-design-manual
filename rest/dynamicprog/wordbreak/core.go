package wordbreak

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if dp[i] {
			for _, word := range wordDict {
				right := i + len(word)
				if right <= len(s) && s[i:right] == word {
					dp[right] = true
				}
			}
		}
	}
	return dp[len(s)]
}
