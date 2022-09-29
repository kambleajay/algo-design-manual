package lngprfx

import (
	"math"
	"strings"
)

func charAtEqualForAll(i int, strs []string) bool {
	char := strs[0][i]
	for _, s := range strs[1:] {
		if s[i] != char {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	smallestStrLen := math.MaxInt32
	for _, s := range strs {
		if len(s) < smallestStrLen {
			smallestStrLen = len(s)
		}
	}
	var result strings.Builder
	for i := 0; i < smallestStrLen; i++ {
		if !charAtEqualForAll(i, strs) {
			break
		} else {
			result.WriteByte(strs[0][i])
		}
	}
	return result.String()
}
