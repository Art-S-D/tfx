package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	tfjson "github.com/hashicorp/terraform-json"
)

type RootModuleModel struct {
	module  tfjson.StateModule
	content []*render.Node
}

func RootModuleModelFromJson(json tfjson.StateModule) *render.Node {
	root := &RootModuleModel{module: json}
	out := &render.Node{Liner: root}
	for _, resource := range json.Resources {
		childResource := NewStateResourceModel(resource, out)
		root.content = append(root.content, childResource)
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(module, out)
		root.content = append(root.content, childModule)
	}
	return out
}

func (m *RootModuleModel) Address() string          { return "" }
func (m *RootModuleModel) Children() []*render.Node { return m.content }

func (m *RootModuleModel) GenerateLines(node *render.Node) []render.Line {
	var out []render.Line
	for _, model := range m.content {
		childLines := model.Lines()
		out = append(out, childLines...)
	}
	return out
}
