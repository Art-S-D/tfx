package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateModuleModel struct {
	module  *tfjson.StateModule
	content []*render.Node
}

func StateModuleModelFromJson(json *tfjson.StateModule, parent *render.Node) *render.Node {
	module := &StateModuleModel{module: json}
	out := &render.Node{Address: json.Address, Depth: parent.Depth + 1, Parent: parent, Liner: module}

	for _, resource := range json.Resources {
		childResource := NewStateResourceModel(resource, out)
		module.content = append(module.content, childResource)
	}
	for _, mod := range json.ChildModules {
		childModule := StateModuleModelFromJson(mod, out)
		module.content = append(module.content, out, childModule)
	}
	return out
}

func (m *StateModuleModel) Address() string {
	return m.module.Address
}

func (m *StateModuleModel) Children() []*render.Node {
	return m.content
}

func (m *StateModuleModel) GenerateLines(node *render.Node) []render.Line {
	firstLine := render.Line{Indentation: node.Depth, PointsTo: node}
	firstLine.AddSelectable(
		style.Type("module"),
		style.Default(" "),
		style.Key(m.module.Address),
	)
	firstLine.AddUnselectable(style.Default(" {"))

	if !node.Expanded {
		firstLine.AddUnselectable(
			style.Preview("..."),
			style.Default("}"),
		)
		return []render.Line{firstLine}
	}

	var out []render.Line
	out = append(out, firstLine)

	for _, model := range m.content {
		lines := model.Lines()
		out = append(out, lines...)
	}

	lastLine := render.Line{Indentation: node.Depth, PointsTo: node, PointsToEnd: true}
	lastLine.AddSelectable(style.Default("}"))
	out = append(out, lastLine)
	return out
}
