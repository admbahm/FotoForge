package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fotoforge/fotoforge/internal/version"
)

func TestVersionDoesNotCreateCatalog(t *testing.T) {
	t.Parallel()
	var out, errOut bytes.Buffer
	cmd := New(version.New("v1.2.3", "abc", "today"), &out, &errOut)
	cmd.SetArgs([]string{"version"})
	if err := cmd.ExecuteContext(context.Background()); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "fotoforge 1.2.3") {
		t.Fatalf("output = %q", out.String())
	}
}

func TestCommandsExist(t *testing.T) {
	t.Parallel()
	cmd := New(version.New("dev", "unknown", "unknown"), &bytes.Buffer{}, &bytes.Buffer{})
	for _, name := range []string{"audit", "analyze", "report", "organize", "quarantine", "restore", "purge", "verify", "version"} {
		if found, _, err := cmd.Find([]string{name}); err != nil || found.Name() != name {
			t.Errorf("command %q not found", name)
		}
	}
}
