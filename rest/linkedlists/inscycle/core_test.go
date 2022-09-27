package inscycle

import (
	"reflect"
	"testing"
)

func input1() *Node {
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1 := &Node{Val: 1}
	n3.Next = n4
	n4.Next = n1
	n1.Next = n3
	return n3
}

func want1() *Node {
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3.Next = n4
	n4.Next = n1
	n1.Next = n2
	n2.Next = n3
	return n3
}

func input2() *Node {
	n := &Node{Val: 1}
	n.Next = n
	return n
}

func want2() *Node {
	n0 := &Node{Val: 0}
	n1 := &Node{Val: 1}
	n1.Next = n0
	n0.Next = n1
	return n1
}

func TestInsert(t *testing.T) {
	var tests = []struct {
		input *Node
		val   int
		want  *Node
	}{
		{input1(), 2, want1()},
		{nil, 1, &Node{Val: 1}},
		{input2(), 0, want2()},
	}
	for _, test := range tests {
		if got := insert(test.input, test.val); !reflect.DeepEqual(got, test.want) {
			t.Errorf("insert(%v, %d) = %v, want %v", test.input, test.val, got, test.want)
		}
	}
}
