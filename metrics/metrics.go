package metrics

import (
	"sync"
	"time"
)

type Metrics struct {
	mu           sync.RWMutex
	requestCount map[string]int64
	startTime    time.Time
}

func New() *Metrics {
	return &Metrics{
		requestCount: make(map[string]int64),
		startTime:    time.Now(),
	}
}

func (m *Metrics) IncrRequest(path string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.requestCount[path]++
}

func (m *Metrics) GetRequestCount(path string) int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.requestCount[path]
}

func (m *Metrics) TotalRequests() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var total int64
	for _, count := range m.requestCount {
		total += count
	}
	return total
}

func (m *Metrics) Uptime() time.Duration {
	return time.Since(m.startTime)
}
