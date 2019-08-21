package iirepo_stage

import (
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/logger"

	"path/filepath"
)

func Path(rootpath string) string {

	iirepo_logger.Debugf("iirepo_stage.Path(%q): begin", rootpath)

	repopath := iirepo.Path(rootpath)
	iirepo_logger.Debugf("iirepo_stage.Path(%q): repopath = %q", rootpath, repopath)

	path := filepath.Join(repopath, Name())
	iirepo_logger.Debugf("iirepo_stage.Path(%q): path = %q", rootpath, path)

	iirepo_logger.Debugf("iirepo_stage.Path(%q): end", rootpath)

	return path
}
