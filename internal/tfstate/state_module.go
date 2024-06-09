package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateModuleModel struct {
	render.BaseModel
	module  *tfjson.StateModule
	content []render.Model
}

func StateModuleModelFromJson(json *tfjson.StateModule) *StateModuleModel {
	module := &StateModuleModel{module: json}
	module.Addr = json.Address

	for _, resource := range json.Resources {
		childResource := NewStateResourceModel(resource)
		module.content = append(module.content, childResource)
	}
	for _, mod := range json.ChildModules {
		childModule := StateModuleModelFromJson(mod)
		module.content = append(module.content, childModule)
	}
	return module
}

func (m *StateModuleModel) Children() []render.Model {
	return m.content
}

func (m *StateModuleModel) View() []render.Line {

	firstLine := render.Line{PointsTo: m}
	firstLine.AddSelectable(
		style.Type("module"),
		style.Default(" "),
		style.Key(m.module.Address),
	)
	firstLine.AddUnselectable(style.Default(" {"))

	if !m.Expanded {
		firstLine.AddUnselectable(
			style.Preview("..."),
			style.Default("}"),
		)
		return []render.Line{firstLine}
	}

	var out []render.Line
	out = append(out, firstLine)

	for _, model := range m.content {
		lines := model.View()
		render.Indent(lines)
		out = append(out, lines...)
	}

	lastLine := render.Line{PointsTo: m, PointsToEnd: true}
	lastLine.AddSelectable(style.Default("}"))
	out = append(out, lastLine)
	return out
}
