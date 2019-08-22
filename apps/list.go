package iirepo_apps

import (
	"github.com/reiver/go-iirepo/logger"

	"os"
	"path/filepath"
	"syscall"
)

func List(path string) ([][]string, error) {

	iirepo_logger.Debugf("iirepo_apps.List(%q): begin", path)
	defer iirepo_logger.Debugf("iirepo_apps.List(%q): end", path)

	appspath, err := Locate(path)
	if nil != err {
		switch patherror := err.(type) {
		case *os.PathError:
			switch errno := patherror.Err.(type) {
			case syscall.Errno:
				if syscall.ENOENT == errno {
					iirepo_logger.Debugf("iirepo_apps.List(%q): repo exists, but %s/ not created yet (therefore no apps) (note: this is not an error)", path, Name())
					return nil, nil
				}
			}
		}

		return nil, err
	}
	iirepo_logger.Debugf("iirepo_apps.List(%q): appspath = %q", path, appspath)

	var apps [][]string

	err = filepath.Walk(appspath, func(apppath string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}

		appName := filepath.Base(apppath)
		iirepo_logger.Debugf("iirepo_apps.List(%q): app-name = %q", path, appName)

		app := []string{appName}

		apps = append(apps, app)

		return nil
	})
	if nil != err {
		return nil, err
	}

	return apps, nil
}
