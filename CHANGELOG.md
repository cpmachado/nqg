# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub Actions workflow (`.github/workflows/go.yml`) running on every push, with jobs for build, `go mod tidy -diff`, `govulncheck`, `gofmt` format check and `go test`. (#4)

### Changed
- Moved `main.go` to `cmd/nqg/main.go`; `make build` now builds `./cmd/nqg`.
- Updated Go version to 1.27.2.
- Updated the README and the build script.

### Fixed
- Typo in the README.

## [0.0.2] - 2026-05-28

### Changed
- The version is now read from the Go runtime build info instead of being set in the Makefile.
- Updated Go version to 1.26.3.
- Moved `init` and global variables into `main`.
- Applied `go fix` to the `snq` package.

### Fixed
- Lint issues reported by the linter in `main.go`.

## [0.0.1] - 2025-02-17

### Added
- Usage and installation instructions in the README.
- Package documentation (godoc) for `snq`, including a reference to the original `snq` project.

### Changed
- Capitalized the solver's name: `snq` is now `SNQ`.
- Updated Go version to 1.24.0.

## [0.0.0] - 2025-02-02

### Added
- Initial project skeleton: license, `Makefile`, README, `.gitignore`, `go.mod` and a minimal `main.go`. (2025-02-01)
- Version command support in the CLI and Makefile.
- `snq` package implementing the original Simple N-Queen backtracking solver (MVP).

### Changed
- Fixed the module path to the vanity URL `go.cpmachado.pt/nqg`.
- Refactored initialization out of `main` into `init`.
- Annotated the existing source code.
- Updated Go version to 1.23.5.

[Unreleased]: https://github.com/cpmachado/nqg/compare/v0.0.2...HEAD
[0.0.2]: https://github.com/cpmachado/nqg/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/cpmachado/nqg/compare/v0.0.0...v0.0.1
[0.0.0]: https://github.com/cpmachado/nqg/releases/tag/v0.0.0
