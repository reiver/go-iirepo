package iirepo_contents

import (
	"github.com/reiver/go-iirepo"

	"path/filepath"
)

func Path(grandparent string) string {
	repopath := iirepo.Path(grandparent)

	path := filepath.Join(repopath, Name())

	return path
}
