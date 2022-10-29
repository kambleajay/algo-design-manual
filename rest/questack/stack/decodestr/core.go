package decodestr

import (
	"strings"
	"unicode"
)

type StackOfStrings struct {
	items []string
}

func (ss *StackOfStrings) isEmpty() bool {
	return len(ss.items) == 0
}

func (ss *StackOfStrings) push(s string) {
	ss.items = append(ss.items, s)
}

func (ss *StackOfStrings) pop() string {
	item := ss.items[len(ss.items)-1]
	ss.items = ss.items[:len(ss.items)-1]
	return item
}

type StackOfInts struct {
	items []int
}

func (si *StackOfInts) isEmpty() bool {
	return len(si.items) == 0
}

func (si *StackOfInts) push(n int) {
	si.items = append(si.items, n)
}

func (si *StackOfInts) pop() int {
	item := si.items[len(si.items)-1]
	si.items = si.items[:len(si.items)-1]
	return item
}

func decodeString(s string) string {
	var result strings.Builder
	numStack := StackOfInts{}
	strStack := StackOfStrings{}
	k := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			k = k*10 + int(r-'0')
		} else if r == '[' {
			numStack.push(k)
			strStack.push(result.String())
			result.Reset()
			k = 0
		} else if r == ']' {
			var decodedStr strings.Builder
			decodedStr.WriteString(strStack.pop())
			count := numStack.pop()
			for k := 0; k < count; k++ {
				decodedStr.WriteString(result.String())
			}
			result.Reset()
			result.WriteString(decodedStr.String())
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
