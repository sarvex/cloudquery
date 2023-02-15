package plugin

import (
	"testing"
)

func TestGoogleAnalytics(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The GoogleAnalytics() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := GoogleAnalytics()
	name := p.Name()
	if name != "google-analytics" {
		t.Errorf("Name() = %q, want %q", name, "google-analytics")
	}
}
