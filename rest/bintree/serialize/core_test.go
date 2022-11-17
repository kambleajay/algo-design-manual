package serialize

import (
	"math"
	"reflect"
	"testing"
)

func makeTreeTest(nodes []int, k int) *TreeNode {
	if nodes == nil {
		return nil
	}
	if nodes[k] == math.MaxInt64 {
		return nil
	}
	node := &TreeNode{Val: nodes[k]}
	left, right := 2*k, (2*k)+1
	if left < len(nodes) {
		node.Left = makeTreeTest(nodes, left)
	}
	if right < len(nodes) {
		node.Right = makeTreeTest(nodes, right)
	}
	return node
}

func TestSerialization(t *testing.T) {
	var tests = []struct {
		nodes []int
		want  string
	}{
		{nil, "#"},
		{[]int{0, 1, 2, 3, math.MaxInt64, math.MaxInt64, 4, 5}, "1,2,#,#,3,4,#,#,5,#,#"},
		{[]int{0, 3, 5, 1, 6, 2, 0, 8, math.MaxInt64, math.MaxInt64, 7, 4}, "3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#"},
		{[]int{0, 1, 2, 3, math.MaxInt64, math.MaxInt64, 4, 5, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 6, 7, math.MaxInt64, math.MaxInt64}, "1,2,#,#,3,4,6,#,#,7,#,#,5,#,#"},
	}
	for _, test := range tests {
		root := makeTreeTest(test.nodes, 1)
		serializer := Constructor()
		if got := serializer.serialize(root); got != test.want {
			t.Errorf("serialize(%v) = %s, want %s", root, got, test.want)
		}

	}
}

func TestDeserialization(t *testing.T) {
	var tests = []struct {
		data  string
		nodes []int
	}{
		{"#", nil},
		{"1,2,#,#,3,4,#,#,5,#,#", []int{0, 1, 2, 3, math.MaxInt64, math.MaxInt64, 4, 5}},
		{"3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#", []int{0, 3, 5, 1, 6, 2, 0, 8, math.MaxInt64, math.MaxInt64, 7, 4}},
		{"1,2,#,#,3,4,6,#,#,7,#,#,5,#,#", []int{0, 1, 2, 3, math.MaxInt64, math.MaxInt64, 4, 5, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 6, 7, math.MaxInt64, math.MaxInt64}},
	}
	for _, test := range tests {
		want := makeTreeTest(test.nodes, 1)
		deserializer := Constructor()
		if got := deserializer.deserialize(test.data); !reflect.DeepEqual(got, want) {
			t.Errorf("deserialize(%s) = %v, want %v", test.data, got, want)
		}

	}
}
