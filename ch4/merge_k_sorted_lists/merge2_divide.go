package merge_k_sorted_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}
	return head.Next
}

func mergeKLists3(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	N := len(lists)
	interval := 1
	for interval < N {
		for i := 0; i < (N - interval); i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}
