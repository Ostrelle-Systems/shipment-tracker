package circuitbreaker

import (
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	cb := New(3, 2, 100*time.Millisecond)

	// Initial state should be closed
	if cb.State() != Closed {
		t.Error("expected initial state to be Closed")
	}

	// Should allow requests in closed state
	if !cb.Allow() {
		t.Error("expected Allow() to return true in Closed state")
	}

	// Record failures
	for i := 0; i < 3; i++ {
		cb.RecordFailure()
	}

	// Should be open after threshold
	if cb.State() != Open {
		t.Errorf("expected state to be Open, got %d", cb.State())
	}

	// Should not allow requests in open state
	if cb.Allow() {
		t.Error("expected Allow() to return false in Open state")
	}

	// Wait for timeout
	time.Sleep(150 * time.Millisecond)

	// Should allow request and transition to half-open
	if !cb.Allow() {
		t.Error("expected Allow() to return true after timeout")
	}

	// Record successes to close circuit
	for i := 0; i < 2; i++ {
		cb.RecordSuccess()
	}

	if cb.State() != Closed {
		t.Error("expected state to be Closed after successes")
	}
}

func TestHalfOpenFailure(t *testing.T) {
	cb := New(2, 2, 50*time.Millisecond)

	// Open the circuit
	cb.RecordFailure()
	cb.RecordFailure()

	// Wait for timeout
	time.Sleep(60 * time.Millisecond)

	// Allow request (half-open)
	cb.Allow()

	// Record failure in half-open state
	cb.RecordFailure()

	if cb.State() != Open {
		t.Error("expected state to be Open after failure in HalfOpen")
	}
}
