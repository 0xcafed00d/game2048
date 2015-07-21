package main

import (
	"github.com/simulatedsimian/assert"
	"testing"
	"time"
)

func TestPhaseMachine(t *testing.T) {
	assert.GetFailFunc = func(t *testing.T) assert.FailFunc {
		return t.Fatalf
	}

	pm := PhaseMachine{}

	done := false

	val := 0
	count := 0
	for !done {
		pm.Begin()
		if pm.Once() {
			assert.Equal(t, val, 0)
		}
		if pm.Once() {
			val++
			assert.Equal(t, val, 1)
		}
		if pm.Manual() {
			assert.True(t, count <= 5)
			val++
			count++
			if count > 5 {
				pm.Step()
			}
		}
		if pm.Once() {
			assert.Equal(t, val, 7)
		}
		if pm.Timed(time.Second * 3) {
			val++
		}
		if pm.Once() {
			assert.True(t, val > 30)
			done = true
		}
		time.Sleep(time.Millisecond * 100)
	}
}
