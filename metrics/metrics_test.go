package metrics

import (
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {
	m := New()

	// Test initial state
	if m.TotalRequests() != 0 {
		t.Errorf("expected 0 total requests, got %d", m.TotalRequests())
	}

	// Test incrementing
	m.IncrRequest("/api/v1/shipments")
	m.IncrRequest("/api/v1/shipments")
	m.IncrRequest("/api/v1/health")

	// Test individual path counts
	if got := m.GetRequestCount("/api/v1/shipments"); got != 2 {
		t.Errorf("expected 2 requests for /api/v1/shipments, got %d", got)
	}

	if got := m.GetRequestCount("/api/v1/health"); got != 1 {
		t.Errorf("expected 1 request for /api/v1/health, got %d", got)
	}

	// Test total
	if got := m.TotalRequests(); got != 3 {
		t.Errorf("expected 3 total requests, got %d", got)
	}

	// Test uptime
	if m.Uptime() < time.Duration(0) {
		t.Error("uptime should be non-negative")
	}
}
