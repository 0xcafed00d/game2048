package glib

import (
	"github.com/simulatedsimian/assert"
	"testing"
)

func TestPanics(t *testing.T) {

	assert.MustPanic(t, func(t *testing.T) {
		s := MakeStateMachine()
		s.Goto(5)
	})

	assert.MustPanic(t, func(t *testing.T) {
		s := MakeStateMachine()
		s.Return()
	})

	assert.MustPanic(t, func(t *testing.T) {
		s := MakeStateMachine()
		s.Gosub(5)
	})
}

func makeStateFunction(name string, result *[]string) StateFunc {
	return func(sm *StateMachine) {
		*result = append(*result, name)
	}
}

func TestStates(t *testing.T) {
	var result []string

	sm := MakeStateMachine()

	sm.AddState(1, State{
		Enter:  makeStateFunction("s1Enter", &result),
		Action: makeStateFunction("s1Action", &result),
		Exit:   makeStateFunction("s1Leave", &result),
	})

	sm.AddState(2, State{
		Enter:  makeStateFunction("s2Enter", &result),
		Action: makeStateFunction("s2Action", &result),
		Exit:   makeStateFunction("s2Leave", &result),
	})

	sm.AddState(3, State{})

	sm.DoAction()
	sm.Goto(1)
	sm.DoAction()
	sm.Goto(2)
	sm.DoAction()
	sm.Goto(3)
	sm.DoAction()
	sm.Goto(1)
	sm.DoAction()
	sm.Gosub(2)
	sm.DoAction()
	sm.Gosub(1)
	sm.DoAction()
	sm.Return()
	sm.Return()
	sm.DoAction()
	sm.Goto(3)

	expect := []string{
		"s1Enter",
		"s1Action",
		"s1Leave",
		"s2Enter",
		"s2Action",
		"s2Leave",
		"s1Enter",
		"s1Action",
		"s2Enter",
		"s2Action",
		"s1Enter",
		"s1Action",
		"s1Leave",
		"s2Leave",
		"s1Action",
		"s1Leave",
	}

	assert.Equal(t, result, expect)
}
