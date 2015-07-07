package main

import (
	tbud "github.com/simulatedsimian/testbuddy"
	"testing"
)

func TestPanics(t *testing.T) {
	tt := &tbud.T{t}

	tt.MustPanic(func(tt *tbud.T) {
		s := MakeStateMachine()
		s.Goto(5)
	})

	tt.MustPanic(func(tt *tbud.T) {
		s := MakeStateMachine()
		s.Return()
	})

	tt.MustPanic(func(tt *tbud.T) {
		s := MakeStateMachine()
		s.Gosub(5)
	})
}

func TestStates(t *testing.T) {

}
