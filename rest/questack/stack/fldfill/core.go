package fldfill

type Point struct {
	x int
	y int
}

func between(n, min, max int) bool {
	return n >= min && n < max
}

func adj(p Point, m, n int) []Point {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var result []Point
	for _, dir := range directions {
		nb := Point{p.x + dir[0], p.y + dir[1]}
		if between(nb.x, 0, m) && between(nb.y, 0, n) {
			result = append(result, nb)
		}
	}
	return result
}

func dfs(x, y int, image [][]int, sourceColor, color int, seen map[Point]bool) {
	p := Point{x, y}
	seen[p] = true
	pointColor := image[x][y]
	if pointColor == sourceColor {
		image[x][y] = color
	} else {
		image[x][y] = pointColor
	}
	for _, nb := range adj(p, len(image), len(image[0])) {
		if !seen[nb] && pointColor == sourceColor {
			dfs(nb.x, nb.y, image, sourceColor, color, seen)
		}
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	seen := make(map[Point]bool)
	dfs(sr, sc, image, image[sr][sc], color, seen)
	return image
}
