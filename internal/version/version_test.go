package version

import (
	"strings"
	"testing"
)

func TestNewNormalizesSemanticVersion(t *testing.T) {
	t.Parallel()
	info := New("v1.2.3", "abc123", "2026-07-05T00:00:00Z")
	if info.Version != "1.2.3" {
		t.Fatalf("version = %q", info.Version)
	}
	if !strings.Contains(info.String(), "fotoforge 1.2.3") {
		t.Fatalf("unexpected string: %s", info.String())
	}
}
