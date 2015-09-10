package main

import (
	//	"fmt"
	"github.com/simulatedsimian/assert"
	"testing"
)

// pack a number of values into a slice containing those values
func pack(vals ...interface{}) []interface{} {
	return vals
}

func TestNavigation(t *testing.T) {

	assert.Equal(t, pack(RelativePos(1, 1, Left)), pack(0, 1, true))
	assert.Equal(t, pack(RelativePos(0, 1, Left)), pack(0, 1, false))

	assert.Equal(t, pack(RelativePos(1, 1, Right)), pack(2, 1, true))
	assert.Equal(t, pack(RelativePos(BoardSize-1, 1, Right)), pack(BoardSize-1, 1, false))

	assert.Equal(t, pack(RelativePos(1, 1, Up)), pack(1, 0, true))
	assert.Equal(t, pack(RelativePos(1, 0, Up)), pack(1, 0, false))

	assert.Equal(t, pack(RelativePos(1, 1, Down)), pack(1, 2, true))
	assert.Equal(t, pack(RelativePos(1, BoardSize-1, Down)), pack(1, BoardSize-1, false))
}

func TestCanMove(t *testing.T) {

	var gb GameBoard

	assert.Equal(t, gb.CanMove(1, 1, Left), false)

	gb[1][1].val = 2

	assert.Equal(t, gb.CanMove(1, 1, Up), true)
	assert.Equal(t, gb.CanMove(1, 1, Down), true)
	assert.Equal(t, gb.CanMove(1, 1, Left), true)
	assert.Equal(t, gb.CanMove(1, 1, Right), true)

	gb[2][1].val = 4
	assert.Equal(t, gb.CanMove(1, 1, Right), false)

	gb[0][1].val = 2
	assert.Equal(t, gb.CanMove(1, 1, Left), true)

	gb[0][1].locked = true
	assert.Equal(t, gb.CanMove(1, 1, Left), false)
}

var (
	_0  = Cell{0, false}
	_2  = Cell{2, false}
	_4  = Cell{4, false}
	_8  = Cell{8, false}
	_0l = Cell{0, true}
	_2l = Cell{2, true}
	_4l = Cell{4, true}
	_8l = Cell{8, true}
)

func TestMove(t *testing.T) {
	gb := GameBoard{}

	assert.Equal(t, pack(gb.MoveCell(1, 1, Up)), pack(false, 0))

	gb[3][3] = _2
	assert.Equal(t, pack(gb.MoveCell(3, 3, Left)), pack(true, 0))
	assert.Equal(t, gb[3][3], _0)
	assert.Equal(t, gb[2][3], _2)

	gb.Reset()

	gb[0][0] = _2
	gb[1][0] = _2

	assert.Equal(t, pack(gb.MoveCell(1, 0, Up)), pack(false, 0))
	assert.Equal(t, pack(gb.MoveCell(1, 0, Left)), pack(true, 4))
	assert.Equal(t, gb[1][0], _0)
	assert.Equal(t, gb[0][0], _4l)

	gb.ClearLocks()
	assert.Equal(t, gb[0][0], _4)
}

// need this otherwise initialising gameboard directly
// will transpose axis
func makeBoard(cells ...Cell) (gb GameBoard) {
	i := 0
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			gb[x][y] = cells[i]
			i++
		}
	}
	return
}

func TestStep(t *testing.T) {

	gb := makeBoard(
		_0, _0, _2, _2,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _4)

	assert.Equal(t, gb[2][0], _2)
	assert.Equal(t, gb[1][0], _0)

	assert.Equal(t, pack(gb.SingleStep(Left)), pack([]Move{{2, 0}, {3, 0}, {3, 3}}, 0))

	assert.Equal(t, gb, makeBoard(
		_0, _2, _2, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _4, _0))

	assert.Equal(t, pack(gb.SingleStep(Left)), pack([]Move{{1, 0}, {2, 0}, {2, 3}}, 0))

	assert.Equal(t, gb, makeBoard(
		_2, _2, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _4, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Left)), pack([]Move{{1, 0}, {1, 3}}, 4))

	assert.Equal(t, gb, makeBoard(
		_4l, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_4, _0, _0, _0))

	gb.ClearLocks()

	assert.Equal(t, gb, makeBoard(
		_4, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_4, _0, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Down)), pack([]Move{{0, 0}}, 0))

	assert.Equal(t, gb, makeBoard(
		_0, _0, _0, _0,
		_4, _0, _0, _0,
		_0, _0, _0, _0,
		_4, _0, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Down)), pack([]Move{{0, 1}}, 0))

	assert.Equal(t, gb, makeBoard(
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_4, _0, _0, _0,
		_4, _0, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Down)), pack([]Move{{0, 2}}, 8))

	assert.Equal(t, gb, makeBoard(
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0,
		_8l, _0, _0, _0))
}

func TestStep2(t *testing.T) {

	gb := makeBoard(
		_0, _0, _0, _0,
		_0, _2, _2, _0,
		_0, _2, _2, _0,
		_0, _0, _0, _0)

	assert.Equal(t, pack(gb.SingleStep(Up)),
		pack([]Move{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 0))

	assert.Equal(t, gb, makeBoard(
		_0, _2, _2, _0,
		_0, _2, _2, _0,
		_0, _0, _0, _0,
		_0, _0, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Right)),
		pack([]Move{{2, 0}, {1, 0}, {2, 1}, {1, 1}}, 0))

	assert.Equal(t, gb, makeBoard(
		_0, _0, _2, _2,
		_0, _0, _2, _2,
		_0, _0, _0, _0,
		_0, _0, _0, _0))

	assert.Equal(t, pack(gb.SingleStep(Right)),
		pack([]Move{{2, 0}, {2, 1}}, 8))

	assert.Equal(t, gb, makeBoard(
		_0, _0, _0, _4l,
		_0, _0, _0, _4l,
		_0, _0, _0, _0,
		_0, _0, _0, _0))

}
