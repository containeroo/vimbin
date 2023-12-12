package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

// filePermission represents the default file permission used in the application.
const filePermission = 0644

// defaultExample is the default content example used when creating the storage file.
const defaultExample = `
#include "syscalls.h"
/* getchar:  simple buffered version */
int getchar(void)
{
  static char buf[BUFSIZ];
  static char *bufp = buf;
  static int n = 0;
  if (n == 0) {  /* buffer is empty */
  n = read(0, buf, sizeof buf);
  bufp = buf;
}
return (--n >= 0) ? (unsigned char) *bufp++ : EOF;
}
`

// checkStorageFile checks if the storage file exists; if not, it creates it with default content.
//
// Parameters:
//   - filePath: string
//     The path to the storage file.
//
// Returns:
//   - error
//     An error if unable to check or create the storage file.
func checkStorageFile(filePath string) error {
	// Open the file to check if it exists
	_, err := os.Stat(filePath)
	if err != nil {
		// If the file doesn't exist, create it with default content
		log.Debug().Msg("Storage file not found; creating it with default content")
		example := []byte(defaultExample)

		if err := os.WriteFile(filePath, example, filePermission); err != nil {
			return fmt.Errorf("Unable to create storage file: %s", err)
		}
	}

	// Try to open file
	f, err := os.OpenFile(filePath, os.O_RDWR, filePermission)
	if err != nil {
		return fmt.Errorf("Unable to open storage file: %s", err)
	}
	defer f.Close()

	return nil
}
