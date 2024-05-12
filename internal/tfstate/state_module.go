package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateModuleModel struct {
	module   tfjson.StateModule
	expanded bool
	content  []render.Model
}

func StateModuleModelFromJson(json tfjson.StateModule) *StateModuleModel {
	result := StateModuleModel{module: json}

	for _, resource := range json.Resources {
		result.content = append(result.content, NewStateResourceModel(resource))
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(*module)
		result.content = append(result.content, childModule)
	}
	return &result
}

func (m *StateModuleModel) ViewHeight() int {
	if !m.expanded {
		return 1
	}

	// one line for the module name, one line for the closing }
	out := 2
	for _, c := range m.content {
		out += c.ViewHeight()
	}
	return out
}

func (m *StateModuleModel) Address() string {
	return m.module.Address
}
func (m *StateModuleModel) Expand() {
	m.expanded = true
}
func (m *StateModuleModel) Collapse() {
	m.expanded = false
}

func (m *StateModuleModel) Children() []render.Model {
	return m.content
}

func (m *StateModuleModel) Lines(indent uint8) []*render.ScreenLine {
	var out []*render.ScreenLine

	firstLine := render.ScreenLine{Indentation: indent, PointsTo: m}
	firstLine.AddString(style.Type, "module")
	firstLine.AddString(style.Default, " ")
	firstLine.AddString(style.Key, m.module.Address)
	firstLine.AddUnSelectableString(style.Default, " {")

	if !m.expanded {
		firstLine.AddUnSelectableString(style.Preview, "...")
		firstLine.AddUnSelectableString(style.Default, "}")
		out = append(out, &firstLine)
		return out
	}

	out = append(out, &firstLine)

	for _, model := range m.content {
		out = append(out, model.Lines(indent+render.INDENT_WIDTH)...)
	}

	lastLine := render.ScreenLine{Indentation: indent, PointsTo: m, PointsToModelEnd: true}
	lastLine.AddString(style.Default, "}")
	out = append(out, &lastLine)

	return out
}
