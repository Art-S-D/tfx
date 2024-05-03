package tfstate

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateModuleModel struct {
	module       tfjson.StateModule
	expanded     bool
	resources    []*StateResourceModel
	childModules []*StateModuleModel
}

func StateModuleModelFromJson(json tfjson.StateModule) StateModuleModel {
	result := StateModuleModel{module: json}

	for _, resource := range json.Resources {
		result.resources = append(result.resources, NewStateResourceModel(resource))
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(*module)
		result.childModules = append(result.childModules, &childModule)
	}
	return result
}

func (m *StateModuleModel) ViewHeight() int {
	if !m.expanded {
		return 1
	}

	// one line for the module name, one line for the closing }
	out := 2
	if m.module.Address == "" {
		// as the root module is not shown, out should start at 0
		out = 0
	}

	for _, r := range m.resources {
		out += r.ViewHeight()
	}
	for _, c := range m.childModules {
		out += c.ViewHeight()
	}
	return out
}

func (m *StateModuleModel) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor < 0 {
		panic(fmt.Sprintf("negative cursor %d on %v", cursor, m))
	} else if cursor == 0 && m.module.Address != "" {
		return m, 0
	} else {
		// add one line to account for the module name
		if m.module.Address != "" {
			cursor -= 1
		}
		for _, r := range m.resources {
			height := r.ViewHeight()
			if cursor < height {
				return r.Selected(cursor)
			} else {
				cursor -= height
			}
		}
		for _, c := range m.childModules {
			height := c.ViewHeight()
			if cursor < height {
				return c.Selected(cursor)
			} else {
				cursor -= height
			}
		}

		// closing bracket of the module is selected
		if cursor == 0 && m.module.Address != "" {
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

func (m *StateModuleModel) View(params render.ViewParams) string {
	var sb strings.Builder

	if m.module.Address == "" && !m.expanded {
		panic("root module is hidden, this should never happen")
	}

	if m.module.Address != "" {
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Type, "module"))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, " "))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Key, m.module.Address))
		sb.WriteString(" {")

		if !m.expanded {
			sb.WriteString(style.Preview.Render("..."))
			sb.WriteString("}")
			return sb.String()
		} else {
			sb.WriteString("\n")
			params.Cursor -= 1
		}
	}

	for _, resource := range m.resources {
		if m.module.Address != "" {
			sb.WriteString(style.Indented.Render(resource.View(params)))
			params.Cursor -= resource.ViewHeight()
		} else {
			// ignore indentation for the root module
			sb.WriteString(resource.View(params))
			params.Cursor -= resource.ViewHeight()
		}
		sb.WriteString("\n")
	}
	for _, model := range m.childModules {
		if m.module.Address != "" {
			sb.WriteString(style.Indented.Render(model.View(params)))
			params.Cursor -= model.ViewHeight()
		} else {
			// ignore indentation for the root module
			sb.WriteString(model.View(params))
			params.Cursor -= model.ViewHeight()
		}
		sb.WriteString("\n")
	}

	if m.module.Address != "" {
		if params.Cursor == 0 {
			sb.WriteString(style.Cursor.Render("}"))
		} else {
			sb.WriteString("}")
		}
	}
	params.Cursor -= 1

	result := sb.String()
	return result
}
