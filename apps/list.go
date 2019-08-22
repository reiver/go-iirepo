package iirepo_apps

import (
	"github.com/reiver/go-iirepo/logger"

	"os"
	"path/filepath"
)

func List(path string) ([][]string, error) {

	iirepo_logger.Debugf("iirepo_apps.List(%q): begin", path)

	appspath, err := Locate(path)
	if nil != err {
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

	iirepo_logger.Debugf("iirepo_apps.List(%q): end", path)

	return apps, nil
}
