package dfs

import (
	"algo/testing/utils"
	"testing"
)

func TestFindWords(t *testing.T) {
	var tests = []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}, nil},
		{[][]byte{{'a'}}, []string{"a"}, []string{"a"}},
		{[][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}}, []string{"oa", "oaa"}, []string{"oa", "oaa"}},
		{[][]byte{{'a', 'a'}}, []string{"aaa"}, nil},
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain", "hklf", "hf"}, []string{"oath", "eat", "hklf", "hf"}},
		{[][]byte{{'a', 'b', 'e'}, {'b', 'c', 'd'}}, []string{"abcdeb"}, []string{"abcdeb"}},
	}
	for _, test := range tests {
		if got := findWords(test.board, test.words); !utils.ContainsAllStringSlice(got, test.want) {
			t.Errorf("findWords(%v, %v) = %v, want %v", test.board, test.words, got, test.want)
		}
	}
}

func BenchmarkFindWords(b *testing.B) {
	board := [][]byte{{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}}
	words := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i < b.N; i++ {
		findWords(board, words)
	}
}
