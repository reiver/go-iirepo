package iirepo

import (
	"path/filepath"
)

// Path returns the repo path.
//
// Note that Path does NOT create this directory itself.
func Path(parent string) string {
	path := filepath.Join(parent, Name())

	return path
}
