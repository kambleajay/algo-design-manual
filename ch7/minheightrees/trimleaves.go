package minheightrees

func makeAdjList(n int, edges [][]int) []map[int]bool {
	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		v, w := edge[0], edge[1]
		adj[v][w] = true
		adj[w][v] = true
	}
	return adj
}

func initialLeaves(adj []map[int]bool) []int {
	var leaves []int
	for v, adjOfV := range adj {
		if len(adjOfV) == 1 {
			leaves = append(leaves, v)
		}
	}
	return leaves
}

func neighbor(adj []map[int]bool, leaf int) int {
	var neighbor int
	for k, _ := range adj[leaf] {
		neighbor = k
	}
	return neighbor
}

func findMinHeightTrees2(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	adj := makeAdjList(n, edges)
	leaves := initialLeaves(adj)
	remainingNodes := n
	var newLeaves []int
	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		for _, leaf := range leaves {
			neighbor := neighbor(adj, leaf)
			delete(adj[neighbor], leaf)
			if len(adj[neighbor]) == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
		}
		leaves = newLeaves
		newLeaves = nil
	}
	return leaves
}
