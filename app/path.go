package iirepo_app

import (
	"github.com/reiver/go-iirepo"

	"path/filepath"
)

func Path(rootpath string) string {
	repopath := iirepo.Path(rootpath)

	path := filepath.Join(repopath, Name())

	return path
}
