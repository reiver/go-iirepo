package iirepo

import (
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

		repopath = Path(x)

		isADir, err := isDir(repopath)
		if nil == err && isADir {
			return x, nil
		}

		parentdir := filepath.Dir(x)
		if x == parentdir {
			break
		}
		x = parentdir
	}

	return "", fmt.Errorf("iirepo: repo not found for %s", path)
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
	repopath, err := Locate(path)
	if nil != err {
		return "", err
	}

	rootpath = filepath.Dir(repopath)

	return rootpath, nil
}
