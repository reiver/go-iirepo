package iirepo_contents

import (
	"github.com/reiver/go-iirepo/logger"
	"github.com/reiver/go-iirepo/stage"

	"os"
	"path/filepath"
)

func CommitStaged() error {

	iirepo_logger.Debug("iirepo_contents.CommitStaged(): begin")

	wd, err := os.Getwd()
	if nil != err {
		return err
	}
	iirepo_logger.Debugf("iirepo_contents.CommitStaged(): wd = %q", wd)

	stagepath, err := iirepo_stage.Locate(wd)
	if nil != err {
		return err
	}
	iirepo_logger.Debugf("iirepo_contents.CommitStaged(): stagepath = %q", stagepath)

	err = filepath.Walk(stagepath, func(path string, fileinfo os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		if fileinfo.IsDir() {
			return nil
		}

		iirepo_logger.Debugf("iirepo_contents.CommitStaged(): walked path = %q", path)

//@TODO

		return nil
	})
	if nil != err {
		return err
	}

	iirepo_logger.Debug("iirepo_contents.CommitStaged(): end")

	return nil
}
