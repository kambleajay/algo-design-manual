package range_sum_query

import "testing"

func TestRangeSumQuery1(t *testing.T) {
	obj := Constructor([]int{1, 3, 5})
	if obj.SumRange(0, 2) != 9 {
		t.Errorf("SumRange(0,2) should be 9, but was %d", obj.SumRange(0, 2))
	}
	obj.Update(1, 2)
	if obj.SumRange(0, 2) != 8 {
		t.Errorf("SumRange(0,2) should be 8, but was %d", obj.SumRange(0, 2))
	}
}

func TestRangeSumQuery2(t *testing.T) {
	obj := Constructor([]int{18, 17, 13, 19, 15, 11, 20, 12, 33, 25})
	answer1 := obj.SumRange(2, 8)
	if answer1 != 123 {
		t.Errorf("SumRange(2,8) should be 123, but was %d", answer1)
	}
	answer2 := obj.SumRange(0, 4)
	if answer2 != 82 {
		t.Errorf("SumRange(0,4) should be 82, but was %d", answer2)
	}
	answer3 := obj.SumRange(4, 5)
	if answer3 != 26 {
		t.Errorf("SumRange(4,5) should be 46, but was %d", answer3)
	}
	answer4 := obj.SumRange(5, 9)
	if answer4 != 101 {
		t.Errorf("SumRange(5,9) should be 101, but was %d", answer4)
	}
}
