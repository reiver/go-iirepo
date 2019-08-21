package iirepo_stage

import (
	"github.com/reiver/go-iirepo/logger"

	"fmt"
	"io"
	"os"
	"path/filepath"
)

func storewhere(path string) (srcpath string, dstpath string, dstpathdir string, err error) {

	iirepo_logger.Debugf("iirepo_stage.storewhere(%q): begin", path)

	//var srcpath string
	{
		var err error

		srcpath, err = filepath.Abs(path)
		if nil != err {
			return "", "", "", err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.storewhere(%q): srcpath = %q", path, srcpath)

	//var dstpath string
	{
		var err error

		dstpath, err = StagedPath(path)
		if nil != err {
			return "", "", "", err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.storewhere(%q): dstpath = %q", path, dstpath)

	//var dstpathdir string
	{
		dstpathdir = filepath.Dir(dstpath)
	}
	iirepo_logger.Debugf("iirepo_stage.storewhere(%q): dstpathdir = %q", path, dstpathdir)

	iirepo_logger.Debugf("iirepo_stage.storewhere(%q): end", path)

	return srcpath, dstpath, dstpathdir, nil
}

// Store moves the file at ‘path’ to the stage.
//
// See also: StoreCopy
func StoreOriginal(path string) (err error) {

	iirepo_logger.Debugf("iirepo_stage.StoreOriginal(%q): begin", path)

	var srcpath    string
	var dstpath    string
	var dstpathdir string
	{
		var err error

		srcpath, dstpath, dstpathdir, err = storewhere(path)
		if nil != err {
			return err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.StoreOriginal(%q): srcpath = %q", path, srcpath)
	iirepo_logger.Debugf("iirepo_stage.StoreOriginal(%q): dstpath = %q", path, dstpath)
	iirepo_logger.Debugf("iirepo_stage.StoreOriginal(%q): dstpathdir = %q", path, dstpathdir)

	if err := os.MkdirAll(dstpathdir, os.ModePerm); nil != err {
		return err
	}

	if err := os.Rename(srcpath, dstpath); nil != err {
		return err
	}

	iirepo_logger.Debugf("iirepo_stage.StoreOriginal(%q): end", path)

	return nil
}

// Store copies the file at ‘path’ to the stage.
//
// See also: StoreOriginal
func StoreCopy(path string) (err error) {

	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): begin", path)

	var srcpath    string
	var dstpath    string
	var dstpathdir string
	{
		var err error

		srcpath, dstpath, dstpathdir, err = storewhere(path)
		if nil != err {
			return err
		}
	}
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): srcpath = %q", path, srcpath)
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): dstpath = %q", path, dstpath)
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): dstpathdir = %q", path, dstpathdir)

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
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): src file opened", path)

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
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): dst file opened", path)

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
	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): copy from src to dst completed", path)

//@TODO: add some checks to make sure file was copied OK.

	iirepo_logger.Debugf("iirepo_stage.StoreCopy(%q): end", path)

	return nil
}
