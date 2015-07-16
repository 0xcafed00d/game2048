package main

import (
	"time"
)

type PhaseMachine struct {
	current   int
	counting  int
	intimed   bool
	beginTime time.Duration
}

func (ph *PhaseMachine) Once() bool {
	inphase := false
	if ph.current == ph.counting {
		inphase = true
		ph.current++
	}
	ph.counting++
	return inphase
}

func (ph *PhaseMachine) Manual() bool {
	inphase := false
	if ph.current == ph.counting {
		inphase = true
	}
	ph.counting++
	return inphase
}

func (ph *PhaseMachine) Timed(period time.Duration) bool {
	return false
}

func (ph *PhaseMachine) Step() {
	ph.current++
}

func (ph *PhaseMachine) Reset() {
	ph.current = 0
}

func (ph *PhaseMachine) TimedProgress() float64 {
	return 0
}

func (ph *PhaseMachine) TimedReverseProgress() float64 {
	return 1
}
