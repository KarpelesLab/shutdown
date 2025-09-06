# CLAUDE.md for shutdown package

## Build & Test Commands
- Build & format code: `make`
- Run tests: `make test`
- Run single test: `go test -v -run TestName`
- Update dependencies: `make deps`

## Code Style Guidelines
- **Formatting**: Use goimports for consistent formatting
- **Imports**: Group standard library imports first, then third-party packages
- **Naming**: Use CamelCase for exported functions, camelCase for non-exported
- **Error Handling**: Return errors rather than using panic, use structured logging with slog
- **Documentation**: All exported functions must have godoc comments
- **Concurrency**: Use sync package primitives for synchronization
- **Testing**: Write unit tests for all exported functions
- **Logging**: Use log/slog package for structured logging

## Package Conventions
- The package provides graceful shutdown functionality for Go applications
- Use SetupSignals() early in main() to handle OS signals
- Use Defer() to register cleanup functions
- Call Wait() to block until shutdown is triggered