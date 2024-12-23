package analysis

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tphakala/birdnet-go/internal/conf"
)

// DirectoryAnalysis processes all .wav files in the given directory for analysis.
func DirectoryAnalysis(settings *conf.Settings) error {
	// Initialize BirdNET interpreter
	if err := initializeBirdNET(settings); err != nil {
		return err
	}

	analyzeFunc := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Return the error to stop the walking process.
			return err
		}

		if d.IsDir() {
			// If recursion is not enabled and this is a subdirectory, skip it.
			if !settings.Input.Recursive && path != settings.Input.Path {
				return filepath.SkipDir
			}
			// If it's the root directory or recursion is enabled, continue walking.
			return nil
		}

		// Check for both .wav and .flac files
		if strings.HasSuffix(strings.ToLower(d.Name()), ".wav") ||
			strings.HasSuffix(strings.ToLower(d.Name()), ".flac") {

			settings.Input.Path = path
			if err := FileAnalysis(settings); err != nil {
				// If FileAnalysis returns an error log it and continue
				log.Printf("Error analyzing file '%s': %v", path, err)
				return nil
			}
		}
		return nil
	}

	// Start walking through the directory
	err := filepath.WalkDir(settings.Input.Path, analyzeFunc)
	if err != nil {
		log.Fatalf("Failed to walk directory: %v", err)
	}

	return nil
}
