package decodeways

import "strconv"

func ways(s string, i int, memo map[int]int) int {
	ans, ok := memo[i]
	if ok {
		return ans
	}
	if i == len(s) {
		return 1
	}
	if s[i] == '0' {
		return 0
	}
	//there is only 1 digit which is not zero
	if i == len(s)-1 {
		return 1
	}
	answer := ways(s, i+1, memo)
	twoDigit, _ := strconv.Atoi(s[i : i+2])
	if twoDigit >= 10 && twoDigit <= 26 {
		answer += ways(s, i+2, memo)
	}
	memo[i] = answer
	return answer
}

func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		twoDigitStr := s[i-2 : i]
		twoDigit, _ := strconv.Atoi(twoDigitStr)
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
	//memo := make(map[int]int)
	//return ways(s, 0, memo)
}
