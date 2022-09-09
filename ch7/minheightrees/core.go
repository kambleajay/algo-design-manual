package minheightrees

import (
	"math"
)

//-- Graph --

type Graph struct {
	v   int
	adj [][]int
}

func (g *Graph) addEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

func NewGraph(v int, edges [][]int) *Graph {
	g := Graph{v: v}
	adj := make([][]int, v)
	g.adj = adj
	for _, edge := range edges {
		v, w := edge[0], edge[1]
		g.addEdge(v, w)
	}
	return &g
}

//-- Queue --

type Node struct {
	item int
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
}

func (q *Queue) isEmpty() bool {
	return q.first == nil
}

func (q *Queue) enqueue(item int) {
	oldLast := q.last
	q.last = &Node{item: item}
	if q.isEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *Queue) dequeue() int {
	item := q.first.item
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	return item
}

//-- BFS --

type BreadthFirstPaths struct {
	marked      []bool
	distTo      []int
	longestPath int
}

func NewBFSPaths(v int) *BreadthFirstPaths {
	marked := make([]bool, v)
	distTo := make([]int, v)
	longestPath := math.MinInt64
	bfsp := BreadthFirstPaths{marked, distTo, longestPath}
	return &bfsp
}

func (bf *BreadthFirstPaths) bfs(g *Graph, s int) {
	q := Queue{}
	q.enqueue(s)
	bf.marked[s] = true
	for !q.isEmpty() {
		v := q.dequeue()
		for _, w := range g.adj[v] {
			if !bf.marked[w] {
				bf.marked[w] = true
				bf.distTo[w] = bf.distTo[v] + 1
				if bf.distTo[w] > bf.longestPath {
					bf.longestPath = bf.distTo[w]
				}
				q.enqueue(w)
			}
		}
	}
}

func findMinHeightTrees1(n int, edges [][]int) []int {
	g := NewGraph(n, edges)
	var bfp *BreadthFirstPaths
	minLongestPath := math.MaxInt64
	var answer []int
	for s := 0; s < n; s++ {
		bfp = NewBFSPaths(n)
		bfp.bfs(g, s)
		if bfp.longestPath < minLongestPath {
			minLongestPath = bfp.longestPath
			answer = []int{s}
		} else if bfp.longestPath == minLongestPath {
			answer = append(answer, s)
		}
	}
	return answer
}

func findMinHeightTrees(n int, edges [][]int) []int {
	return findMinHeightTrees2(n, edges)
}
