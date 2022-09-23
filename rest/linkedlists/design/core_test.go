package design

import (
	"algo/rest/linkedlists/design/dll"
	"algo/rest/linkedlists/design/sll"
	"testing"
)

func singlyLinkedList() *sll.MyLinkedList {
	list := sll.Constructor()
	return &list
}

func doublyLinkedList() *dll.MyLinkedList {
	list := dll.Constructor()
	return &list
}

func makeList() List {
	return doublyLinkedList()
}

func TestAddAtHead(t *testing.T) {
	var tests = []struct {
		elems []int
		wants [][]int
	}{
		{[]int{3, 2, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, -1}, {-2, -1}}},
	}
	for _, test := range tests {
		list := makeList()
		for _, elem := range test.elems {
			list.AddAtHead(elem)
		}
		for _, want := range test.wants {
			index, value := want[0], want[1]
			if v := list.Get(index); v != value {
				t.Errorf("Get(%d) = %d, want %d", index, v, value)
			}
		}
	}
}

func TestAddAtTail(t *testing.T) {
	var tests = []struct {
		elems []int
		wants [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, -1}}},
	}
	for _, test := range tests {
		list := makeList()
		for _, elem := range test.elems {
			list.AddAtTail(elem)
		}
		for _, want := range test.wants {
			index, value := want[0], want[1]
			if v := list.Get(index); v != value {
				t.Errorf("Get(%d) = %d, want %d", index, v, value)
			}
		}
	}
}

func TestAddAtIndex(t *testing.T) {
	var tests = []struct {
		elems [][]int
		wants [][]int
	}{
		{[][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, -1}}},
		{[][]int{{0, 3}, {-1, 2}, {2, 1}, {3, 4}, {10, 10}}, [][]int{{0, 2}, {1, 3}, {2, 1}, {3, 4}, {10, -1}}},
	}
	for _, test := range tests {
		list := makeList()
		for _, elem := range test.elems {
			i, e := elem[0], elem[1]
			list.AddAtIndex(i, e)
		}
		for _, want := range test.wants {
			index, value := want[0], want[1]
			if v := list.Get(index); v != value {
				t.Errorf("Get(%d) = %d, want %d", index, v, value)
			}
		}
	}
}

func TestDeleteAtIndex(t *testing.T) {
	var tests = []struct {
		elemsToInsert   []int
		indexesToDelete []int
		wants           [][]int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2}, [][]int{{0, 1}, {1, 3}, {2, 5}, {3, -1}}},
	}
	for _, test := range tests {
		list := makeList()
		for _, elem := range test.elemsToInsert {
			list.AddAtHead(elem)
		}
		for _, index := range test.indexesToDelete {
			list.DeleteAtIndex(index)
		}
		for _, want := range test.wants {
			index, value := want[0], want[1]
			if v := list.Get(index); v != value {
				t.Errorf("Get(%d) = %d, want %d", index, v, value)
			}
		}
	}
}

func TestList(t *testing.T) {
	list := makeList()
	//["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
	//[[],            [2],         [1],           [2],        [7],         [3],       [2],        [5],        [5],        [5],    [6],          [4]]
	list.AddAtHead(2)
	list.DeleteAtIndex(1)
	list.AddAtHead(2)
	list.AddAtHead(7)
	list.AddAtHead(3)
	list.AddAtHead(2)
	list.AddAtHead(5)
	list.AddAtTail(5)
	if v := list.Get(5); v != 2 {
		t.Errorf("Get(5) = %d, want %d", v, 2)
	}
	list.DeleteAtIndex(6)
	list.DeleteAtIndex(4)
	if v := list.Get(0); v != 5 {
		t.Errorf("Get(0) = %d, want %d", v, 5)
	}
}
