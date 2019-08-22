package iirepo_apps

import (
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/logger"
)

// Locate returns the path to the repo's apps for ‘path’.
//
// For example, if name of the repo directory is “.ii/”, and the directory structure looks like:
//
//	/home/joeblow/workspaces/
//	/home/joeblow/workspaces/myproject/
//	/home/joeblow/workspaces/myproject/.ii    <--------- Note the repo directory here.
//	/home/joeblow/workspaces/myproject/apple.go
//	/home/joeblow/workspaces/myproject/banana.go
//	/home/joeblow/workspaces/myproject/cherry.go
//	/home/joeblow/workspaces/myproject/date.go
//	/home/joeblow/workspaces/myproject/numbers/
//	/home/joeblow/workspaces/myproject/numbers/one.go
//	/home/joeblow/workspaces/myproject/numbers/two.go
//	/home/joeblow/workspaces/myproject/numbers/three.go
//	/home/joeblow/workspaces/myproject/emoji/
//	/home/joeblow/workspaces/myproject/emoji/grinning.go
//	/home/joeblow/workspaces/myproject/emoji/slightlysmiling.go
//
// Then calling:
//
//	location, err := iirepo_apps.Locate("/home/joeblow/workspaces/myproject/numbers")
//
// Would return a value for ‘location’ of:
//
//	"/home/joeblow/workspaces/myproject/.ii/apps"
func Locate(path string) (appspath string, err error) {
	return locate(path, iirepo.LocateRoot)
}

func locate(path string, locateRootFunc func(string)(string,error)) (appspath string, err error) {
	iirepo_logger.Debugf("iirepo_apps.Locate(%q): begin", path)

	rootpath, err := locateRootFunc(path)
	if nil != err {
		return "", err
	}
	iirepo_logger.Debugf("iirepo_apps.Locate(%q): rootpath = %q", path, rootpath)

	appspath = Path(rootpath)
	iirepo_logger.Debugf("iirepo_apps.Locate(%q): appspath = %q", path, appspath)

	iirepo_logger.Debugf("iirepo_apps.Locate(%q): end", path)

	return appspath, nil
}
