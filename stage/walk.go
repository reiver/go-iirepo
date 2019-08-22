package iirepo_stage

import (
	"os"
	"path/filepath"
)

// Walk will locate the ii repo stage for ‘path’ and call ‘fn’ for each staged file in deterministic order
// based on lexical ordering.
func Walk(path string, fn func(relstagedpath string)error) error {

	stagepath, err := Locate(path)
	if nil != err {
		return err
	}

	return filepath.Walk(stagepath, walker{stagepath, fn}.WalkFunc)
}

type walker struct {
	BasePath string
	Func func(relstagedpath string)error
}

func (receiver walker) WalkFunc(path string, info os.FileInfo, err error) error {
	if nil != err {
		return err
	}

	// Skip "."
	if "." == path {
		return nil
	}

	relpath, err := filepath.Rel(receiver.BasePath, path)
	if nil != err {
		return err
	}

	return receiver.Func(relpath)
}
