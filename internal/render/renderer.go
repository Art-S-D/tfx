package render

type ViewParams struct {
	Cursor int
	Width  int
	// Height int
}

type Model interface {
	View(opts ViewParams) string
	Selected(cursor int) (selected Model, cursorPosition int)
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
}
