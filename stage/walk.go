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

	return filepath.Walk(stagepath, walker{fn}.WalkFunc)
}

type walker struct {
	Func func(relstagedpath string)error
}

func (receiver walker) WalkFunc(path string, info os.FileInfo, err error) error {
	if nil != err {
		return err
	}

	return receiver.Func(path)
}
