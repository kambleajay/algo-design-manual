package reconstruct_queue

import (
	"fmt"
	"sort"
)

type Node struct {
	item []int
	next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.item)
}

type List struct {
	first, last *Node
	size        int
}

func (list *List) IsEmpty() bool {
	return list.first == nil
}

func (list *List) ToArray() [][]int {
	res := make([][]int, list.size)
	i := 0
	for x := list.first; x != nil; x = x.next {
		res[i] = x.item
		i++
	}
	return res
}

func (list *List) LinkFirst(item []int) {
	oldFirst := list.first
	node := Node{item: item, next: oldFirst}
	list.first = &node
	if oldFirst == nil {
		list.last = &node
	}
	list.size++
}

func (list *List) LinkLast(item []int) {
	oldLast := list.last
	node := &Node{item: item}
	list.last = node
	if list.IsEmpty() {
		list.first = node
	} else {
		oldLast.next = node
	}
	list.size++
}

func (list *List) LinkBefore(item []int, pred *Node) {
	node := &Node{item: item, next: pred.next}
	if pred == nil {
		list.first = node
	} else {
		pred.next = node
	}
	list.size++
}

func (list *List) NodeBefore(index int) *Node {
	x := list.first
	for i := 0; i < index-1; i++ {
		x = x.next
	}
	return x
}

func (list *List) InsertAt(index int, item []int) {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index is out of range: %d, size: %d\n", index, list.size))
	}
	if index == list.size {
		list.LinkLast(item)
	} else if index == 0 {
		list.LinkFirst(item)
	} else {
		list.LinkBefore(item, list.NodeBefore(index))
	}
}

func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		height1, kIndex1 := people[i][0], people[i][1]
		height2, kIndex2 := people[j][0], people[j][1]
		if height1 != height2 {
			return height1 > height2
		}
		return kIndex1 <= kIndex2
	})
	var list List
	for _, person := range people {
		kIndex := person[1]
		list.InsertAt(kIndex, person)
	}
	return list.ToArray()
}
