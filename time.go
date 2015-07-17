package main

import (
	"time"
)

func getTimeNow() time.Duration {
	return time.Duration(time.Now().UnixNano())
}

type CountdownTimer struct {
	startTime   time.Duration
	duration    time.Duration
	pausedAt    time.Duration
	pausedCount int
}

func MakeCountdownTimer(duration time.Duration) CountdownTimer {
	return CountdownTimer{startTime: getTimeNow(), duration: duration}
}

// GetElapsedTime returns the elapsed time since the time started
// (taking into account pause status)
func (t *CountdownTimer) GetElapsedTime() time.Duration {
	if t.pausedCount > 0 {
		return t.pausedAt - t.startTime
	} else {
		return getTimeNow() - t.startTime
	}
}

// Reset resets the timer with new duration. Can be called any time irrespective
// of the Expried Status of the timer
func (t *CountdownTimer) Reset(duration time.Duration) {
	t.pausedCount = 0
	t.duration = duration
	t.startTime = getTimeNow()
}

// Pause pauses the Countdown timer. call UnPause to resume or Reset to start timer again.
func (t *CountdownTimer) Pause() {
	if t.pausedCount == 0 {
		t.pausedAt = getTimeNow()
	}
	t.pausedCount++
}

// Unpause resumes the countdown
func (t *CountdownTimer) Unpause() {
	if t.pausedCount == 1 {
		pausedFor := getTimeNow() - t.pausedAt
		t.startTime += pausedFor
	}

	if t.pausedCount > 0 {
		t.pausedCount--
	}
}

// IsPaused returns true if timer is paused
func (t *CountdownTimer) IsPaused() bool {
	return t.pausedCount > 0
}

// ForceExpire makes the timer expire before its duration
func (t *CountdownTimer) ForceExpire() {
	t.duration = 0
}

// HasExpired returns true if the timer has expired
func (t *CountdownTimer) HasExpired() bool {
	return t.GetElapsedTime() >= t.duration
}

// GetTimeRemaining returns how much time is remaining until the timer expires
func (t *CountdownTimer) GetTimeRemaining() time.Duration {
	if t.HasExpired() {
		return 0
	} else {
		return t.duration - t.GetElapsedTime()
	}
}

// GetProgress returns how far along the duration the timer has gone (from 0.0 -> 1.0)
func (t *CountdownTimer) GetProgress() float64 {
	if t.HasExpired() {
		return 1.0
	} else {
		return float64(t.GetElapsedTime()) / float64(t.duration)
	}
}

// GetReverseProgress returns how far along the duration the timer has gone (from 1.0 -> 0.0) in reverse
func (t *CountdownTimer) GetReverseProgress() float64 {
	return 1.0 - t.GetProgress()
}
