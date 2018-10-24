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
}

type colorFlow struct {
	customWriter bool
	writer       io.Writer
}

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

func (f *colorFlow) String() string {
	if f.customWriter {
		return ""
	}
	buf := f.writer.(*bytes.Buffer)
	return buf.String()
}

func (f *colorFlow) Print(str interface{}, colors ... Color) IColorFlow {
	f.doColors(colors...)
	fmt.Fprint(f.writer, str, clear)

	return f
}

func (f *colorFlow) PrintLn(str interface{}, colors ... Color) IColorFlow {
	f.doColors(colors...)
	fmt.Fprintln(f.writer, str, clear)
	return f
}

func (f *colorFlow) Printf(format string, args []interface{}, colors ...Color) IColorFlow {
	f.doColors(colors...)
	fmt.Fprintf(f.writer, format, args...)
	fmt.Fprint(f.writer, clear)
}

func (f *colorFlow) doColors(colors ...Color) {
	for _, c := range colors {
		fmt.Fprint(f.writer, esc, c, "m")
	}
}
