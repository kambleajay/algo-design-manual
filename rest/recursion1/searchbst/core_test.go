package searchbst

import (
	"math"
	"reflect"
	"testing"
)

func makeTree(nodes []int, k int) *TreeNode {
	if nodes == nil {
		return nil
	}
	if nodes[k] == math.MaxInt64 {
		return nil
	}
	node := &TreeNode{Val: nodes[k]}
	left, right := 2*k, (2*k)+1
	if left < len(nodes) {
		node.Left = makeTree(nodes, left)
	}
	if right < len(nodes) {
		node.Right = makeTree(nodes, right)
	}
	return node
}

func TestSearchBST(t *testing.T) {
	var tests = []struct {
		inputNodes []int
		val        int
		wantNodes  []int
	}{
		{[]int{0, 4, 2, 7, 1, 3, math.MaxInt64, math.MaxInt64}, 2, []int{0, 2, 1, 3}},
		{[]int{0, 4, 2, 7, 1, 3, math.MaxInt64, math.MaxInt64}, 5, nil},
		{nil, 0, nil},
	}
	for _, test := range tests {
		input := makeTree(test.inputNodes, 1)
		want := makeTree(test.wantNodes, 1)
		if got := searchBST(input, test.val); !reflect.DeepEqual(got, want) {
			t.Errorf("searchBST(%v, %d) = %v, want %v", input, test.val, got, want)
		}
	}
}
