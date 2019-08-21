package iirepo

import (
	"github.com/reiver/go-iirepo/logger"

	"fmt"
	"path/filepath"
)

// Locate returns the path to the repo for ‘path’.
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
//	location, err := iirepo.Locate("/home/joeblow/workspaces/myproject/numbers")
//
// Would return a value for ‘location’ of:
//
//	"/home/joeblow/workspaces/myproject/.ii"
func Locate(path string) (repopath string, err error) {

	iirepo_logger.Debugf("iirepo.Locate(%q): begin", path)

	var rootpath string
	{
		var err error

		rootpath, err = LocateRoot(path)
		if nil != err {
			return "", err
		}
	}
	iirepo_logger.Debugf("iirepo.Locate(%q): rootpath = %q", path, rootpath)

	repopath = filepath.Join(rootpath, Name())
	iirepo_logger.Debugf("iirepo.Locate(%q): repopath = %q", path, repopath)

	iirepo_logger.Debugf("iirepo.Locate(%q): end", path)

	return rootpath, nil
}

// LocateRoot returns the path to the root for ‘path’.
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
//	location, err := iirepo.LocateRoot("/home/joeblow/workspaces/myproject/numbers")
//
// Would return a value for ‘location’ of:
//
//	"/home/joeblow/workspaces/myproject"
func LocateRoot(path string) (rootpath string, err error) {

	iirepo_logger.Debugf("iirepo.LocateRoot(%q): begin", path)

	x := path

	{
		isADir, err := isDir(path)
		if nil != err {
			return "", err
		}

		if !isADir {
			x = filepath.Dir(x)
		}
	}

	for "" != x {
		if isADir, err := isDir(x); nil != err {
			return "", err
		} else if !isADir {
			return "", fmt.Errorf("iirepo: %s is not a directory", x)
		}

		repopath := Path(x)

		isADir, err := isDir(repopath)
		if nil == err && isADir {
			break
		}

		parentdir := filepath.Dir(x)
		if x == parentdir {
			return "", errNotFoundf("%s repo not found for %s", Name(), path)
		}
		x = parentdir
	}
	iirepo_logger.Debugf("iirepo.LocateRoot(%q): found root = %q", path, x)


	iirepo_logger.Debugf("iirepo.LocateRoot(%q): end", path)

	return x, nil
}
