# Building CyberBASIC

This document describes how to build and run the CyberBASIC compiler.

## Prerequisites

- Go 1.24.1 or later

## Building the Compiler

### Windows

```bash
# Build for Windows only
.\build.bat

# Build for all platforms
go run build.go
```

### Linux/macOS

```bash
# Build for current platform only
go build -o cyberbasic cmd/cyberbasic/main.go

# Build for all platforms
go run build.go
```

## Creating Distribution Packages

After building for all platforms, you can create distribution packages with:

```bash
go run package.go
```

This will create ZIP files in the `packages` directory for each platform.

## Running the Compiler

```bash
# Compile a CyberBASIC file
./cyberbasic examples/hello.cyber

# Compile and run a CyberBASIC file
./cyberbasic --run examples/hello.cyber

# Specify output file name
./cyberbasic -o output.go examples/hello.cyber
```

## Command-line Options

- `--run`: Run the program after compilation
- `--no-progress`: Disable progress bar during compilation
- `--no-banner`: Disable ASCII art banner
- `-o filename.go`: Specify output file name

## Running the Example

On Windows, you can use the included batch file to run the example:

```bash
.\run_example.bat
```

## Project Structure

- `ast/`: Abstract Syntax Tree definitions
- `cmd/`: Command-line interfaces
  - `cyberbasic/`: Main compiler command
- `compiler/`: CyberBASIC to Go compiler
- `examples/`: Example CyberBASIC programs
- `lexer/`: Lexical analyzer
- `object/`: Runtime object definitions
- `parser/`: Parser for CyberBASIC
- `build.go`: Cross-platform build script
- `package.go`: Distribution package creator 