package robotroomcleaner

const (
	North = iota
	East
	South
	West
)

type Robot struct {
	rooms   [][]int
	curr    Cell
	dir     int
	cleaned map[Cell]bool
	m, n    int
}

func NewRobot(rooms [][]int, x, y int) Robot {
	p := Cell{x, y}
	m := make(map[Cell]bool)
	return Robot{rooms: rooms, curr: p, cleaned: m, m: len(rooms), n: len(rooms[0])}
}

// Returns true if the cell in front is open and robot directions into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool {
	p := Cell{robot.curr.row, robot.curr.col}
	switch robot.dir {
	case North:
		p.row--
	case East:
		p.col++
	case South:
		p.row++
	case West:
		p.col--
	}
	if p.row < 0 || p.row >= robot.m {
		return false
	}
	if p.col < 0 || p.col >= robot.n {
		return false
	}
	if robot.rooms[p.row][p.col] == Wall {
		return false
	}
	robot.curr = p
	return true
}

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft() {
	switch robot.dir {
	case North:
		robot.dir = West
	case East:
		robot.dir = North
	case South:
		robot.dir = East
	case West:
		robot.dir = South
	}
}
func (robot *Robot) TurnRight() {
	switch robot.dir {
	case North:
		robot.dir = East
	case East:
		robot.dir = South
	case South:
		robot.dir = West
	case West:
		robot.dir = North
	}
}

// Clean the current cell.
func (robot *Robot) Clean() {
	robot.cleaned[robot.curr] = true
}

const Wall = 0

type Cell struct {
	row, col int
}

var directions = []int{North, East, South, West}

type Solution struct {
	robot   *Robot
	seen    map[Cell]bool
	a       []int
	currDir int
}

func NewSolution(robot *Robot) Solution {
	seen := make(map[Cell]bool)
	a := make([]int, 10)
	return Solution{seen: seen, a: a, robot: robot, currDir: North}
}

func (s *Solution) resize(capacity int) {
	tmp := make([]int, capacity)
	copy(tmp, s.a)
	s.a = tmp
}

const (
	NoTurn = iota
	Right
	TwoRights
	Left
)

var fromDirToToDirMoves = map[int]map[int]int{
	North: {North: NoTurn, East: Right, South: TwoRights, West: Left},
	East:  {North: Left, East: NoTurn, South: Right, West: TwoRights},
	South: {North: TwoRights, East: Left, South: NoTurn, West: Right},
	West:  {North: Right, East: TwoRights, South: Left, West: NoTurn},
}

var oppositeDirs = map[int]int{North: South, East: West, South: North, West: East}

func (s *Solution) TurnTwoRights() {
	s.robot.TurnRight()
	s.robot.TurnRight()
}

func (s *Solution) turn(from, to int) {
	move := fromDirToToDirMoves[from][to]
	switch move {
	case NoTurn:
	case Right:
		s.robot.TurnRight()
	case Left:
		s.robot.TurnLeft()
	case TwoRights:
		s.TurnTwoRights()
	}
	s.currDir = to
}

func (s *Solution) makeMove(dir int) bool {
	s.turn(s.currDir, dir)
	return s.robot.Move()
}

func (s *Solution) unmakeMove(dir int) {
	toDir := oppositeDirs[dir]
	s.makeMove(toDir)
}

func cellAt(currCell Cell, dir int) Cell {
	result := Cell{currCell.row, currCell.col}
	switch dir {
	case North:
		result.row--
	case East:
		result.col++
	case South:
		result.row++
	case West:
		result.col--
	}
	return result
}

func (s *Solution) backtrack(k int, currCell Cell) {
	k++
	for _, dir := range directions {
		toCell := cellAt(currCell, dir)
		if !s.seen[toCell] {
			if k == len(s.a) {
				s.resize(2 * len(s.a))
			}
			s.a[k] = dir
			validMove := s.makeMove(dir)
			if validMove {
				s.robot.Clean()
				s.seen[toCell] = true
				s.backtrack(k, toCell)
				s.unmakeMove(dir)
			}
		}
	}
}

func cleanRoom(robot *Robot) {
	solution := NewSolution(robot)
	source := Cell{0, 0}
	solution.seen[source] = true
	solution.robot.Clean()
	solution.backtrack(0, source)
}
