package revwords3

import "strings"

func reverseString(s string) string {
	arr := []byte(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}
	return strings.Join(words, " ")
}
