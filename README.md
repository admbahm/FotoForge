# FotoForge

[![CI](https://github.com/fotoforge/fotoforge/actions/workflows/ci.yml/badge.svg)](https://github.com/fotoforge/fotoforge/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**Preserve memories. Remove chaos.**

FotoForge is a safety-first command-line application for auditing, reconciling, organizing, and preserving large personal photo and video collections. It is designed for people whose archives have outgrown ad hoc scripts and conventional duplicate finders.

> FotoForge never automatically deletes user data. Every mutating workflow must be deterministic, explainable, explicitly approved, and reversible.

## Status

FotoForge is in its foundation phase. The CLI and catalog infrastructure are usable for development, but collection analysis and media operations are placeholders. Do not rely on unreleased commands for production archives.

## Philosophy

- **Safety before convenience:** read-only analysis is the default; destructive actions require a separately reviewed plan.
- **Determinism:** identical inputs and configuration produce identical findings.
- **Explainability:** reports retain the evidence and rule that produced each conclusion.
- **Provenance:** the catalog records where observations came from and how they changed.
- **Reversibility:** organization and quarantine operations retain enough information to undo them.
- **Verification:** BLAKE3 provides fast primary identity and SHA-256 provides independent verification.

## Commands

| Command | Purpose |
|---|---|
| `fotoforge` | Show help and available commands |
| `fotoforge version` | Print semantic version and build metadata (`--json` supported) |
| `fotoforge audit` | Inventory a collection without changing source files |
| `fotoforge analyze` | Derive findings from cataloged observations |
| `fotoforge report` | Generate an explainable collection report |
| `fotoforge organize` | Plan reversible organization |
| `fotoforge quarantine` | Isolate explicitly selected media with restoration metadata |
| `fotoforge restore` | Restore quarantined media |
| `fotoforge purge` | Permanently remove explicitly approved quarantined data |
| `fotoforge verify` | Verify hashes, catalog records, and provenance |

Except for `version`, these commands currently initialize the SQLite catalog and report that implementation is pending. They do not modify media.

Global flags:

```text
--config string   configuration file
--verbose         enable verbose structured logging
```

Configuration is loaded in this order: defaults, config file, `FOTOFORGE_*` environment variables, then command-line flags. The default catalog is `.fotoforge/catalog.db`. See [examples/fotoforge.yaml](examples/fotoforge.yaml).

## Architecture

FotoForge keeps its public surface deliberately small. The executable lives in `cmd/fotoforge`; application code is under `internal` and grouped by domain:

- `cli`, `config`, `logger`, `version`: process boundary and build/runtime concerns.
- `db`, `catalog`, `provenance`: durable facts, schema evolution, and traceability.
- `audit`, `analyze`, `artifacts`, `media`: future inventory and classification pipeline.
- `hash`: streaming BLAKE3 and SHA-256 primitives.
- `report`, `tui`: non-mutating presentation layers.

Future external integrations (ExifTool, ffprobe, Immich) will enter through narrow adapters at process or network boundaries. Domain packages will not depend directly on CLI or presentation code. See [docs/architecture.md](docs/architecture.md) and [docs/safety.md](docs/safety.md).

## Roadmap

1. Read-only filesystem audit with resumable catalog writes.
2. Metadata extraction through ExifTool and ffprobe.
3. Exact duplicate and backup-tree analysis with evidence records.
4. Artifact classification and timeline generation.
5. Reviewable organization and quarantine plans with rollback journals.
6. Static HTML reports and optional Immich integration.

Duplicate detection is intentionally not part of the initial foundation.

## Development

Prerequisites: Go 1.24 or newer, Git, and Make. SQLite is embedded through a pure-Go driver; no system SQLite library is required.

```sh
git clone https://github.com/fotoforge/fotoforge.git
cd fotoforge
make check
make build
./bin/fotoforge version
```

Release builds inject version data:

```sh
make build VERSION=0.1.0
```

Useful targets are `fmt`, `fmt-check`, `vet`, `test`, `build`, `check`, and `clean`.

## Contributing

Start with [CONTRIBUTING.md](CONTRIBUTING.md), follow the [Code of Conduct](CODE_OF_CONDUCT.md), and report vulnerabilities according to [SECURITY.md](SECURITY.md). Changes that can move or remove media require explicit threat analysis, tests for interruption and rollback, and documentation of their safety properties.

## License

Licensed under the [MIT License](LICENSE).
