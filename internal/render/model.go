package render

type Model interface {
	View(params ViewParams) []Token
	Address() string
}

type Collapser interface {
	Expand()
	Collapse()
}
type Childrener interface {
	Children() []Model
}

type BaseCollapser struct {
	Expanded bool
}

func (c *BaseCollapser) Expand()   { c.Expanded = true }
func (c *BaseCollapser) Collapse() { c.Expanded = false }

type BaseModel struct {
	Addr string
}

func (m *BaseModel) Address() string { return m.Addr }
