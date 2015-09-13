package main

import (
	"github.com/simulatedsimian/game2048/glib"
)

type BoardDrawer interface {
	DrawScores(scr, hi int)
	DrawBoardNow(gb *GameBoard)
	StartSlideTiles(tiles []Move, dir Direction)
	DoneSlideTiles() bool
}

type GameInfo struct {
	Score   int
	HiScore int
	Board   GameBoard
}

type GameState struct {
	glib.StateMachine

	Drawer   BoardDrawer
	Current  GameInfo
	Previous GameInfo
	Move     Direction
}

const (
	StateGameOver = iota
	StateReadyForMove
	StateMoveInProgress
)

func MakeGameState() *GameState {
	gs := GameState{}

	gs.AddState(StateGameOver, glib.State{})
	gs.AddState(StateReadyForMove, glib.State{})
	gs.AddState(StateMoveInProgress, glib.State{})

	gs.NewGame()

	return &gs
}

func (gs *GameState) NewGame() {
	gs.Current.Board.Reset()
	gs.Current.Score = 0
	gs.Previous = gs.Current

	gs.Drawer.DrawScores(gs.Current.Score, gs.Current.HiScore)
	gs.Drawer.DrawBoardNow(&gs.Current.Board)

	gs.Goto(StateReadyForMove)
}

func (gs *GameState) ReadyForMove() bool {
	id, ok := gs.CurrentId()
	return ok && id == StateReadyForMove
}

func (gs *GameState) DoMove(dir Direction) {
	if gs.ReadyForMove() {
		gs.Move = dir
		gs.Goto(StateMoveInProgress)
	}
}

func (gs *GameState) Tick() {
	gs.DoAction()
}

func (gs *GameState) UndoMove() {
	if gs.ReadyForMove() {
		gs.Current = gs.Previous
		gs.Drawer.DrawScores(gs.Current.Score, gs.Current.HiScore)
		gs.Drawer.DrawBoardNow(&gs.Current.Board)
	}
}
