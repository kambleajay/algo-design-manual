package longsubs

func lengthOfLongestSubstring1(s string) int {
	var max, curMax int
	seen := make(map[rune]bool)
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			curMax++
		} else {
			if curMax > max {
				max = curMax
			}
			seen = make(map[rune]bool)
			seen[r] = true
			curMax = 1
		}
	}
	if curMax > max {
		max = curMax
	}
	return max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	var ans int
	chToIndex := make(map[uint8]int)
	for left, right := 0, 0; right < len(s); right++ {
		ch := s[right]
		_, ok := chToIndex[ch]
		if ok {
			left = max(chToIndex[ch], left)
		}
		ans = max(ans, right-left+1)
		chToIndex[ch] = right + 1
	}
	return ans
}
