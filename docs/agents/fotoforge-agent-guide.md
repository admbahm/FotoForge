# FotoForge Agent Guide

This guide defines the product and engineering constraints for AI coding agents working on FotoForge. It supplements [the root agent instructions](../../AGENTS.md), [the architecture](../architecture.md), and [the safety model](../safety.md).

## Product philosophy

FotoForge helps people understand and preserve personal photo and video archives without asking them to surrender control of their memories.

Its central principle is:

> Preserve memories. Remove chaos.

FotoForge is a digital archivist, not merely a duplicate finder. It should collect evidence, preserve provenance, identify meaningful relationships, explain recommendations, and support reversible workflows. Users must be able to inspect how the tool reached a conclusion instead of blindly trusting it.

Safety and explainability are product behavior, not optional implementation details. Faster processing or a simpler interface does not justify weakening either property.

## Safety model

Treat original media as immutable. Reading content and metadata is allowed; rewriting, normalizing, rotating, transcoding, renaming, moving, or deleting an original is not a safe default.

Agents must design around these constraints:

- Analysis defaults to read-only operation.
- Identical inputs and configuration produce deterministically ordered results.
- A recommendation does not authorize an action.
- Planned changes are complete, stable, and reviewable before execution.
- Expected source hashes and destination conflicts are checked immediately before any copy or move.
- Operations record enough information to explain, resume, verify, or reverse them.
- Cancellation and partial failure leave explicit, inspectable state.
- Unknown or ambiguous cases remain unresolved rather than being guessed.

Do not add destructive behavior as a convenience, cleanup step, error recovery mechanism, or default. Do not weaken safeguards for scripting or non-interactive use.

## Expected behavior for destructive operations

FotoForge does not currently implement destructive operations. Before any future purge or deletion feature is accepted, its design must include all of the following:

1. A separate analysis phase produces a deterministic plan.
2. The plan identifies every affected file, supporting evidence, expected content hash, and reason.
3. The user explicitly selects or approves the plan.
4. Execution requires explicit confirmation and strong, unambiguous flags in non-interactive use.
5. Quarantine precedes permanent removal.
6. Quarantine records original paths, hashes, timestamps, operation identity, and restoration state.
7. Purge operates only on explicitly approved quarantine entries and never on source-discovery results directly.
8. Stale plans, changed files, missing provenance, hash mismatches, and destination conflicts fail safely.
9. Logs and reports clearly distinguish reversible operations from an irreversible purge.
10. Tests cover interruption, partial completion, retries, conflicts, permission errors, and restoration.

Never bypass confirmation because a file appears redundant. Never interpret `--force` as permission to discard evidence or overwrite unrelated files. If safe behavior is unclear, stop and report the ambiguity.

## Folder and artifact classification principles

Folder structure and artifact classification provide evidence, not absolute truth.

- Preserve the observed relative path, filesystem metadata, and containing-tree identity.
- Treat names such as `backup`, `export`, `cache`, or `originals` as signals that require corroboration.
- Use explicit, versioned classification rules with stable precedence.
- Record which rule matched and the observations that supported it.
- Distinguish media, sidecars, thumbnails, caches, application metadata, temporary files, and unknown files without discarding any category automatically.
- Preserve sidecars and related metadata even when they are not independently viewable media.
- Classify uncertain files as unknown and keep them available for review.
- Do not infer that an entire folder is disposable from its name or duplicate percentage.
- Keep artifact classification separate from decisions about retention.

Classification output should be reproducible and suitable for both human-readable reports and future machine-readable manifests.

## Duplicate detection principles

Exact duplicate detection must be based on content, not filenames, paths, timestamps, or file size alone.

- Use streaming BLAKE3 as the primary content digest.
- Use SHA-256 as independent verification where required by the workflow.
- File size may be used only as a safe prefilter, never as duplicate proof.
- Define a duplicate group by explicit digest and size evidence.
- Sort groups and members deterministically.
- Retain every observed path and its provenance; do not reduce a group to a single unexplained winner.
- Explain any recommendation for a preferred copy using recorded criteria such as provenance, metadata completeness, or archive role.
- Treat perceptual similarity, burst photos, edits, transcoded videos, and metadata variants as separate future analyses. Never label them exact duplicates without byte-level evidence.
- Recheck content before a future operation acts on an earlier finding.

A duplicate finding is evidence of identical content. It is not permission to delete a copy.

## Provenance principles

Provenance answers where an observation came from, when it was recorded, and which operation produced or changed state.

Preserve, where applicable:

- audit and operation identifiers;
- source root and normalized relative path;
- observed path representation when normalization could lose information;
- file size, timestamps, type, and content digests;
- scanner and rule versions;
- configuration relevant to the result;
- classification rule and supporting evidence;
- relationships between originals, copies, backup trees, sidecars, and generated artifacts;
- planned and completed operation steps; and
- errors, cancellations, verification results, and restoration history.

Prefer append-only evidence and explicit supersession over silently rewriting historical facts. Schema changes must be versioned and idempotent. Reports should distinguish observed facts from derived conclusions and user decisions.

