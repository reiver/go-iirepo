package iirepo

import (
	"os"
)

func isDir(path string) (tis bool, err error) {
	var file *os.File
	{
		file, err = os.Open(path)
		if nil != err {
			return false, err
		}
		defer func(){
			e := file.Close()
			if nil == err {
				err = e
			}
		}()
	}

	var fileinfo os.FileInfo
	{
		fileinfo, err = file.Stat()
		if nil != err {
			return false, err
		}
	}

	if ! fileinfo.IsDir() {
		return false, nil
	}

	return true, nil
}
