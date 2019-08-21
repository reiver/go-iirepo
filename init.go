package iirepo

import (
	"os"
)

// Init creates the (hidden) .ii/ repo directory, if it doesn't exist, under ‘rootpath’.
func Init(rootpath string) error {
	repopath := Path(rootpath)

	if err := os.MkdirAll(repopath, os.ModePerm); nil != err {
		return err
	}

	return nil
}
