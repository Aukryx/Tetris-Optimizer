package algo

import (
	"time"
)

// Timer struct to measure execution time
type Timer struct {
	startTime time.Time
}

// NewTimer creates a new Timer instance
func NewTimer() *Timer {
	return &Timer{
		startTime: time.Now(),
	}
}

// Start sets the start time to the current time
func (t *Timer) Start() {
	t.startTime = time.Now()
}

// Stop returns the duration since the timer started
func (t *Timer) Stop() time.Duration {
	return time.Since(t.startTime)
}

// Elapsed returns the duration since the timer started
func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.startTime)
}

// ElapsedSeconds returns the number of seconds elapsed since the timer started
func (t *Timer) ElapsedSeconds() float64 {
	return time.Since(t.startTime).Seconds()
}
