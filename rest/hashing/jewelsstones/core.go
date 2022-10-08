package jewelsstones

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[rune]bool)
	for _, j := range jewels {
		jewelsMap[j] = true
	}
	var countOfJewels int
	for _, stone := range stones {
		if jewelsMap[stone] {
			countOfJewels++
		}
	}
	return countOfJewels
}
