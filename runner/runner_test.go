package runner

import (
	"errors"
	"net/http"
	"testing"
)

type erroringRoundTripper struct{}

func (erroringRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("connection refused")
}

// TestLoggingRoundTripper_TransportError ensures that a transport-level error
// (where the underlying RoundTripper returns a nil response) does not cause a
// nil pointer dereference when the response status is inspected for logging.
func TestLoggingRoundTripper_TransportError(t *testing.T) {
	l := &LoggingRoundTripper{Name: "test", RoundTripper: erroringRoundTripper{}}
	req, err := http.NewRequest(http.MethodGet, "http://example.invalid", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	resp, err := l.RoundTrip(req)
	if err == nil {
		t.Fatal("expected an error from RoundTrip")
	}
	if resp != nil {
		t.Fatalf("expected a nil response, got %v", resp)
	}
}
