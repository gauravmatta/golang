# AGENTS.md

## Architecture
This project consists of multiple independent Go modules, each in a subdirectory (e.g., `hello_world/`, `declaring variables/`). Each module has its own `go.mod` file, allowing isolated development and execution. Modules are simple executables demonstrating basic Go concepts, with no inter-module dependencies or shared components.

## Developer Workflows
- **Running a module**: Navigate to the module's directory and execute `go run main.go` (e.g., `cd hello_world && go run main.go` prints "Hello World").
- **Building**: Use `go build` in the module directory to create an executable.
- **No tests present**: The codebase lacks test files; focus on manual execution for verification.

## Conventions and Patterns
- **Variable declarations**: Prefer short assignment `:=` for local variables (e.g., `firstName, lastName := "Raghavi", "Matta"` in `declaring variables/main.go` line 11). Use `var` for package-level or when type is needed.
- **Import style**: Group imports in a block with parentheses (e.g., `hello_world/main.go` lines 3-5).
- **Package comments**: Place descriptive comments on the package declaration line (e.g., `// compiles to executable` in `hello_world/main.go` line 1).
- **Constants**: Define constants with `const` at package level (e.g., `const PI = 3.14` in `declaring variables/main.go` line 19).
- **Zero values**: Variables declared without assignment get zero values (e.g., `emptyint`, `emptystring`, `emptybool` in `declaring variables/main.go` lines 14-15 print `0  false`).

## Key Files
- `hello_world/main.go`: Basic executable structure.
- `declaring variables/main.go`: Variable declaration examples.
- Each `go.mod`: Module definitions using Go 1.26.2.</content>
<parameter name="filePath">/home/gauravmatta/projects/golang/basics/AGENTS.md
