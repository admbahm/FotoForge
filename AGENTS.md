# FotoForge Agent Instructions

## Mission

FotoForge is a safety-first CLI for auditing, reconciling, organizing, and preserving personal photo and video archives.

## Non-negotiable safety rules

- Never implement automatic deletion.
- Never modify original media.
- Prefer copy and quarantine workflows over destructive actions.
- Any purge or delete behavior must require explicit confirmation and strong, unambiguous flags.
- Every recommendation must be deterministic and explainable.
- Preserve provenance whenever possible.

Read [the full agent guide](docs/agents/fotoforge-agent-guide.md) before designing behavior that classifies media, identifies duplicates, moves files, or changes the catalog.

## Engineering rules

- Use idiomatic Go.
- Keep packages small and cohesive.
- Pass `context.Context` into long-running and I/O-bound operations.
- Prefer clear code over clever code.
- Add tests for non-trivial behavior, including failure and interruption paths.
- Run `make check` before reporting implementation work complete.

## Tech stack

Cobra, Viper, `slog`, SQLite, BLAKE3, SHA-256, Bubble Tea, and Lip Gloss.

## Standard commands

```sh
make fmt
make vet
make test
make build
make smoke
make check
```

## Current priority

Implement non-destructive audit functionality first. Do not skip ahead to destructive archive operations.

## Devlogs

At the end of meaningful work, suggest or create a devlog entry when architectural, product, safety, or workflow decisions were made. Use `make devlog slug=short-topic` and follow [the devlog guidance](docs/devlog/README.md). Do not fabricate accomplishments, verification, or decisions. Clearly distinguish completed work from planned work.
