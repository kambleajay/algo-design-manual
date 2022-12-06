package treetodoublelist

import (
	"testing"
)

func assert(t *testing.T, got bool, msg string) {
	if !got {
		t.Errorf(msg)
	}
}

func TestTreeToDoublyList1(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n2.Left = n1
	n2.Right = n3

	treeToDoublyList(n2)

	assert(t, n1.Left == n3, "n1's left is not n3")
	assert(t, n1.Right == n2, "n1's right is not n2")
	assert(t, n2.Right == n3, "n2's right is not n3")
	assert(t, n3.Right == n1, "n3's right is not n1")
	assert(t, n3.Left == n2, "n3's left is not n2")
	assert(t, n2.Left == n1, "n2's left is not n1")
}

func TestTreeToDoublyList2(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n4.Left, n4.Right = n2, n5
	n2.Left, n2.Right = n1, n3

	treeToDoublyList(n4)

	assert(t, n1.Left == n5, "n1's left is not n5")
	assert(t, n1.Right == n2, "n1's right is not n2")
	assert(t, n2.Right == n3, "n2's right is not n3")
	assert(t, n3.Right == n4, "n3's right is not n4")
	assert(t, n4.Right == n5, "n4's right is not n5")
	assert(t, n5.Right == n1, "n5's left is not n1")
	assert(t, n5.Left == n4, "n5's left is not n4")
	assert(t, n4.Left == n3, "n4's left is not n3")
	assert(t, n3.Left == n2, "n3's left is not n2")
	assert(t, n2.Left == n1, "n2's left is not n1")
}

func TestTreeToDoublyList3(t *testing.T) {
	treeToDoublyList(nil)
}
