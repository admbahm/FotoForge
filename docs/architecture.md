# Architecture

FotoForge separates observation, analysis, planning, and execution. This boundary is a safety mechanism, not just code organization.

1. **Audit** observes files and stores immutable evidence in the catalog.
2. **Analyze** derives reproducible findings from evidence without touching source media.
3. **Plan** turns selected findings into a complete, human-reviewable operation manifest.
4. **Execute** applies an approved manifest and journals every completed step.
5. **Verify or restore** proves the result or reverses journaled operations.

SQLite is the local system of record. Schema changes are versioned and applied idempotently. Content identity uses streaming BLAKE3; SHA-256 is retained as an independent verification digest. Filesystem paths alone are never considered identity.

Adapters for ExifTool, ffprobe, and Immich will remain outside domain logic. Reports and the terminal UI consume catalog data and do not own analysis rules.

## Dependency direction

`cmd/fotoforge` → `internal/cli` → domain and infrastructure packages. Domain packages must not import the CLI, TUI, or report layers. Interfaces should be introduced only at real process, filesystem, clock, or network boundaries.
