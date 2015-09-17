package glib

import (
	"github.com/simulatedsimian/assert"
	"testing"
)

// pack a number of values into a slice containing those values
func pack(vals ...interface{}) []interface{} {
	return vals
}

func TestPanics(t *testing.T) {

	assert.MustPanic(t, func(t *testing.T) {
		s := StateMachine{}
		s.Goto(5)
	})

	assert.MustPanic(t, func(t *testing.T) {
		s := StateMachine{}
		s.Return()
	})

	assert.MustPanic(t, func(t *testing.T) {
		s := StateMachine{}
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

	sm := StateMachine{}

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
	assert.Equal(t, pack(sm.CurrentId()), pack(1, true))
	sm.DoAction()
	sm.Goto(2)
	assert.Equal(t, pack(sm.CurrentId()), pack(2, true))
	sm.DoAction()
	sm.Goto(3)
	assert.Equal(t, pack(sm.CurrentId()), pack(3, true))
	sm.DoAction()
	sm.Goto(1)
	assert.Equal(t, pack(sm.CurrentId()), pack(1, true))
	sm.DoAction()
	sm.Gosub(2)
	assert.Equal(t, pack(sm.CurrentId()), pack(2, true))
	sm.DoAction()
	sm.Gosub(1)
	assert.Equal(t, pack(sm.CurrentId()), pack(1, true))
	sm.DoAction()
	sm.Return()
	assert.Equal(t, pack(sm.CurrentId()), pack(2, true))
	sm.Return()
	assert.Equal(t, pack(sm.CurrentId()), pack(1, true))
	sm.DoAction()
	sm.Goto(3)
	assert.Equal(t, pack(sm.CurrentId()), pack(3, true))

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
