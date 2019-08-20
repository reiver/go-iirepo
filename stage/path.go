package iirepo_stage

import (
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/logger"

	"path/filepath"
)

func Path(grandparent string) string {

	iirepo_logger.Debugf("iirepo_stage.Path(%q): begin", grandparent)

	repopath := iirepo.Path(grandparent)
	iirepo_logger.Debugf("iirepo_stage.Path(%q): repopath = %q", grandparent, repopath)

	path := filepath.Join(repopath, Name())
	iirepo_logger.Debugf("iirepo_stage.Path(%q): path = %q", grandparent, path)

	iirepo_logger.Debugf("iirepo_stage.Path(%q): end", grandparent)

	return path
}
