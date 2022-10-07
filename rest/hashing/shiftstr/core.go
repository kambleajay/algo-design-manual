package shiftstr

import "strings"

const A = int('a')

func shiftLeft(r rune, shift int) rune {
	i := (int(r)-shift+26)%26 + A
	return rune(i)
}

func hash(str string) string {
	var sb strings.Builder
	first := int(rune(str[0]))
	for _, s := range str {
		ch := shiftLeft(s, first)
		sb.WriteRune(ch)
	}
	return sb.String()
}

func groupStrings(strings []string) [][]string {
	codeToStrings := make(map[string][]string)
	for _, s := range strings {
		h := hash(s)
		_, ok := codeToStrings[h]
		if !ok {
			codeToStrings[h] = []string{s}
		} else {
			codeToStrings[h] = append(codeToStrings[h], s)
		}
	}
	var result [][]string
	for _, v := range codeToStrings {
		result = append(result, v)
	}
	return result
}
