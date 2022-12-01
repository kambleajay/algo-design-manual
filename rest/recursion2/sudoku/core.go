package sudoku

import "fmt"

const (
	Empty     = '.'
	Dimension = 9
)

type Cell struct {
	row, col int
}

type Solution struct {
	board      [][]byte
	moves      []*Cell
	a          []int
	steps      int
	finished   bool
	emptyCount int
}

func (s *Solution) resize(capacity int) {
	tmp := make([]int, capacity)
	copy(tmp, s.a)
	s.a = tmp
}

func (s *Solution) exclusions(cell Cell) map[int]bool {
	exclusions := make(map[int]bool)
	for i := 0; i < 9; i++ {
		val := s.board[cell.row][i]
		if val != Empty {
			exclusions[byteToInt(val)] = true
		}
	}
	for j := 0; j < 9; j++ {
		val := s.board[j][cell.col]
		if val != Empty {
			exclusions[byteToInt(val)] = true
		}
	}
	boxRow, boxCol := cell.row/3, cell.col/3
	for i := boxRow * 3; i < (boxRow+1)*3; i++ {
		for j := boxCol * 3; j < (boxCol+1)*3; j++ {
			val := s.board[i][j]
			if val != Empty {
				exclusions[byteToInt(val)] = true
			}
		}
	}
	return exclusions
}

func (s *Solution) possibleValues(cell Cell) []int {
	exc := s.exclusions(cell)
	var possibleValues []int
	for i := 1; i <= 9; i++ {
		if !exc[i] {
			possibleValues = append(possibleValues, i)
		}
	}
	return possibleValues
}

func NewSolution(board [][]byte) Solution {
	m, n := len(board), len(board[0])
	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == Empty {
				count++
			}
		}
	}
	a := make([]int, 10)
	moves := make([]*Cell, 82)
	return Solution{board: board, a: a, emptyCount: count, moves: moves}
}

func (s *Solution) isSolution() bool {
	return s.emptyCount == 0
}

func (s *Solution) processSolution() {
	s.steps++
	s.finished = true
}

func byteToInt(b byte) int {
	return int(b - '0')
}

func intToByte(n int) byte {
	return byte(n + '0')
}

func (s *Solution) less(cell1, cell2 *Cell) bool {
	pvs1 := s.possibleValues(*cell1)
	pvs2 := s.possibleValues(*cell2)
	return len(pvs1) < len(pvs2)
}

func (s *Solution) partition(cells []*Cell, lo, hi int) int {
	firstHigh, p := lo, hi
	for i := lo; i < hi; i++ {
		if s.less(cells[i], cells[p]) {
			cells[i], cells[firstHigh] = cells[firstHigh], cells[i]
			firstHigh++
		}
	}
	cells[firstHigh], cells[p] = cells[p], cells[firstHigh]
	return firstHigh
}

func (s *Solution) qselect(cells []*Cell, k int) *Cell {
	lo, hi := 0, len(cells)-1
	for lo < hi {
		p := s.partition(cells, lo, hi)
		if k < p {
			hi = p - 1
		} else if k > p {
			lo = p + 1
		} else {
			return cells[k]
		}
	}
	return cells[k]
}

func (s *Solution) min(cells []*Cell) *Cell {
	return s.qselect(cells, 0)
}

func (s *Solution) nextEmptyCell() *Cell {
	var cells []*Cell
	for i := 0; i < Dimension; i++ {
		for j := 0; j < Dimension; j++ {
			if s.board[i][j] == Empty {
				cell := &Cell{i, j}
				cells = append(cells, cell)
			}
		}
	}
	return s.min(cells)
}

func (s *Solution) constructCandidates(k int) []int {
	cell := s.nextEmptyCell()
	if cell == nil {
		panic(fmt.Sprintf("k = %d, no empty cells", k))
	}
	s.moves[k] = cell
	return s.possibleValues(*cell)
}

func (s *Solution) makeMove(k int) {
	cell := s.moves[k]
	s.board[cell.row][cell.col] = intToByte(s.a[k])
	s.emptyCount--
}

func (s *Solution) unmakeMove(k int) {
	if s.finished {
		return
	}
	cell := s.moves[k]
	s.board[cell.row][cell.col] = Empty
	s.emptyCount++
}

func (s *Solution) backtrack(k int) {
	if s.isSolution() {
		s.processSolution()
	} else {
		k++
		for _, candidate := range s.constructCandidates(k) {
			if k == len(s.a) {
				s.resize(2 * len(s.a))
			}
			s.a[k] = candidate
			s.makeMove(k)
			s.backtrack(k)
			s.unmakeMove(k)
			if s.finished {
				return
			}
		}
	}
}

func solveSudoku(board [][]byte) {
	s := NewSolution(board)
	s.backtrack(0)
}
