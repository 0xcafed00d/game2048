package main

import ()

type GameState struct {
	Board     GameBoard
	PrevBoard GameBoard
}

func (gs *GameState) NewGame() {

}

func (gs *GameState) DoMove(dir Direction) int {
	return 0
}

func (gs *GameState) UndoMove() int {
	return 0
}

func (gs *GameState) Tick() {

}
