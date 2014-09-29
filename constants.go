package tint

const escape = "\x1b"
const DefaultColor = 39

const (
	reset = iota + 0
	bold
	dim
	underline
	blink
	invert
	hidden
)

const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
)

const (
	LightBlack = (iota + 90)
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)
