package reconstruct_queue

import (
	"fmt"
	"testing"
)

func TestListInsert(t *testing.T) {
	var l List
	l.InsertAt(0, []int{7, 0})
	l.InsertAt(1, []int{7, 1})
	l.InsertAt(1, []int{6, 1})
	l.InsertAt(0, []int{5, 0})
	l.InsertAt(2, []int{5, 2})
	l.InsertAt(4, []int{4, 4})
	node := l.first
	for node != nil {
		fmt.Printf("%v ", node)
		node = node.next
	}
}
