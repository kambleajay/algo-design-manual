package firunich

func firstUniqChar(s string) int {
	freq := make(map[int32]int)
	for _, ch := range s {
		freq[ch]++
	}
	for i, ch := range s {
		count, ok := freq[ch]
		if ok && count == 1 {
			return i
		}
	}
	return -1
}