## CLI UX expectations

The CLI is a safety boundary and should behave predictably in terminals and automation.

- Default to read-only behavior.
- Use clear command names and help text that state whether a command reads, catalogs, copies, moves, quarantines, restores, or purges.
- Require explicit source roots and reject ambiguous arguments.
- Provide deterministic output ordering.
- Send primary output to stdout and diagnostics or structured logs to stderr.
- Use meaningful exit codes and return non-zero status for incomplete or invalid operations.
- Respect `context.Context` cancellation and surface partial results honestly.
- Support a reviewable plan or dry-run representation before future mutating commands.
- Avoid prompts when stdin is not interactive; require explicit flags instead.
- Never hide skipped files, hash failures, permission errors, or conflicting destinations behind a success summary.
- Keep human output concise while retaining a path toward stable machine-readable output.

Bubble Tea and Lip Gloss may improve interactive review, but core safety must not depend on terminal capabilities or styling.

## Testing expectations

Use the Go testing package and deterministic synthetic fixtures. Never commit personal media or private metadata.

Non-trivial behavior should include tests for:

- stable traversal and output ordering;
- empty inputs, unusual names, symlinks, unsupported types, and permission failures;
- cancellation and interrupted operations;
- hashing errors and files that change during observation;
- database initialization, migrations, constraints, and transaction rollback;
- conflicting paths and stale plans;
- classification rule precedence and unknown cases;
- exact duplicate grouping and independent verification;
- CLI arguments, help, stdout/stderr separation, and exit behavior; and
- restoration for every future reversible operation.

Prefer table-driven tests where they improve clarity. Use temporary directories and synthetic byte content. Avoid timing-dependent assertions and platform assumptions unless the test is explicitly platform-specific.

Run the narrow tests while developing, then run `make check` before reporting implementation work complete. A behavior-changing pull request should explain its safety impact and verification coverage.

## Documentation expectations

Update documentation in the same change as behavior.

- Keep command help, README examples, configuration examples, and architecture documents consistent with code.
- State whether a command is read-only or mutating.
- Document evidence inputs, deterministic ordering, exclusions, and limitations.
- Explain safety checks, confirmation requirements, journaling, rollback, and irreversible boundaries.
- Mark planned behavior as future work rather than implying it exists.
- Record significant architectural or safety decisions where future contributors can find them.
- Use synthetic paths and metadata in examples.

Avoid claims such as “safe,” “verified,” or “duplicate” unless the documented implementation establishes exactly what those terms mean.

## Examples of good agent behavior

- Before implementing `audit`, an agent defines deterministic traversal, symlink handling, error reporting, manifest ordering, and cancellation behavior, then tests each boundary.
- An agent reports a file as unknown when classification evidence is insufficient and records why no rule matched.
- An agent groups byte-identical files by digest while preserving every path and makes no retention decision.
- An agent proposes a quarantine plan with source hashes, destinations, conflicts, and restoration metadata before adding execution code.
- An agent notices unrelated working-tree changes and avoids overwriting or staging them.
- An agent updates command documentation and tests together with behavior, runs `make check`, and reports any remaining limitation.

## Examples of bad agent behavior

- Deleting all but one member of each duplicate group automatically.
- Rewriting EXIF data, rotating images, or transcoding videos in place.
- Calling files duplicates because their names, sizes, or capture timestamps match.
- Treating every file in a directory named `cache` or `backup` as disposable.
- Choosing an “original” through an undocumented heuristic and hiding the alternatives.
- Adding `--force` to bypass hash mismatches, stale plans, conflicts, confirmation, or provenance requirements.
- Logging success while silently skipping unreadable or changed files.
- Introducing nondeterministic traversal or report ordering.
- Building purge functionality before audit evidence, plan review, quarantine, restoration, and verification are established.
- Reporting completion without tests, documentation, and `make check` for a behavior change.

## Current implementation priority

The first product feature should be a non-destructive `fotoforge audit` command that walks an explicitly selected directory and produces a deterministic manifest and report. It must not modify source files. Audit evidence should become the foundation for later analysis, explanation, provenance, quarantine planning, restoration, and verification.

## Devlog expectations

Use [the repository devlog](../devlog/README.md) to preserve context from meaningful development sessions. A devlog is appropriate when work establishes or changes architecture, product direction, safety constraints, archive-handling behavior, or contributor workflow. It may also record substantial investigations whose findings will influence later implementation.

At the end of qualifying work, agents should suggest an entry or create one when the task authorizes documentation changes. Generate the initial file with `make devlog slug=short-topic`, then replace the template prompts with concise facts.

Devlogs must:

- describe only work and verification that actually occurred;
- distinguish completed changes from proposals and next steps;
- record relevant decisions and their rationale, especially safety tradeoffs;
- include commands run and observed results without implying unrun checks passed;
- preserve unresolved risks, limitations, and assumptions;
- avoid credentials, personal media, private metadata, and sensitive local paths; and
- remain concise enough to help a future contributor recover context quickly.

Do not create entries merely to make a session appear substantial. Never rewrite history to conceal failed approaches or unresolved safety concerns.
