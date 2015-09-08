package main

import (
	"github.com/simulatedsimian/assert"
	"testing"
)

func pack(vals ...interface{}) []interface{} {
	return vals
}

func TestNavigation(t *testing.T) {

	assert.Equal(t, pack(RelativePos(1, 1, Left)), pack(0, 1, true))
	assert.Equal(t, pack(RelativePos(0, 1, Left)), pack(0, 1, false))

	assert.Equal(t, pack(RelativePos(1, 1, Right)), pack(2, 1, true))
	assert.Equal(t, pack(RelativePos(BoardSizeX-1, 1, Right)), pack(BoardSizeX-1, 1, false))

	assert.Equal(t, pack(RelativePos(1, 1, Up)), pack(1, 0, true))
	assert.Equal(t, pack(RelativePos(1, 0, Up)), pack(1, 0, false))

	assert.Equal(t, pack(RelativePos(1, 1, Down)), pack(1, 2, true))
	assert.Equal(t, pack(RelativePos(1, BoardSizeY-1, Down)), pack(1, BoardSizeY-1, false))
}

func TestCanMove(t *testing.T) {

	var gb GameBoard

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
