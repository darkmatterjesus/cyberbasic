@echo off
echo Building CyberBASIC compiler...

REM Build for the current platform first
go build -o cyberbasic.exe cmd/cyberbasic/main.go
if %ERRORLEVEL% neq 0 (
    echo Build failed!
    exit /b %ERRORLEVEL%
)

echo.
echo CyberBASIC compiler built successfully.
echo Run "cyberbasic.exe examples/hello.cyber" to test it.
echo.
echo To build for all platforms, run:
echo    go run build.go
echo.
echo To create distribution packages, run:
echo    go run build.go
echo    go run package.go 