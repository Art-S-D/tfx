package render

type Model interface {
	View() []Line
	Address() string
	Expand()
	Collapse()
	IsCollapsed() bool
	Children() []Model
}

type BaseModel struct {
	Addr     string
	Expanded bool
}

func (m *BaseModel) Address() string   { return m.Addr }
func (m *BaseModel) Expand()           { m.Expanded = true }
func (m *BaseModel) Collapse()         { m.Expanded = false }
func (m *BaseModel) IsCollapsed() bool { return !m.Expanded }
func (m *BaseModel) Children() []Model { return []Model{} }
