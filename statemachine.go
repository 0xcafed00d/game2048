package main

import (
	"fmt"
)

type StateFunc func(sm *StateMachine)

type StateInfo struct {
	Enter  StateFunc
	Action StateFunc
	Exit   StateFunc
}

type StateMachine struct {
	states       map[int]*StateInfo
	currentState *StateInfo
	gosubStack   []*StateInfo
}

func MakeStateMachine() StateMachine {
	return StateMachine{states: make(map[int]*StateInfo)}
}

func (sm *StateMachine) AddState(id int, si StateInfo) {
	sm.states[id] = &si
}

func (sm *StateMachine) DoAction() {
	if sm.currentState != nil {
		if sm.currentState.Action != nil {
			sm.currentState.Action(sm)
		}
	}
}

func (sm *StateMachine) Goto(id int) {
	if sm.currentState != nil {
		if sm.currentState.Exit != nil {
			sm.currentState.Exit(sm)
		}
	}

	sm.currentState = sm.states[id]

	if sm.currentState != nil {
		if sm.currentState.Enter != nil {
			sm.currentState.Enter(sm)
		}
	} else {
		panic(fmt.Sprint("Invalid State ID: ", id))
	}
}

func (sm *StateMachine) Gosub(id int) {
	sm.gosubStack = append(sm.gosubStack, sm.currentState)

	sm.currentState = sm.states[id]

	if sm.currentState != nil {
		if sm.currentState.Enter != nil {
			sm.currentState.Enter(sm)
		}
	} else {
		panic(fmt.Sprint("Invalid State ID: ", id))
	}
}

func (sm *StateMachine) Return(id int) {
	if len(sm.gosubStack) > 0 {
		if sm.currentState.Exit != nil {
			sm.currentState.Exit(sm)
		}

		sm.currentState = sm.gosubStack[len(sm.gosubStack)-1]
		sm.gosubStack = sm.gosubStack[0 : len(sm.gosubStack)-1]
	} else {
		panic("Empty Gosub Stack")
	}
}
