package main

import (
	"fmt"
)

const (
	BoardSize = 4
)

type Cell struct {
	val    int
	locked bool
}

// game board [0][0] is top left
type GameBoard [BoardSize][BoardSize]Cell

func printBoard(gb *GameBoard) {
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			fmt.Print(gb[x][y])
		}
		fmt.Println()
	}
	return
}

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

func RelativePos(x, y int, dir Direction) (newx, newy int, ok bool) {
	dx, dy := DxDy(dir)
	newx = x + dx
	newy = y + dy
	ok = newx >= 0 && newy >= 0 && newx < BoardSize && newy < BoardSize

	if !ok {
		newx, newy = x, y
	}

	return
}

func (gb *GameBoard) MoveCell(x, y int, dir Direction) (moved bool, score int) {

	if gb.CanMove(x, y, dir) {
		newx, newy, _ := RelativePos(x, y, dir)
		if gb[newx][newy].val == 0 {
			gb[newx][newy] = gb[x][y]
		} else {
			gb[newx][newy].val += gb[x][y].val
			gb[newx][newy].locked = true
			score = gb[newx][newy].val
		}
		gb[x][y] = Cell{}
		moved = true
	}

	return
}

func (gb *GameBoard) CanMove(x, y int, dir Direction) bool {
	if gb[x][y].val == 0 {
		return false
	}

	tox, toy, onBoard := RelativePos(x, y, dir)

	fromCell := gb[x][y]
	toCell := gb[tox][toy]

	// can move if new location is on the board & not locked and destinaltion cell is empty
	// or destination is same value as source
	return onBoard && !toCell.locked && (toCell.val == 0 || toCell.val == fromCell.val)
}

func (gb *GameBoard) Reset() {
	*gb = GameBoard{}
}

func (gb *GameBoard) ClearLocks() {
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			gb[x][y].locked = false
		}
	}
}

type Move struct {
	x, y int
}

func (gb *GameBoard) SingleStep(dir Direction) (moves []Move, score int) {

	scanStart, scanStop, scanStep := 0, BoardSize-1, 1
	if dir == Right || dir == Down {
		scanStart, scanStop, scanStep = BoardSize-1, 0, -1
	}

	doMove := func(x, y int, dir Direction) {
		moved, sc := gb.MoveCell(x, y, dir)
		score += sc
		if moved {
			moves = append(moves, Move{x, y})
		}
	}

	if dir == Left || dir == Right {
		for y := 0; y < BoardSize; y++ {
			for x := scanStart; x <= scanStop; x += scanStep {
				doMove(x, y, dir)
			}
		}
	}

	if dir == Up || dir == Down {
		for x := 0; x < BoardSize; x++ {
			for y := scanStart; y <= scanStop; y += scanStep {
				doMove(x, y, dir)
			}
		}
	}

	return
}

func (gb *GameBoard) FindFreeCell() (x int, y int, ok bool) {
	return 0, 0, true
}
