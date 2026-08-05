package circuitbreaker

import (
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu              sync.RWMutex
	state           State
	failureCount    int
	successCount    int
	failureThreshold int
	successThreshold int
	timeout         time.Duration
	lastFailureTime time.Time
}

func New(failureThreshold, successThreshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:           Closed,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		timeout:         timeout,
	}
}

func (cb *CircuitBreaker) Allow() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	switch cb.state {
	case Closed:
		return true
	case Open:
		if time.Since(cb.lastFailureTime) > cb.timeout {
			return true
		}
		return false
	case HalfOpen:
		return true
	default:
		return false
	}
}

func (cb *CircuitBreaker) RecordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case HalfOpen:
		cb.successCount++
		if cb.successCount >= cb.successThreshold {
			cb.state = Closed
			cb.failureCount = 0
			cb.successCount = 0
		}
	case Closed:
		cb.failureCount = 0
	}
}

func (cb *CircuitBreaker) RecordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case HalfOpen:
		cb.state = Open
		cb.lastFailureTime = time.Now()
		cb.successCount = 0
	case Closed:
		cb.failureCount++
		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
			cb.lastFailureTime = time.Now()
		}
	}
}

func (cb *CircuitBreaker) State() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}
