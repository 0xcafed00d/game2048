package main

const (
	BoardSizeX = 4
	BoardSizeY = 4
)

type Cell struct {
	val    int
	locked bool
}

type GameBoard [BoardSizeX][BoardSizeY]Cell

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func DxDy(d Direction) (dx, dy int) {
	switch d {
	case Up:
		dx, dy = 0, -1
	case Down:
		dx, dy = 0, 1
	case Left:
		dx, dy = -1, 0
	case Right:
		dx, dy = 1, 0
	}
	return
}

type GameState struct {
	Board     GameBoard
	PrevBoard GameBoard
}

func (gb *GameBoard) RelativePos(x, y int, dir Direction) (newx, newy int, ok bool) {
	dx, dy := DxDy(dir)
	newx = x + dx
	newy = y + dy
	ok = newx >= 0 && newy >= 0 && newx < BoardSizeX && newy < BoardSizeY

	return
}

func (gb *GameBoard) MoveCell(x, y int, dir Direction) {
	if gb.CanMove(x, y, dir) {
		dx, dy, _ := gb.RelativePos(x, y, dir)
		if gb[x+dx][y+dy].val == 0 {
			gb[x+dx][y+dy] = gb[x][y]
		} else {
			gb[x+dx][y+dy].val += gb[x][y].val
			gb[x+dx][y+dy].locked = true
		}
		gb[x][y] = Cell{}
	}
}

func (gb *GameBoard) CanMove(x, y int, dir Direction) bool {
	dx, dy, onBoard := gb.RelativePos(x, y, dir)

	fromCell := gb[x][y]
	toCell := gb[x+dx][y+dy]

	// can move if new location is on the board & not locked and destinaltion cell is empty
	// or destination is same value as source
	return onBoard && !toCell.locked && (toCell.val == 0 || toCell.val == fromCell.val)
}

func (gb *GameBoard) FindFreeCell() (x int, y int, ok bool) {
	return 0, 0, true
}
