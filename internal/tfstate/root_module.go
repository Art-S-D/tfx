package tfstate

import (
	"fmt"

	"github.com/Art-S-D/tfview/internal/render"
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

func (m *RootModuleModel) ViewHeight() int {
	out := 0
	for _, c := range m.content {
		out += c.ViewHeight()
	}
	return out
}

func (m *RootModuleModel) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor < 0 {
		panic(fmt.Sprintf("negative cursor %d on %v", cursor, m))
	} else {
		for _, c := range m.content {
			height := c.ViewHeight()
			if cursor < height {
				return c.Selected(cursor)
			} else {
				cursor -= height
			}
		}

		panic(fmt.Sprintf("cursor out of bounds %d for root module of height %d", cursor, m.ViewHeight()))
	}
}

func (m *RootModuleModel) Address() string {
	return ""
}
func (m *RootModuleModel) Expand() {
	panic("callint expand on the root module")
}
func (m *RootModuleModel) Collapse() {
	panic("calling collapse on the root module")
}

func (m *RootModuleModel) View(r *render.Renderer) {
	for i, model := range m.content {
		model.View(r)

		// skip last line for the root module
		if i < len(m.content)-1 {
			r.NewLine()
		}
	}
}
