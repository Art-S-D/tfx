package tfstate

import (
	"github.com/Art-S-D/tfx/internal/render"
	tfjson "github.com/hashicorp/terraform-json"
)

type RootModuleModel struct {
	module  tfjson.StateModule
	content []render.Model
}

func RootModuleModelFromJson(json tfjson.StateModule) *RootModuleModel {
	root := &RootModuleModel{module: json}
	for _, resource := range json.Resources {
		childResource := NewStateResourceModel(resource)
		root.content = append(root.content, childResource)
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(module)
		root.content = append(root.content, childModule)
	}
	return root
}

func (m *RootModuleModel) Address() string          { return "" }
func (m *RootModuleModel) Children() []render.Model { return m.content }

func (m *RootModuleModel) View() []render.Line {
	var out []render.Line
	for _, model := range m.content {
		childLines := model.View()
		out = append(out, childLines...)
	}
	return out
}
