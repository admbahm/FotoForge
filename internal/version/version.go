// Package version describes the FotoForge build.
package version

import (
	"fmt"
	"runtime"
	"strings"
)

// Info contains reproducible build metadata.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// New constructs build information. Values are normally supplied with -ldflags.
func New(version, commit, buildDate string) Info {
	return Info{Version: normalize(version), Commit: commit, BuildDate: buildDate, GoVersion: runtime.Version(), OS: runtime.GOOS, Arch: runtime.GOARCH}
}

func normalize(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "dev"
	}
	return strings.TrimPrefix(value, "v")
}

func (i Info) String() string {
	return fmt.Sprintf("fotoforge %s (commit %s, built %s, %s, %s/%s)", i.Version, i.Commit, i.BuildDate, i.GoVersion, i.OS, i.Arch)
}
