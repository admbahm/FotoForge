# Contributing to FotoForge

Thank you for helping preserve personal media safely.

## Workflow

1. Search existing issues and open one for substantial behavioral changes.
2. Create a focused branch and keep commits reviewable.
3. Run `make check` before opening a pull request.
4. Explain user-visible behavior, safety implications, and rollback behavior in the PR.

Code must be idiomatic Go, pass `gofmt` and `go vet`, and include tests proportional to risk. Keep packages cohesive and add interfaces only at demonstrated boundaries. Pass `context.Context` into operations that perform I/O or may block.

Never include personal media, private metadata, credentials, or machine-specific catalog files in fixtures. Use synthetic files under `testdata`.

By contributing, you agree that your contribution is licensed under MIT.
