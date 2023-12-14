package config

import (
	"os"
	"path"
	"testing"
)

// TestParse is a unit test for the Parse method.
func TestParse(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a temporary storage file for testing
	tempStoragePath := path.Join(tempDir, "test_storage.txt")
	err := os.WriteFile(tempStoragePath, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary storage file: %v", err)
	}

	// Initialize Config with test values
	cfg := &Config{
		Storage: Storage{
			Directory: tempDir,
			Name:      "test_storage.txt",
			Path:      tempStoragePath,
		},
		Server: Server{
			Web: Web{
				Address: "localhost:8080",
			},
		},
	}

	// Run the Parse method
	err = cfg.Parse()

	// Check if there was an error
	if err != nil {
		t.Fatalf("Parse method returned an error: %v", err)
	}

	// Add additional assertions as needed to validate the state of the Config struct
	// For example, you can check if the working directory, storage path, content, etc., are set correctly.

	// Example assertion: Check if the working directory is set correctly
	if cfg.Storage.Directory != tempDir {
		t.Errorf("Working directory not set correctly. Expected: %s, Got: %s", tempDir, cfg.Storage.Directory)
	}

	// Example assertion: Check if the storage path is set correctly
	if cfg.Storage.Path != tempStoragePath {
		t.Errorf("Storage path not set correctly. Expected: %s, Got: %s", tempStoragePath, cfg.Storage.Path)
	}
}
