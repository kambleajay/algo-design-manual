package isomor

func isIsomorphic(s string, t string) bool {
	mapSToT := make(map[uint8]uint8)
	mapTToS := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		chs, cht := s[i], t[i]
		prevSToTMapping, ok1 := mapSToT[chs]
		prevToToSMapping, ok2 := mapTToS[cht]
		if ok1 && prevSToTMapping != cht {
			return false
		}
		if ok2 && prevToToSMapping != chs {
			return false
		}
		mapSToT[chs] = cht
		mapTToS[cht] = chs
	}
	return true
}
