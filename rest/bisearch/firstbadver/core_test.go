package firstbadver

import "testing"

var badVersionAt4 = func(x int) bool {
	return x == 4
}

var badVersionAt1 = func(x int) bool {
	return x == 1
}

func TestFirstBadVersion(t *testing.T) {
	isBadVersion = badVersionAt4
	if got := firstBadVersion(5); got != 4 {
		t.Errorf("firstBadVersion(%d) = %d, want %d", 5, got, 4)
	}
	isBadVersion = badVersionAt1
	if got := firstBadVersion(1); got != 1 {
		t.Errorf("firstBadVersion(%d) = %d, want %d", 1, got, 1)
	}
}
