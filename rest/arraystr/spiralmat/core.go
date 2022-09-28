package spiralmat

const (
	Right = "right"
	Down  = "down"
	Left  = "left"
	Up    = "up"
)

type Node struct {
	Dir  string
	Next *Node
}

type Elem struct {
	Row int
	Col int
}

func directions() *Node {
	right := &Node{Dir: Right}
	down := &Node{Dir: Down}
	left := &Node{Dir: Left}
	up := &Node{Dir: Up}
	right.Next = down
	down.Next = left
	left.Next = up
	up.Next = right
	return right
}

func indexesForRightTraversal(matrix [][]int, row int) [][]int {
	var indexes [][]int
	for j := 0; j < len(matrix[row]); j++ {
		indexes = append(indexes, []int{row, j})
	}
	return indexes
}

func indexesForDownTraversal(matrix [][]int, col int) [][]int {
	var indexes [][]int
	for i := 0; i < len(matrix); i++ {
		indexes = append(indexes, []int{i, col})
	}
	return indexes
}

func indexesForLeftTraversal(matrix [][]int, row int) [][]int {
	var indexes [][]int
	for j := len(matrix[row]) - 1; j >= 0; j-- {
		indexes = append(indexes, []int{row, j})
	}
	return indexes
}

func indexesForUpTraversal(matrix [][]int, col int) [][]int {
	var indexes [][]int
	for i := len(matrix) - 1; i >= 0; i-- {
		indexes = append(indexes, []int{i, col})
	}
	return indexes
}

func readElems(matrix [][]int, indexes [][]int, seen map[Elem]bool) ([]int, Elem) {
	var result []int
	lastElemAdded := Elem{}
	for _, index := range indexes {
		row, col := index[0], index[1]
		elem := Elem{row, col}
		if !seen[elem] {
			result = append(result, matrix[row][col])
			seen[elem] = true
			lastElemAdded = elem
		}
	}
	return result, lastElemAdded
}

func spiralOrder(matrix [][]int) []int {
	follow := directions()
	var row, col int
	m, n := len(matrix), len(matrix[0])
	var result, elems []int
	var lastElem Elem
	seen := make(map[Elem]bool)
	for len(result) < m*n {
		if follow.Dir == Right {
			elems, lastElem = readElems(matrix, indexesForRightTraversal(matrix, row), seen)
			result = append(result, elems...)
		} else if follow.Dir == Down {
			elems, lastElem = readElems(matrix, indexesForDownTraversal(matrix, col), seen)
			result = append(result, elems...)
		} else if follow.Dir == Left {
			elems, lastElem = readElems(matrix, indexesForLeftTraversal(matrix, row), seen)
			result = append(result, elems...)
		} else if follow.Dir == Up {
			elems, lastElem = readElems(matrix, indexesForUpTraversal(matrix, col), seen)
			result = append(result, elems...)
		}
		row, col = lastElem.Row, lastElem.Col
		follow = follow.Next
	}
	return result
}
