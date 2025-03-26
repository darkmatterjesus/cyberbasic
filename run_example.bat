@echo off
echo Compiling and running CyberBASIC example...

REM Build the compiler if it doesn't exist
if not exist cyberbasic.exe (
    echo Building CyberBASIC compiler...
    go build -o cyberbasic.exe cmd/cyberbasic/main.go
    if %ERRORLEVEL% neq 0 (
        echo Compiler build failed!
        exit /b %ERRORLEVEL%
    )
)

REM Compile and run the example
echo.
echo Compiling and running examples/full_example.cyber with progress bar:
.\cyberbasic.exe examples/full_example.cyber --run
if %ERRORLEVEL% neq 0 (
    echo Compilation or execution failed!
    exit /b %ERRORLEVEL%
)

echo.
echo Example compiled and executed successfully.
echo.
echo Available flags:
echo  --run: Compile and run the program
echo  --no-progress: Disable progress bar during compilation
echo  --no-banner: Disable ASCII art banner
echo  --o filename.go: Specify output filename 