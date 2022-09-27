package copyrandlist

import (
	"reflect"
	"testing"
)

func list1() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2
	return n1
}

func list2() *Node {
	n7 := &Node{Val: 7}
	n13 := &Node{Val: 13}
	n11 := &Node{Val: 11}
	n10 := &Node{Val: 10}
	n1 := &Node{Val: 1}
	n7.Next = n13
	n13.Next = n11
	n11.Next = n10
	n10.Next = n1
	n13.Random = n7
	n11.Random = n1
	n10.Random = n11
	n1.Random = n7
	return n7
}

func list3() *Node {
	n31 := &Node{Val: 3}
	n32 := &Node{Val: 3}
	n33 := &Node{Val: 3}
	n31.Next = n32
	n32.Next = n33
	n32.Random = n31
	return n31
}

func TestCopyRandomList(t *testing.T) {
	var tests = []*Node{list1(), list2(), list3()}
	for _, test := range tests {
		got := copyRandomList(test)
		if !reflect.DeepEqual(got, test) || got == test {
			t.Errorf("copyRandomList(%v) = %v, want %v", test, got, test)
		}
	}

}
