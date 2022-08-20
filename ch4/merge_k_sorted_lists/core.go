package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeEmpty(lists []*ListNode) []*ListNode {
	var res []*ListNode
	for _, l := range lists {
		if l != nil {
			res = append(res, l)
		}
	}
	return res
}

func atLeastOneNotNil(lists []*ListNode) bool {
	removeEmpty(lists)
	for _, l := range lists {
		if l != nil {
			return true
		}
	}
	return false
}

func findIndexOfMin(lists []*ListNode) int {
	var minIndex *int
	for i, l := range lists {
		if minIndex == nil && l != nil {
			minIndex = new(int)
			*minIndex = i
		} else if l != nil && l.Val < lists[*minIndex].Val {
			*minIndex = i
		}
	}
	return *minIndex
}

func mergeKLists1(lists []*ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for lists = removeEmpty(lists); atLeastOneNotNil(lists); lists = removeEmpty(lists) {
		indexOfMin := findIndexOfMin(lists)
		tail.Next = lists[indexOfMin]
		lists[indexOfMin] = lists[indexOfMin].Next
		tail = tail.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists3(lists)
}
