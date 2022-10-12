package numisls

type Point struct {
	row int
	col int
}

const Land = '1'

type ConnectedComponents struct {
	grid        [][]byte
	seen        map[Point]bool
	components  map[int]bool
	m, n, count int
}

func NewConnectedComponent(grid [][]byte) *ConnectedComponents {
	m := len(grid)
	n := len(grid[0])
	seen := make(map[Point]bool)
	components := make(map[int]bool)
	return &ConnectedComponents{grid: grid, seen: seen, components: components, m: m, n: n}
}

func (cc *ConnectedComponents) adj(p *Point) []*Point {
	north := &Point{p.row - 1, p.col}
	east := &Point{p.row, p.col + 1}
	south := &Point{p.row + 1, p.col}
	west := &Point{p.row, p.col - 1}
	neighbours := []*Point{north, east, south, west}
	var result []*Point
	for _, nb := range neighbours {
		if nb.row >= 0 && nb.row < cc.m && nb.col >= 0 && nb.col < cc.n {
			result = append(result, nb)
		}
	}
	return result
}

func (cc *ConnectedComponents) dfs(p *Point) {
	cc.seen[*p] = true
	cc.components[cc.count] = true
	for _, nb := range cc.adj(p) {
		if !cc.seen[*nb] && cc.grid[nb.row][nb.col] == Land {
			cc.dfs(nb)
		}
	}
}

func (cc *ConnectedComponents) findComponents(points []*Point) {
	for _, p := range points {
		if !cc.seen[*p] {
			cc.dfs(p)
			cc.count++
		}
	}
}

func numIslands(grid [][]byte) int {
	cc := NewConnectedComponent(grid)
	var landPoints []*Point
	for i, row := range grid {
		for j, item := range row {
			if item == Land {
				landPoints = append(landPoints, &Point{i, j})
			}
		}
	}
	cc.findComponents(landPoints)
	return len(cc.components)
}
