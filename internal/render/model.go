package render

type Model interface {
	Lines(indent uint8) []*ScreenLine
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
	Children() []Model
}
