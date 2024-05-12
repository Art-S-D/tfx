package render

const INDENT_WIDTH = 2

type Model interface {
	Lines(indent uint8) []*ScreenLine
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
	Children() []Model
}
