// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config defines the structure of the JSON configuration file.
type Config struct {
	HeaderLines         []string          `json:"headerLines"`         // Expected file header.
	FileCommentPrefixes map[string]string `json:"fileCommentPrefixes"` // Comment styles for file types.
	IgnoredPaths        []string          `json:"ignoredPaths"`        // Paths to ignore.
	MaxScanLines        int               `json:"maxScanLines"`        // Number of lines to check for the header.
}

// Global variables to hold the configuration and results.
var config Config
var filesWithIssues []string
var filesWithError []string
var processedCount = 0

var providedConfigPath = flag.String("config", "", "Path to the configuration file")

// main is the entry point for the header check.
func main() {
	flag.Parse()

	if *providedConfigPath == "" {
		fmt.Println("Provide a JSON configuration using the --config flag.")
		os.Exit(1)
	}

	fmt.Printf("Configuration file: %s\n", *providedConfigPath)
	configPathToLoad := *providedConfigPath

	// Load the configuration file.
	if err := loadConfig(configPathToLoad); err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1) // Configuration error is critical.
	}

	// Start directory traversal.
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories.
		if info.IsDir() {
			return nil
		}

		// Skip ignored paths.
		if isIgnored(path) {
			return nil
		}

		// Check for file types supported by `fileCommentPrefixes`.
		ext := filepath.Ext(path)
		if commentPrefix, ok := config.FileCommentPrefixes[ext]; ok {
			processedCount++
			currentHeader := transformHeader(config.HeaderLines, commentPrefix)

			// Check the file for the required header.
			if missing, err := checkHeader(path, currentHeader); err != nil {
				filesWithError = append(filesWithError, fmt.Sprintf("%s (error: %v)", path, err))
			} else if missing {
				filesWithIssues = append(filesWithIssues, path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("❌ Error during directory traversal: %v\n", err)
		os.Exit(2) // Critical error during file traversal.
	}

	// Print the summary and get the exit code.
	exitCode := printSummary()
	os.Exit(exitCode)
}

// loadConfig loads the JSON configuration into the global config variable.
func loadConfig(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open configuration file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("failed to decode configuration: %w", err)
	}

	return nil
}

// isIgnored determines if a file or directory should be skipped based on ignoredPaths.
func isIgnored(path string) bool {
	normalizedPath := strings.ReplaceAll(path, string(filepath.Separator), "/")
	for _, pattern := range config.IgnoredPaths {
		if matched, _ := filepath.Match(pattern, normalizedPath); matched {
			return true
		}
		if strings.Contains(pattern, "**") {
			prefix := strings.Split(pattern, "**")[0]
			if strings.HasPrefix(normalizedPath, prefix) {
				return true
			}
		}
	}
	return false
}

// transformHeader adjusts the header format for the target file's comment style.
func transformHeader(header []string, prefix string) []string {
	transformed := make([]string, len(header))
	for i, line := range header {
		transformed[i] = prefix + " " + strings.TrimPrefix(line, "// ")
	}
	return transformed
}

// checkHeader verifies if the file contains the required header.
func checkHeader(path string, headerLines []string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	skipShebang := filepath.Ext(path) == ".sh" // Handle shebang for shell scripts.

	// Only scan the first `MaxScanLines` lines.
	for i := 0; i < config.MaxScanLines && scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())

		// Skip shebang if present.
		if skipShebang && strings.HasPrefix(line, "#!") {
			skipShebang = false
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return !containsHeader(lines, headerLines), nil
}

// containsHeader checks if the required header lines exist in a file.
func containsHeader(fileLines, headerLines []string) bool {
	i := 0
	for _, line := range fileLines {
		if line == headerLines[i] {
			i++
			if i == len(headerLines) {
				return true
			}
		}
	}
	return false
}

// printSummary displays the results after processing all files and returns the exit code.
func printSummary() int {
	fmt.Println("Processing complete.")
	fmt.Printf("Total files processed: %d\n", processedCount)

	exitCode := 0

	if len(filesWithIssues) > 0 {
		fmt.Printf("Missing headers: %d\n", len(filesWithIssues))
		for _, file := range filesWithIssues {
			fmt.Printf("   - %s\n", file)
		}

		exitCode = 1 // Missing headers found.
	}

	if len(filesWithError) > 0 {
		fmt.Printf("Errors encountered in files: %d\n", len(filesWithError))
		for _, file := range filesWithError {
			fmt.Printf("   - %s\n", file)
		}

		if exitCode == 1 {
			exitCode = 3 // Both missing headers and errors.
		} else {
			exitCode = 2 // Only errors occurred.
		}
	}

	if len(filesWithIssues) == 0 && len(filesWithError) == 0 {
		fmt.Println("All processed files have the expected header.")
	}

	return exitCode
}
