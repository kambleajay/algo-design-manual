package twosum3

type TwoSum struct {
	freq map[int]int
}

func Constructor() TwoSum {
	m := make(map[int]int)
	return TwoSum{m}
}

func (this *TwoSum) Add(number int) {
	this.freq[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num, count := range this.freq {
		complement := value - num
		if complement == num {
			if count > 1 {
				return true
			}
		} else if _, ok := this.freq[complement]; ok {
			return true
		}
	}
	return false
}
