package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStorageFile(t *testing.T) {
	t.Run("File does not exist, create with default content", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := tempDir + "/test_storage.txt"

		err := checkStorageFile(filePath)
		assert.NoError(t, err)

		// Check if the file was created with default content
		content, err := os.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, defaultExample, string(content))
	})

	t.Run("File already exists, no modification", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := tempDir + "/test_storage.txt"

		// Create a file with some content
		existingContent := []byte("existing content")
		err := os.WriteFile(filePath, existingContent, filePermission)
		assert.NoError(t, err)

		err = checkStorageFile(filePath)
		assert.NoError(t, err)

		// Check if the existing file content is not modified
		content, err := os.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, existingContent, content)
	})

	t.Run("Unable to open file for checking", func(t *testing.T) {
		// Create a directory with the same name as the file
		tempDir := t.TempDir()
		filePath := tempDir + "/test_storage"

		err := os.Mkdir(filePath, 0755)
		assert.NoError(t, err)

		err = checkStorageFile(filePath)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Unable to open storage file")
	})

	t.Run("Invalid file path", func(t *testing.T) {
		filePath := "/non_existent_path/test_storage.txt"

		err := checkStorageFile(filePath)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Unable to create storage file: open /non_existent_path/test_storage.txt: no such file or directory")
	})
}
