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
	result := RootModuleModel{module: json}
	for _, resource := range json.Resources {
		result.content = append(result.content, NewStateResourceModel(resource))
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(*module)
		result.content = append(result.content, childModule)
	}
	return &result
}

func (m *RootModuleModel) Address() string          { return "" }
func (m *RootModuleModel) Children() []render.Model { return m.content }

func (m *RootModuleModel) View(params render.ViewParams) []render.Token {
	var out []render.Token
	for _, model := range m.content {
		childLines := model.View(params)
		out = append(out, childLines...)
	}
	return out
}
