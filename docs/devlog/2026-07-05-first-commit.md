# First Commit: From Family Archive to FotoForge

**Date:** July 5, 2026  
**Repository:** `git@github.com:admbahm/FotoForge.git`  
**Commit:** `5c91b764ca65c2f94cbcbb51932f59d29cd95499`

## Session Summary

FotoForge went from an idea to a public GitHub scaffold in this first session. The initial foundation includes a command-line interface, project documentation, continuous integration, tests, semantic build information, structured logging, configuration loading, SQLite initialization, hashing utilities, and the repository hygiene expected of a maintainable open-source project.

The implementation remains deliberately foundational. Commands that will eventually inspect or modify archive state are placeholders, and no destructive file behavior has been implemented.

## Origin Story

FotoForge emerged from the practical work of reconciling a chaotic family photo archive. The archive contained **99,949 files**, enough that manual inspection and intuition were no longer reliable ways to understand its contents.

A SHA-256 manifest completed successfully and provided a deterministic basis for analysis. It identified **15,916 duplicate hash groups**. One backup tree, Drive3BAK, contained **22,505 files**, of which **19,891 were duplicates**—an **88.4% duplicate rate**. Extracting content unique to Drive3BAK left **2,614 unique files**.

That smaller set still mixed meaningful media with generated or incidental files. Artifact classification reduced the review set to **893 real media files** and **1,721 artifacts**.

These results shaped FotoForge's product philosophy. Hashes can establish content identity, but useful archive reconciliation also requires provenance, backup-tree awareness, artifact classification, and explanations that a person can review before anything changes.

## Core Philosophy

FotoForge is guided by a small set of non-negotiable principles:

- **Preserve memories. Remove chaos.**
- **Safety first.** Archive integrity takes priority over convenience or speed.
- **Originals are sacred.** Original media must be identified and protected.
- **Never delete automatically.** Destructive action must never be inferred from an analysis result.
- **Quarantine before purge.** Reversible isolation must precede any permanent removal workflow.
- **Deterministic evidence over guessing.** Recommendations must follow reproducible observations and rules.
- **Every recommendation must be explainable.** Users should be able to inspect the evidence and reasoning behind it.
- **Users should not have to blindly trust the tool.** FotoForge must make verification possible and preserve enough provenance to challenge or reverse its conclusions.

These principles apply to implementation details as well as product behavior. Future file operations must be explicit, reviewable, journaled, and reversible wherever reversal is technically possible.

## Product Insight

FotoForge is not merely a duplicate finder. It is a **digital archivist**.

Its job is to help users:

- understand the structure and condition of their archive;
- preserve provenance across originals, copies, exports, and backup trees;
- understand why each recommendation was made;
- distinguish real photos and videos from thumbnails, caches, sidecars, and other artifacts;
- identify backup trees and measure how they overlap with the primary archive; and
- safely simplify chaotic collections without putting memories at unnecessary risk.

Duplicate detection is one capability within that larger responsibility, not the product's defining boundary.

## Engineering Pattern

OpenHunt, The Forge, and FotoForge share an engineering pattern: each addresses a real problem encountered through lived experience. Their common thread is reducing uncertainty through deterministic systems.

That pattern matters for future contributors. FotoForge should favor observable facts, stable manifests, explicit state transitions, and inspectable reports over opaque heuristics. When heuristics are eventually useful, their inputs, confidence, and limitations must remain visible.

## First Commit Contents

The first commit established:

- a Cobra command scaffold with the initial FotoForge command surface;
- configuration loading through Viper;
- structured logging with `slog`;
- basic SQLite catalog initialization;
- a streaming BLAKE3 and SHA-256 hashing utility;
- a Bubble Tea and Lip Gloss foundation for future terminal interfaces;
- architecture, safety, contribution, security, and project documentation;
- an MIT license;
- a Makefile for formatting, vetting, testing, building, and smoke checks;
- a `.gitignore` covering Go output, SQLite artifacts, local configuration, editors, and operating-system files;
- GitHub Actions continuous integration;
- Dependabot configuration;
- CODEOWNERS assigned to `@admbahm`; and
- structured issue templates and a pull request template with explicit safety considerations.

The first commit was pushed to `main`, which tracks `origin/main`, at commit `5c91b764ca65c2f94cbcbb51932f59d29cd95499`. The local working tree was clean after the push, and the remote commit was verified to match local `HEAD`.

## Verification Results

Before the initial push, the scaffold passed the complete local quality gate:

- `gofmt` verification passed;
- `go vet ./...` passed;
- `go test ./...` passed;
- `go build ./cmd/fotoforge` passed;
- `fotoforge version` and `fotoforge --help` smoke checks passed;
- `make check` passed; and
- the commit on remote `main` was verified against the local commit hash.

## Roadmap Seeds

The current and potential command surface provides a useful outline for future work:

- `fotoforge audit` — inventory files and produce deterministic observations;
- `fotoforge analyze` — derive findings from cataloged evidence;
- `fotoforge explain` — show the evidence and rules behind a finding or recommendation;
- `fotoforge history` — inspect provenance and recorded archive operations;
- `fotoforge quarantine` — move explicitly selected files into reversible isolation;
- `fotoforge restore` — restore quarantined files to their recorded locations;
- `fotoforge purge` — permanently remove only explicitly reviewed quarantine entries;
- `fotoforge verify` — verify content hashes, catalog integrity, and operation results.

These are roadmap seeds, not commitments to immediate implementation. In particular, quarantine, restore, and purge require a carefully designed plan-and-journal model before they should touch user files.

## Immediate Next Step

The first real feature should be `fotoforge audit`: a non-destructive inventory command that walks a directory and produces a deterministic manifest and human-readable report.

The initial implementation should collect stable filesystem observations, stream file contents through the existing hash utility, record errors without abandoning the entire audit, and avoid modifying source files. It should define deterministic traversal and output ordering, accept cancellation through `context.Context`, and make partial or interrupted runs explicit. This creates the evidence layer that later analysis, provenance, explanation, and verification features can safely build upon.
