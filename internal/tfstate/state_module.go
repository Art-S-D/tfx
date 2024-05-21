package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
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

	firstLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: m, LineBreak: true}
	firstLine.AddSelectable(
		params.Theme.Type("module"),
		params.Theme.Default(" "),
		params.Theme.Key(m.module.Address),
	)
	firstLine.AddUnselectable(params.Theme.Default(" {"))

	if !m.Expanded {
		firstLine.AddUnselectable(
			params.Theme.Preview("..."),
			params.Theme.Default("}"),
		)
		return []render.Token{firstLine}
	}

	var out []render.Token
	out = append(out, firstLine)

	for _, model := range m.content {
		lines := model.View(params.IndentedRight())
		out = append(out, lines...)
	}

	lastLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: m, PointsToEnd: true, LineBreak: true}
	lastLine.AddSelectable(params.Theme.Default("}"))
	out = append(out, lastLine)
	return out
}
