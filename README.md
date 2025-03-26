# CyberBASIC

CyberBASIC is a modern BASIC programming language implementation with support for game development, graphics, and multimedia.

## Features

- Traditional BASIC syntax with modern features
- Built-in graphics and game development capabilities
- Cross-platform support (Windows, Linux, macOS)
- Object-oriented programming support
- JSON and file handling
- Audio and input support

## Installation

1. Download the appropriate version for your platform from the releases page:
   - Windows: `cyberbasic_windows_amd64.exe`
   - Linux: `cyberbasic_linux_amd64`
   - macOS: `cyberbasic_darwin_amd64`

2. Add the compiler to your system PATH or place it in your working directory.

## Usage

Create a `.cyber` file with your BASIC code:

```basic
' hello.cyber
PRINT "Hello, World!"

FOR i = 1 TO 5
    PRINT "Count: "; i
NEXT i
```

Compile and run your program:

```bash
# Windows
cyberbasic.exe hello.cyber

# Linux/macOS
./cyberbasic hello.cyber
```

## Example Programs

Check the `examples` directory for sample programs demonstrating various features:

- `hello.cyber`: Basic syntax and control structures
- `simple_game_demo.cyber`: Game development features
- `simple_json.cyber`: JSON handling

## Language Syntax

### Basic Statements
```basic
PRINT "Hello"
LET x = 42
INPUT "Enter name: ", name$
```

### Control Structures
```basic
IF x > 0 THEN
    PRINT "Positive"
ELSE
    PRINT "Non-positive"
END IF

FOR i = 1 TO 10
    PRINT i
NEXT i

WHILE x > 0
    x = x - 1
WEND
```

### Graphics and Game Functions
```basic
SCREEN 800, 600
CLEARSCREEN
DRAWRECT 100, 100, 50, 50
DRAWSPRITE "player.png", 200, 200
```

### Object-Oriented Features
```basic
CLASS Player
    LET x = 0
    LET y = 0
    
    FUNCTION move(dx, dy)
        x = x + dx
        y = y + dy
    END FUNCTION
END CLASS
```

## Building from Source

Requires Go 1.21 or later.

```bash
go run build.go
```

The compiled binaries will be available in the `dist` directory.

## License

MIT License - See LICENSE file for details. 