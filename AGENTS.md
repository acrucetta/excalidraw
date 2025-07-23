# Agent Guidelines for multi-draw

## Build/Test/Lint Commands
- Build: `go build ./cmd/main.go` or `go build -o multi-draw ./cmd/main.go`
- Run: `go run ./cmd/main.go`
- Test: `go test ./...` (run all tests) or `go test ./internal/package` (single package)
- Format: `go fmt ./...`
- Vet: `go vet ./...`
- Mod tidy: `go mod tidy`

## Code Style Guidelines
- Use standard Go formatting (`go fmt`)
- Package names: lowercase, single word (e.g., `hub`, `game`, `player`)
- Struct names: PascalCase (e.g., `Hub`, `Client`, `Player`)
- Function names: PascalCase for exported, camelCase for unexported
- Constants: ALL_CAPS with underscores (e.g., `writeWait`, `pongWait`)
- Import grouping: standard library first, then third-party, then local packages
- Use meaningful variable names (e.g., `client`, `message`, `conn`)
- Error handling: always check errors, use `log.Printf` for logging
- Comments: document exported functions and types, use `//` style
- Channel operations: use `select` statements for non-blocking operations
- Struct initialization: use struct literals with field names
- File organization: one main type per file, group related functionality