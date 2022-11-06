package mediandatastr

import "testing"

func TestMedianFinder(t *testing.T) {
	testSets := [][][]interface{}{{{1, 1.0}, {2, 1.5}, {3, 2.0}, {4, 2.5}, {5, 3.0}, {6, 3.5}, {7, 4.0}},
		{{6, 6.0}, {10, 8.0}, {2, 6.0}, {6, 6.0}, {5, 6.0}, {0, 5.5}, {6, 6.0}, {3, 5.5}, {1, 5.0}, {0, 4.0}, {0, 3.0}}}
	for i, testSet := range testSets {
		mf := Constructor()
		for j, test := range testSet {
			num := test[0].(int)
			want := test[1].(float64)
			mf.AddNum(num)
			if got := mf.FindMedian(); got != want {
				t.Errorf("(Set %d, Test %d) [%d] FindMedian() = %f, want %f", i, j, num, got, want)
			}
		}

	}
}
