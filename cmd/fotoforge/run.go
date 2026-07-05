package main

import (
	"context"
	"os"

	"github.com/fotoforge/fotoforge/internal/cli"
	"github.com/fotoforge/fotoforge/internal/version"
)

var (
	buildVersion = "dev"
	buildCommit  = "unknown"
	buildDate    = "unknown"
)

func run(ctx context.Context) error {
	info := version.New(buildVersion, buildCommit, buildDate)
	return cli.New(info, os.Stdout, os.Stderr).ExecuteContext(ctx)
}
