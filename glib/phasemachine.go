package glib

import (
	"time"
)

type PhaseMachine struct {
	current      int
	counting     int
	timerRunning bool
	timer        CountdownTimer
}

func (ph *PhaseMachine) Begin() {
	ph.counting = 0
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
	inphase := false
	if ph.current == ph.counting {
		inphase = true

		if !ph.timerRunning {
			ph.timer.Reset(period)
			ph.timerRunning = true
		} else {
			if ph.timer.HasExpired() {
				ph.Step()
			}
		}
	}
	ph.counting++
	return inphase
}

func (ph *PhaseMachine) Step() {
	ph.timer.ForceExpire()
	ph.current++
}

func (ph *PhaseMachine) Reset() {
	ph.timer.ForceExpire()
	ph.current = 0
}

func (ph *PhaseMachine) TimedProgress() float64 {
	return ph.timer.GetProgress()
}

func (ph *PhaseMachine) TimedReverseProgress() float64 {
	return ph.timer.GetReverseProgress()
}
