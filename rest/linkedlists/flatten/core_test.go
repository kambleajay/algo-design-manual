package flatten

import (
	"reflect"
	"testing"
)

func link(n1 *Node, n2 *Node) {
	n1.Next = n2
	n2.Prev = n1
}

func input1() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}
	link(n1, n2)
	link(n2, n3)
	link(n3, n4)
	link(n4, n5)
	link(n5, n6)
	link(n7, n8)
	link(n8, n9)
	link(n9, n10)
	link(n11, n12)
	n3.Child = n7
	n8.Child = n11
	return n1
}

func toList(nums ...int) *Node {
	head := &Node{}
	tail := head
	for _, num := range nums {
		node := &Node{Val: num, Prev: tail}
		tail.Next = node
		tail = tail.Next
	}
	head.Next.Prev = nil
	return head.Next
}

func want1() *Node {
	return toList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6)
}

func input2() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	link(n1, n2)
	n1.Child = n3
	return n1
}

func want2() *Node {
	return toList(1, 3, 2)
}

func TestFlatten(t *testing.T) {
	var tests = []struct {
		input *Node
		want  *Node
	}{
		{input1(), want1()},
		{input2(), want2()},
		{nil, nil},
	}
	for _, test := range tests {
		if got := flatten(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("flatten(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
