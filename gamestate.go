package main

import (
	//"fmt"
	"github.com/simulatedsimian/game2048/glib"
)

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
	gs := &GameState{}

	gs.AddState(StateGameOver, glib.State{
		Enter: func(sm *glib.StateMachine) {
		},
		Action: func(sm *glib.StateMachine) {
		},
		Exit: func(sm *glib.StateMachine) {
		},
	})

	gs.AddState(StateReadyForMove, glib.State{
		Enter: func(sm *glib.StateMachine) {
			gs.Drawer.DrawScores(gs.Current.Score, gs.Current.HiScore)
			gs.Drawer.DrawBoardNow(&gs.Current.Board)
		},
		Action: func(sm *glib.StateMachine) {
		},
		Exit: func(sm *glib.StateMachine) {
		},
	})

	gs.AddState(StateMoveInProgress, glib.State{
		Enter: func(sm *glib.StateMachine) {
			moves, score := gs.Current.Board.SingleStep(gs.Move)
			if len(moves) == 0 {
				sm.Goto(StateReadyForMove)
			} else {
				gs.Current.Score += score
			}
		},
		Action: func(sm *glib.StateMachine) {
			sm.Goto(StateMoveInProgress)
		},
		Exit: func(sm *glib.StateMachine) {
		},
	})

	return gs
}

func (gs *GameState) NewGame() {
	gs.Current.Board.Reset()

	x, y, _ := gs.Current.Board.FindFreeCell()
	gs.Current.Board[x][y].val = 2

	x, y, _ = gs.Current.Board.FindFreeCell()
	gs.Current.Board[x][y].val = 2

	gs.Current.Score = 0
	gs.Previous = gs.Current

	gs.Goto(StateReadyForMove)
}

func (gs *GameState) ReadyForMove() bool {
	id, ok := gs.CurrentId()
	return ok && id == StateReadyForMove
}

func (gs *GameState) DoMove(dir Direction) {
	if gs.ReadyForMove() {

		gs.Previous = gs.Current
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
