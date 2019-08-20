package iirepo_stage

import (
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/logger"
)

// Locate returns the path to the repo's stage for ‘path’.
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
//	location, err := iirepo_stage.Locate("/home/joeblow/workspaces/myproject/numbers")
//
// Would return a value for ‘location’ of:
//
//	"/home/joeblow/workspaces/myproject/.ii/stage"
func Locate(path string) (stagepath string, err error) {
	return locate(path, iirepo.LocateRoot)
}

func locate(path string, locateRootFunc func(string)(string,error)) (stagepath string, err error) {
	iirepo_logger.Debugf("iirepo_stage.Locate(%q): begin", path)

	repopath, err := locateRootFunc(path)
	if nil != err {
		return "", err
	}
	iirepo_logger.Debugf("iirepo_stage.Locate(%q): repopath = %q", path, repopath)

	stagepath = Path(repopath)
	iirepo_logger.Debugf("iirepo_stage.Locate(%q): stagepath = %q", path, stagepath)

	iirepo_logger.Debugf("iirepo_stage.Locate(%q): end", path)

	return stagepath, nil
}
