package main

import (
	"fmt"
	"strings"

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
		result.resources = append(result.resources, &StateResourceModel{*resource, false})
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleModelFromJson(*module)
		result.childModules = append(result.childModules, &childModule)
	}
	return result
}

func (m *StateModuleModel) RenderingHeight() int {
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
		out += r.RenderingHeight()
	}
	for _, c := range m.childModules {
		out += c.RenderingHeight()
	}
	return out
}

func (m *StateModuleModel) Selected(cursor int) CursorSelector {
	if cursor < 0 {
		panic(fmt.Sprintf("negative cursor %d on %v", cursor, m))
	} else if cursor == 0 && m.module.Address != "" {
		return m
	} else {
		// add one line to account for the module name
		if m.module.Address != "" {
			cursor -= 1
		}
		for _, r := range m.resources {
			height := r.RenderingHeight()
			if cursor < height {
				return r
			} else {
				cursor -= height
			}
		}
		for _, c := range m.childModules {
			height := c.RenderingHeight()
			if cursor < height {
				return c.Selected(cursor)
			} else {
				cursor -= height
			}
		}

		// closing bracket of the module is selected
		if cursor == 0 && m.module.Address != "" {
			return m
		}
		panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, m, m.RenderingHeight()))
	}
}

func (m *StateModuleModel) Address() string {
	return m.module.Address
}
func (m *StateModuleModel) Toggle() {
	m.expanded = !m.expanded
}

func (m *StateModuleModel) View(cursor int) (view string, newCursor int) {
	var sb strings.Builder

	if m.module.Address == "" && !m.expanded {
		panic("root module is hidden, this should never happen")
	}

	if m.module.Address != "" {
		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Type, "module"))
		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Default, " "))
		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Key, m.module.Address))
		sb.WriteString(" {")
		cursor -= 1
	}

	if !m.expanded {
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String(), cursor
	} else if m.module.Address != "" {
		sb.WriteString("\n")
	}

	for _, resource := range m.resources {
		if m.module.Address != "" {
			var resourceView string
			resourceView, cursor = resource.View(cursor)
			sb.WriteString(style.Indented.Render(resourceView))
		} else {
			// ignore indentation for the root module
			var resourceView string
			resourceView, cursor = resource.View(cursor)
			sb.WriteString(resourceView)
		}
		sb.WriteString("\n")
	}
	for _, model := range m.childModules {
		if m.module.Address != "" {
			var modelView string
			modelView, cursor = model.View(cursor)
			sb.WriteString(style.Indented.Render(modelView))
		} else {
			// ignore indentation for the root module
			var modelView string
			modelView, cursor = model.View(cursor)
			sb.WriteString(modelView)
		}
		sb.WriteString("\n")
	}

	if m.module.Address != "" {
		if cursor == 0 {
			sb.WriteString(style.Cursor.Render("}"))
		} else {
			sb.WriteString("}")
		}
	}
	cursor -= 1

	result := sb.String()
	return result, cursor
}
