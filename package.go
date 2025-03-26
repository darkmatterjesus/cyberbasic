//go:build ignore
// +build ignore

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Creating distribution packages...")

	// Get all files in the dist directory
	files, err := os.ReadDir("dist")
	if err != nil {
		fmt.Printf("Error reading dist directory: %v\n", err)
		os.Exit(1)
	}

	// Group files by platform
	platforms := make(map[string][]string)
	commonFiles := []string{"README.md"}

	// Add example files to common files
	exampleFiles, err := filepath.Glob("dist/examples/*")
	if err == nil {
		for _, file := range exampleFiles {
			commonFiles = append(commonFiles, file)
		}
	}

	for _, file := range files {
		if file.IsDir() && file.Name() == "examples" {
			continue // Skip the examples directory itself
		}

		name := file.Name()
		if strings.HasPrefix(name, "cyberbasic_") {
			parts := strings.Split(name, "_")
			if len(parts) >= 3 {
				platform := fmt.Sprintf("%s_%s", parts[1], parts[2])
				if strings.HasSuffix(platform, ".exe") {
					platform = platform[:len(platform)-4]
				}
				platforms[platform] = append(platforms[platform], filepath.Join("dist", name))
			}
		}
	}

	// Create package directory if it doesn't exist
	if err := os.MkdirAll("packages", 0755); err != nil {
		fmt.Printf("Error creating packages directory: %v\n", err)
		os.Exit(1)
	}

	// Create a zip file for each platform
	for platform, binaries := range platforms {
		zipFile := filepath.Join("packages", fmt.Sprintf("cyberbasic_%s.zip", platform))
		fmt.Printf("Creating package for %s: %s\n", platform, zipFile)

		// Create a new zip file
		archive, err := os.Create(zipFile)
		if err != nil {
			fmt.Printf("Error creating zip file: %v\n", err)
			continue
		}
		defer archive.Close()

		zipWriter := zip.NewWriter(archive)
		defer zipWriter.Close()

		// Add binary files
		for _, file := range binaries {
			if err := addFileToZip(zipWriter, file, filepath.Base(file)); err != nil {
				fmt.Printf("Error adding file to zip: %v\n", err)
			}
		}

		// Add common files
		for _, file := range commonFiles {
			targetPath := file
			if strings.HasPrefix(file, "dist/") {
				targetPath = file[5:] // Remove the 'dist/' prefix
			}
			if err := addFileToZip(zipWriter, file, targetPath); err != nil {
				fmt.Printf("Error adding file to zip: %v\n", err)
			}
		}
	}

	fmt.Println("Package creation complete! Distribution packages are in the 'packages' directory.")
}

func addFileToZip(zipWriter *zip.Writer, source, target string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = target
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
