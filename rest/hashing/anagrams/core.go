package anagrams

import (
	"sort"
	"strings"
)

func sortedString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func groupAnagrams1(strs []string) [][]string {
	keyToAnagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortedString(str)
		_, ok := keyToAnagrams[sortedStr]
		if !ok {
			keyToAnagrams[sortedStr] = []string{str}
		} else {
			keyToAnagrams[sortedStr] = append(keyToAnagrams[sortedStr], str)
		}
	}
	var result [][]string
	for _, anagrams := range keyToAnagrams {
		result = append(result, anagrams)
	}
	return result
}

func joinAsString(counts [26]rune) string {
	var sb strings.Builder
	for _, c := range counts {
		sb.WriteRune(c)
		sb.WriteString("-")
	}
	return sb.String()
}

func makeCountString(s string) string {
	var counts [26]rune
	aCode := int32('a')
	for _, ch := range s {
		counts[ch-aCode]++
	}
	return joinAsString(counts)
}

func groupAnagrams2(strs []string) [][]string {
	keyToAnagrams := make(map[string][]string)
	for _, str := range strs {
		countStr := makeCountString(str)
		_, ok := keyToAnagrams[countStr]
		if !ok {
			keyToAnagrams[countStr] = []string{str}
		} else {
			keyToAnagrams[countStr] = append(keyToAnagrams[countStr], str)
		}
	}
	var result [][]string
	for _, anagrams := range keyToAnagrams {
		result = append(result, anagrams)
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	return groupAnagrams2(strs)
}
