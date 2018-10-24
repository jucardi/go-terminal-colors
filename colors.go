package fmtc

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
	BgLightGray
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
