package fmtc

import (
	"fmt"
	"github.com/jucardi/go-strings/stringx"
)

type Color int

// special formats
const (
	Clear Color = iota
	Bold
	Dim
	Italic
	Underline
	Inverted
	Hidden
)

// foreground
const (
	Black Color = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	Gray
)
const (
	DarkGray Color = 90 + iota
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

// background
const (
	BgBlack Color = 40 + iota
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgGray
)
const (
	BgDarkGray Color = 100 + iota
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhite
)

var colorMap = map[string]Color{
	"clear":     Clear,
	"bold":      Bold,
	"dim":       Dim,
	"italic":    Italic,
	"underline": Underline,
	"inverted":  Inverted,
	"hidden":    Hidden,

	"black":        Black,
	"red":          Red,
	"green":        Green,
	"yellow":       Yellow,
	"blue":         Blue,
	"magenta":      Magenta,
	"cyan":         Cyan,
	"gray":         Gray,
	"darkgray":     DarkGray,
	"lightred":     LightRed,
	"lightgreen":   LightGreen,
	"lightyellow":  LightYellow,
	"lightblue":    LightBlue,
	"lightmagenta": LightMagenta,
	"lightcyan":    LightCyan,
	"white":        White,

	"bgblack":        BgBlack,
	"bgred":          BgRed,
	"bggreen":        BgGreen,
	"bgyellow":       BgYellow,
	"bgblue":         BgBlue,
	"bgmagenta":      BgMagenta,
	"bgcyan":         BgCyan,
	"bggray":         BgGray,
	"bgdarkgray":     BgDarkGray,
	"bglightred":     BgLightRed,
	"bglightgreen":   BgLightGreen,
	"bglightyellow":  BgLightYellow,
	"bglightblue":    BgLightBlue,
	"bglightmagenta": BgLightMagenta,
	"bglightcyan":    BgLightCyan,
	"bgwhite":        BgWhite,
}

func Parse(color string) (Color, error) {
	if v, ok := colorMap[stringx.New(color).ToLower().TrimSpace().S()]; ok {
		return v, nil
	}

	var l Color
	return l, fmt.Errorf("not a valid color: %q", color)
}
