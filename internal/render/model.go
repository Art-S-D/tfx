package render

type Model interface {
	View(opts *Renderer)
	Selected(cursor int) (selected Model, cursorPosition int)
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
}
