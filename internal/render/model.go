package render

type Model interface {
	View() []Line
	Address() string
}

type Collapser interface {
	Expand()
	Collapse()
	IsCollapsed() bool
}
type Childrener interface {
	Children() []Model
}

type BaseCollapser struct {
	Expanded bool
}

func (c *BaseCollapser) Expand()           { c.Expanded = true }
func (c *BaseCollapser) Collapse()         { c.Expanded = false }
func (c *BaseCollapser) IsCollapsed() bool { return !c.Expanded }

type BaseModel struct {
	Addr string
}

func (m *BaseModel) Address() string { return m.Addr }
