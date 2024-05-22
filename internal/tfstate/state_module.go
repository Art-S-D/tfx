package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateModuleModel struct {
	render.BaseCollapser

	module  tfjson.StateModule
	content []render.Model
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

func (m *StateModuleModel) Address() string {
	return m.module.Address
}

func (m *StateModuleModel) Children() []render.Model {
	return m.content
}

func (m *StateModuleModel) View(params render.ViewParams) []render.Token {

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
		return []render.Token{firstLine}
	}

	var out []render.Token
	out = append(out, firstLine)

	for _, model := range m.content {
		lines := model.View(params.IndentedRight())
		out = append(out, lines...)
	}

	lastLine := render.Line{PointsTo: m, PointsToEnd: true}
	lastLine.AddSelectable(style.Default("}"))
	out = append(out, lastLine)
	return out
}
