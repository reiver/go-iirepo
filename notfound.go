package iirepo

import (
	"fmt"
)

// NotFound is an error that is returned by Locate(), and LocateRoot() if they cannot locate
// what they are looking for.
//
// Here is an example of with iirepo.LocateRoot():
//
//	import "github.com/reiver/go-iirepo"
//		
//	// ...
//		
//	rootpath, err := iirepo.LocateRoot()
//	if nil != err {
//		switch casted := err.(type) {
//		case iirepo.NotFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
//
// Here is an example of with iirepo.Locate():
//
//	import "github.com/reiver/go-iirepo"
//		
//	// ...
//		
//	rootpath, err := iirepo.Locate()
//	if nil != err {
//		switch casted := err.(type) {
//		case iirepo.NotFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type NotFound interface {
	error
	NotFound() string
}

func errNotFound(msg string) error {
	var e NotFound = &internalNotFound{
		msg: msg,
	}

	return e
}

func errNotFoundf(format string, a ...interface{}) error {
	msg := fmt.Sprintf(format, a...)
	return errNotFound(msg)
}

type internalNotFound struct {
	msg string
}

func (receiver internalNotFound) Error() string {
	return fmt.Sprintf("iirepo: Not Found: %s", receiver.msg)
}

func (receiver internalNotFound) NotFound() string {
	return receiver.msg
}
