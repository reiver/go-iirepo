package iirepo

import (
	"fmt"
	"os"
)

func assertIsDir(path string) (err error) {
	var file *os.File
	{
		file, err = os.Open(path)
		if nil != err {
			return err
		}
		defer func(){
			err = file.Close()
		}()
	}

	var fileinfo os.FileInfo
	{
		fileinfo, err = file.Stat()
	}

	if ! fileinfo.IsDir() {
		return fmt.Errorf("ii: %q is not a directory", path)
	}

	return nil
}
