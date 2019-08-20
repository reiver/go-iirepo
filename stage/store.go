package iirepo_stage

import (
	"github.com/reiver/go-iirepo/logger"

	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Store copied the file at ‘path’ to the stage.
func Store(path string) (err error) {

	iirepo_logger.Debugf("iirepo_stage.Store(%q): begin", path)

	var srcpath string
	{
		var err error

		srcpath, err = filepath.Abs(path)
		if nil != err {
			return err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): srcpath = %q", path, srcpath)

	var dstpath string
	{
		var err error

		dstpath, err = StagedPath(path)
		if nil != err {
			return err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): dstpath = %q", path, dstpath)

	var dstpathdir string
	{
		dstpathdir = filepath.Dir(dstpath)
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): dstpathdir = %q", path, dstpathdir)

	var src *os.File
	{
		var err error

		src, err = os.Open(srcpath)
		if nil != err {
			return err
		}
		defer func() {
			e := src.Close()
			if nil == err {
				err = e
			}
		}()

		fileinfo, err := src.Stat()
		if nil != err {
			return err
		}
		if ! fileinfo.Mode().IsRegular() {
			return fmt.Errorf("iirepo: %q is not a regular file", srcpath)
		}
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): src file opened", path)

	var dst *os.File
	{
		if err := os.MkdirAll(dstpathdir, os.ModePerm); nil != err {
			return err
		}

		var err error

		dst, err = os.Create(dstpath)
		if nil != err {
			return err
		}
		defer func() {
			e := dst.Close()
			if nil == err {
				err = e
			}
		}()
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): dst file opened", path)

	{
		var buffer [64*8]byte // The 64 comes from the typical length of a ‘cache line’.

		for {
			numRead, err := src.Read(buffer[:])
			if nil != err && io.EOF != err {
				return err
			}
			if 0 == numRead {
				break
			}

			numWritten, err :=  dst.Write(buffer[:numRead])
			if nil != err {
				return err
			}

			if expected, actual := numRead, numWritten; expected != actual {
				return fmt.Errorf("iirepo: Short Write: expected=%d actual=%d", expected, actual)
			}
		}
	}
	iirepo_logger.Debugf("iirepo_stage.Store(%q): copy from src to dst completed", path)

//@TODO: add some checks to make sure file was copied OK.

	iirepo_logger.Debugf("iirepo_stage.Store(%q): end", path)

	return nil
}
