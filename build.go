//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var targets = []struct {
	os   string
	arch string
	ext  string
}{
	{"windows", "amd64", ".exe"},
	{"windows", "386", ".exe"},
	{"linux", "amd64", ""},
	{"linux", "386", ""},
	{"darwin", "amd64", ""},
	{"darwin", "arm64", ""},
}

func main() {
	fmt.Println("Building CyberBASIC compiler for multiple platforms...")

	// Create dist directory if it doesn't exist
	if err := os.MkdirAll("dist", 0755); err != nil {
		fmt.Printf("Error creating dist directory: %v\n", err)
		os.Exit(1)
	}

	// Build for each target
	for _, target := range targets {
		fmt.Printf("Building for %s/%s...\n", target.os, target.arch)

		// Set up environment variables for cross-compilation
		env := append(os.Environ(),
			fmt.Sprintf("GOOS=%s", target.os),
			fmt.Sprintf("GOARCH=%s", target.arch),
		)

		// Create output file name with appropriate extension
		outputName := fmt.Sprintf("cyberbasic_%s_%s%s", target.os, target.arch, target.ext)
		outputPath := filepath.Join("dist", outputName)

		// Run go build for the compiler
		cmd := exec.Command("go", "build", "-o", outputPath, "./cmd/cyberbasic")
		cmd.Env = env
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Error building for %s/%s: %v\n", target.os, target.arch, err)
			continue
		}

		fmt.Printf("Successfully built %s\n", outputPath)
	}

	// Copy documentation and examples
	fmt.Println("Copying documentation and examples...")

	// Copy README and LICENSE
	copyFile("README.md", "dist/README.md")
	copyFile("LICENSE", "dist/LICENSE")

	// Create examples directory
	if err := os.MkdirAll("dist/examples", 0755); err != nil {
		fmt.Printf("Error creating examples directory: %v\n", err)
	} else {
		// Copy .cyber example files
		exampleFiles, err := filepath.Glob("examples/*.cyber")
		if err == nil {
			for _, file := range exampleFiles {
				base := filepath.Base(file)
				if err := copyFile(file, filepath.Join("dist/examples", base)); err != nil {
					fmt.Printf("Error copying example %s: %v\n", file, err)
				}
			}
		}
	}

	// Create a simple batch/shell script to run the compiler
	createRunScript()

	fmt.Println("Build complete! Distribution files are in the 'dist' directory.")
	fmt.Println("To compile a .cyber file, run: cyberbasic yourfile.cyber")
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func createRunScript() {
	var script string
	var filename string

	if runtime.GOOS == "windows" {
		script = `@echo off
if "%1"=="" (
    echo Usage: cyberbasic yourfile.cyber
    exit /b 1
)
cyberbasic.exe %*`
		filename = "dist/cyberbasic.bat"
	} else {
		script = `#!/bin/sh
if [ -z "$1" ]; then
    echo "Usage: cyberbasic yourfile.cyber"
    exit 1
fi
./cyberbasic "$@"`
		filename = "dist/cyberbasic.sh"
	}

	err := os.WriteFile(filename, []byte(script), 0755)
	if err != nil {
		fmt.Printf("Error creating run script: %v\n", err)
	}
}
