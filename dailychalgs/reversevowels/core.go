package reversevowels

var vowels = []uint8{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func isVowel(ch uint8) bool {
	for _, vowel := range vowels {
		if ch == vowel {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	lo, hi := 0, len(bytes)-1
	for {
		for lo < len(bytes) && !isVowel(bytes[lo]) {
			lo++
		}
		for hi >= 0 && !isVowel(bytes[hi]) {
			hi--
		}
		if lo >= hi {
			break
		}
		bytes[lo], bytes[hi] = bytes[hi], bytes[lo]
		lo++
		hi--
	}
	return string(bytes)
}
