# Repository Guidelines

## Project Structure & Module Organization
- Source: `internal/domain` with subpackages `game` and `player` for core domain types (e.g., `ID`, `Status`, `Format`, `Game`).
- Tests: `test/game` and `test/player` contain table-driven unit tests for domain behavior.
- Tooling: `go.mod`/`go.sum` manage dependencies; optional Nix shell via `flake.nix` for a pinned Go toolchain.

## Build, Test, and Development Commands
- `go test ./...`: run all tests.
- `go test -cover ./...`: run tests with coverage.
- `go vet ./...`: static checks for common issues.
- `go fmt ./...`: format code before committing.
- `go build ./...`: compile packages (no binary entrypoint yet).
- If using Nix: `nix develop` to enter the dev shell with Go installed.

## Coding Style & Naming Conventions
- Indentation: 2 spaces (see `.editorconfig`); LF line endings, UTF‑8.
- Go style: exported identifiers use PascalCase (`ID`, `NewID`), unexported use camelCase; package names are short and lowercase (`game`, `player`).
- Files follow focused names by concept (`id.go`, `status.go`, `format.go`).
- Run `go fmt` and prefer idiomatic Go; `go vet` before pushing.

## Testing Guidelines
- Framework: standard `testing` package; table-driven tests preferred.
- Location: keep new tests under `test/<package>` with `*_test.go` and functions `TestXxx`.
- Coverage: aim for meaningful coverage of happy paths and edge cases for value objects and state transitions.
- Run: `go test ./...` locally before opening a PR.

## Test-Driven Development (TDD)
- Always follow Red-Green-Refactor.
- Red: write a failing test that specifies the behavior.
- Green: implement the minimal code to make the test pass.
- Refactor: improve design and readability with tests staying green.
- Bug fixes: add a failing regression test first, then fix.

## Commit & Pull Request Guidelines
- Commits: concise, imperative subject; optionally scope with package (e.g., `game: add status transitions`).
- Include rationale when non-obvious; group related changes.
- PRs: clear description, link issues if applicable, include test updates and examples when adding new domain types or transitions.
- Checks: ensure formatting (`go fmt`), vetting, and tests pass.

## Architecture Notes
- Domain-driven core: small, focused value objects and entities under `internal/domain`; no application/server layer yet.
- Keep types cohesive and immutable where possible; prefer constructors like `NewID`, `NewFormatFromString`.
- Shared IDs: use neutral VOs under `internal/domain/ids` (e.g., `ids.PlayerID`) when multiple aggregates reference the same identifier to avoid cross-package coupling.
