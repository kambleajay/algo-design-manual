package lowcommonancestor

import (
	"reflect"
	"testing"
)

type Test struct {
	root, p, q, want *TreeNode
}

func makeNodes(size int) []*TreeNode {
	nodes := make([]*TreeNode, 9)
	for i := 0; i < size; i++ {
		nodes[i] = &TreeNode{Val: i}
	}
	return nodes
}

func test1() Test {
	nodes := makeNodes(9)
	nodes[3].Left, nodes[3].Right = nodes[5], nodes[1]
	nodes[5].Left, nodes[5].Right = nodes[6], nodes[2]
	nodes[2].Left, nodes[2].Right = nodes[7], nodes[4]
	nodes[1].Left, nodes[1].Right = nodes[0], nodes[8]
	return Test{nodes[3], nodes[5], nodes[1], nodes[3]}
}

func test2() Test {
	nodes := makeNodes(9)
	nodes[3].Left, nodes[3].Right = nodes[5], nodes[1]
	nodes[5].Left, nodes[5].Right = nodes[6], nodes[2]
	nodes[2].Left, nodes[2].Right = nodes[7], nodes[4]
	nodes[1].Left, nodes[1].Right = nodes[0], nodes[8]
	return Test{nodes[3], nodes[5], nodes[4], nodes[5]}
}

func test3() Test {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n1.Right = n2
	return Test{n1, n1, n2, n1}
}

func TestLowestCommonAncestor(t *testing.T) {
	var tests = []Test{test1(), test2(), test3()}
	for _, test := range tests {
		if got := lowestCommonAncestor(test.root, test.p, test.q); !reflect.DeepEqual(got, test.want) {
			t.Errorf("lowestCommonAncestor(%v, %d, %d) = %d, want %d", test.root, test.p.Val, test.q.Val, got.Val, test.want.Val)
		}
	}
}
