package clonegraph

import (
	"reflect"
	"testing"
)

func graph1() *Node {
	g1 := &Node{Val: 1}
	g2 := &Node{Val: 2}
	g3 := &Node{Val: 3}
	g4 := &Node{Val: 4}
	g1.Neighbors = []*Node{g2, g4}
	g2.Neighbors = []*Node{g1, g3}
	g3.Neighbors = []*Node{g2, g4}
	g4.Neighbors = []*Node{g1, g3}
	return g1
}

func graph2() *Node {
	return &Node{Val: 1}
}

func TestCloneGraph(t *testing.T) {
	var graphs = []*Node{graph1(), graph2(), nil}
	for _, graph := range graphs {
		if got := cloneGraph(graph); !reflect.DeepEqual(got, graph) {
			t.Errorf("cloneGraph(%v) = %v, want %v", graph, got, graph)
		}
	}

}
