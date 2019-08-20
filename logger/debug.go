package iirepo_logger

import (
	"github.com/reiver/go-tmpl"

	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

var (
	debugWriter io.Writer
)

func init() {
	debugWriter = ioutil.Discard
}

func Debug(v ...interface{}) {
	var builder strings.Builder

	fmt.Fprint(&builder, v...)
	builder.WriteRune('\n')

	io.WriteString(debugWriter, builder.String())
}

func Debugf(format string, v ...interface{}) {
	var builder strings.Builder

	fmt.Fprintf(&builder, format, v...)
	builder.WriteRune('\n')

	io.WriteString(debugWriter, builder.String())
}

func Debugt(template string, data interface{}) {
	var builder strings.Builder

	tmpl.Fprintt(&builder, template, data)
	builder.WriteRune('\n')

	io.WriteString(debugWriter, builder.String())
}
