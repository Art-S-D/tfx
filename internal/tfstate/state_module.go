package tfstate

import (
	"fmt"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
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

func (m *StateModuleModel) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor < 0 {
		panic(fmt.Sprintf("negative cursor %d on %v", cursor, m))
	} else if cursor == 0 {
		return m, 0
	} else {
		// add one line to account for the module name
		if m.module.Address != "" {
			cursor -= 1
		}
		for _, c := range m.content {
			height := c.ViewHeight()
			if cursor < height {
				return c.Selected(cursor)
			} else {
				cursor -= height
			}
		}

		// closing bracket of the module is selected
		if cursor == 0 {
			return m, m.ViewHeight() - 1
		}
		panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, m, m.ViewHeight()))
	}
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

func (m *StateModuleModel) View(r *render.Renderer) {
	r.CursorWrite(style.Type, "module")
	r.CursorWrite(style.Default, " ")
	r.CursorWrite(style.Key, m.module.Address)
	r.Write(" {")

	if !m.expanded {
		r.Write(style.Preview.Render("..."))
		r.Write("}")
		return
	} else {
		r.NewLine()
	}

	r.IndentRight()

	for _, model := range m.content {
		model.View(r)
		r.NewLine()
	}

	r.IndentLeft()
	r.CursorWrite(style.Default, "}")
}
