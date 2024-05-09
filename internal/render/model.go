package render

type Model interface {
	View(params *ViewParams) string
	Selected(cursor int) (selected Model, cursorPosition int)
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
	Children() []Model
}
