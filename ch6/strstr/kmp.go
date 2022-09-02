package strstr

func makeCharToIntMap() map[uint8]int {
	chars := []uint8{'a', 'b', 'c', 'd', 'e',
		'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y',
		'z'}
	charToInt := make(map[uint8]int)
	for i, c := range chars {
		charToInt[c] = i
	}
	return charToInt
}

func makeDFA(pat string, r, m int, charToInt map[uint8]int) [][]int {
	dfa := make([][]int, r)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}
	dfa[charToInt[pat[0]]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < r; c++ {
			dfa[c][j] = dfa[c][x]
		}
		dfa[charToInt[pat[j]]][j] = j + 1
		x = dfa[charToInt[pat[j]]][x]
	}
	return dfa
}

func search(txt, pat string, dfa [][]int, m int, charToInt map[uint8]int) int {
	var i, j int
	n := len(txt)
	for ; i < n && j < m; i++ {
		j = dfa[charToInt[txt[i]]][j]
	}
	if j == m {
		return i - m
	} else {
		return -1
	}

}

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	charToInt := makeCharToIntMap()
	pat := needle
	m := len(pat)
	r := 26
	dfa := makeDFA(pat, r, m, charToInt)
	return search(haystack, pat, dfa, m, charToInt)
}
