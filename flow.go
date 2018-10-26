package fmtc

import (
	"bytes"
	"fmt"
	"io"
)

const (
	esc   = "\033["
	clear = esc + "0m"
)

type IColorFlow interface {
	String() string
	Print(str interface{}, colors ... Color) IColorFlow
	PrintLn(str interface{}, colors ... Color) IColorFlow
	Printf(format string, args Args, colors ...Color) IColorFlow
}

type colorFlow struct {
	customWriter bool
	writer       io.Writer
}

// Args is an alias of []interface{}
type Args []interface{}

func New(writer ...io.Writer) IColorFlow {
	ret := &colorFlow{}

	if len(writer) > 0 && writer[0] != nil {
		ret.writer = writer[0]
		ret.customWriter = true
	} else {
		ret.writer = &bytes.Buffer{}
	}

	return ret
}

func Fprint(w io.Writer, str interface{}, colors ... Color) {
	doColors(colors, w)
	fmt.Fprint(w, str, clear)
}

func Fprintln(w io.Writer, str interface{}, colors ... Color) {
	doColors(colors, w)
	fmt.Fprintln(w, str, clear)
}

func Fprintf(w io.Writer, format string, args Args, colors ...Color) {
	doColors(colors, w)
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, clear)
}

func (f *colorFlow) String() string {
	if f.customWriter {
		return ""
	}
	buf := f.writer.(*bytes.Buffer)
	return buf.String()
}

func (f *colorFlow) Print(str interface{}, colors ... Color) IColorFlow {
	Fprint(f.writer, str, colors...)
	return f
}

func (f *colorFlow) PrintLn(str interface{}, colors ... Color) IColorFlow {
	Fprintln(f.writer, str, colors...)
	return f
}

func (f *colorFlow) Printf(format string, args Args, colors ...Color) IColorFlow {
	Fprintf(f.writer, format, args, colors...)
	return f
}
