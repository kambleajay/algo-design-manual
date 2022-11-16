package connectright2

import (
	"reflect"
	"testing"
)

func tree1() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n7 := &Node{Val: 7}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Right = n7
	return n1
}

func tree1Answer() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n7 := &Node{Val: 7}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Right = n7
	n2.Next = n3
	n4.Next = n5
	n5.Next = n7
	return n1
}

func TestConnect(t *testing.T) {
	var tests = []struct {
		root *Node
		want *Node
	}{
		{tree1(), tree1Answer()},
		{nil, nil},
	}
	for _, test := range tests {
		if got := connect(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("connect(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
