package iirepo

import (
	"path/filepath"
)

// Path returns the repo path.
//
// Note that Path does NOT create this directory itself.
func Path(rootpath string) string {
	path := filepath.Join(rootpath, Name())

	return path
}
