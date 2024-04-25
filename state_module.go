package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	tea "github.com/charmbracelet/bubbletea"
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

func (m *StateModuleModel) Height() int {
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
		out += r.Height()
	}
	for _, c := range m.childModules {
		out += c.Height()
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
			height := r.Height()
			if cursor < height {
				return r
			} else {
				cursor -= height
			}
		}
		for _, c := range m.childModules {
			height := c.Height()
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
		panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, m, m.Height()))
	}
}

func (m *StateModuleModel) Address() string {
	return m.module.Address
}
func (m *StateModuleModel) Toggle() {
	m.expanded = !m.expanded
}

func (m *StateModuleModel) Init() tea.Cmd {
	return nil
}

func (m *StateModuleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}

func (m *StateModuleModel) View() string {
	var sb strings.Builder

	if m.module.Address != "" {
		sb.WriteString(style.Type.Render("module"))
		sb.WriteString(" ")
		sb.WriteString(style.Key.Render(m.module.Address))
		sb.WriteString(" {\n")
	}

	for _, resource := range m.resources {
		if m.module.Address != "" {
			sb.WriteString(style.Indented.Render(resource.View()))
		} else {
			// ignore indentation for the root module
			sb.WriteString(resource.View())
		}
		sb.WriteString("\n")
	}
	for _, model := range m.childModules {
		if m.module.Address != "" {
			sb.WriteString(style.Indented.Render(model.View()))
		} else {
			// ignore indentation for the root module
			sb.WriteString(model.View())
		}
		sb.WriteString("\n")
	}

	if m.module.Address != "" {
		sb.WriteString("}")
	}
	result := sb.String()
	return result
}

func (m *StateModuleModel) ViewCursor(cursor int) (view string, newCursor int) {
	var sb strings.Builder

	if m.module.Address == "" && !m.expanded {
		panic("root module is hidden, this should never happen")
	}

	if m.module.Address != "" {
		if cursor == 0 {
			sb.WriteString(style.Cursor.Render("module"))
			sb.WriteString(style.Cursor.Render(" "))
			sb.WriteString(style.Cursor.Render(m.module.Address))
			sb.WriteString(" {")
		} else {
			sb.WriteString(style.Type.Render("module"))
			sb.WriteString(" ")
			sb.WriteString(style.Key.Render(m.module.Address))
			sb.WriteString(" {")
		}
		cursor -= 1
	}

	if !m.expanded {
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String(), cursor
	} else {
		sb.WriteString("\n")
	}

	for _, resource := range m.resources {
		if m.module.Address != "" {
			var resourceView string
			resourceView, cursor = resource.ViewCursor(cursor)
			sb.WriteString(style.Indented.Render(resourceView))
		} else {
			// ignore indentation for the root module
			var resourceView string
			resourceView, cursor = resource.ViewCursor(cursor)
			sb.WriteString(resourceView)
		}
		sb.WriteString("\n")
	}
	for _, model := range m.childModules {
		if m.module.Address != "" {
			var modelView string
			modelView, cursor = model.ViewCursor(cursor)
			sb.WriteString(style.Indented.Render(modelView))
		} else {
			// ignore indentation for the root module
			var modelView string
			modelView, cursor = model.ViewCursor(cursor)
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
