package uniwordabbr

import "fmt"

type ValidWordAbbr struct {
	abrvToWords map[string]map[string]bool
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := make(map[string]map[string]bool)
	for _, word := range dictionary {
		abrv := abbreviation(word)
		_, ok := m[abrv]
		if !ok {
			m[abrv] = make(map[string]bool)
			m[abrv][word] = true
		} else {
			m[abrv][word] = true
		}
	}
	dict := ValidWordAbbr{m}
	return dict
}

func abbreviation(s string) string {
	if len(s) <= 2 {
		return s
	}
	return fmt.Sprintf("%c%d%c", s[0], len(s)-2, s[len(s)-1])
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abrv := abbreviation(word)
	words, ok := this.abrvToWords[abrv]
	if !ok { //abbreviation not found
		return true
	}
	if len(words) == 1 && words[word] {
		return true
	}
	return false
}
