package render

type Model interface {
	View(params *ViewParams) string
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
	Children() []Model
}
