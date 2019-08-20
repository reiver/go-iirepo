package iirepo_stage

import (
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/logger"

	"fmt"
	"path/filepath"
	"strings"
)

// StagedPath returns the path of where ‘path’ would be put under the stage.
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
//	location, err := iirepo_stage.StagedPath("/home/joeblow/workspaces/myproject/numbers/one.go")
//
// Would return a value for ‘location’ of:
//
//	"/home/joeblow/workspaces/myproject/.ii/stage/numbers/one.go"
func StagedPath(path string) (stagedpath string, err error) {

	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): begin", path)

	var srcpath string
	{
		var err error

		srcpath, err = filepath.Abs(path)
		if nil != err {
			return "", err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): srcpath = %q", path, srcpath)

	var stagepath string
	{
		var err error

		stagepath, err = Locate(srcpath)
		if nil != err {
			return "", err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): stagepath = %q", path, stagepath)

	var rootpath string
	{
		var err error

		rootpath, err = iirepo.LocateRoot(srcpath)
		if nil != err {
			return "", err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): rootpath = %q", path, rootpath)

	var rootpathslash string
	{
		var builder strings.Builder

		builder.WriteString(rootpath)
		builder.WriteRune(filepath.Separator)

		rootpathslash = builder.String()
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): rootpathslash = %q", path, rootpathslash)

	var relsrcpath string
	{
		if !strings.HasPrefix(srcpath, rootpathslash) {
			return "", fmt.Errorf("iirepo: %q is not under repo for %q", srcpath, rootpathslash)
		}

		relsrcpath = srcpath[len(rootpathslash):]
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): relsrcpath = %q", path, relsrcpath)

	//var stagedpath string
	{
		stagedpath = filepath.Join(stagepath, relsrcpath)
	}
	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): stagedpath = %q", path, stagedpath)

	iirepo_logger.Debugf("iirepo_stage.StagedPath(%q): end", path)

	return stagedpath, nil
}
