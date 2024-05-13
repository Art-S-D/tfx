package tfstate

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
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

func (m *StateModuleModel) Children() []render.Model {
	return m.content
}

func (m *StateModuleModel) View(params *render.ViewParams) string {
	builder := render.NewBuilder(params)

	builder.AddString(params.Theme.Type("module"))
	builder.AddString(params.Theme.Default(" "))
	builder.AddString(params.Theme.Key(m.module.Address))

	builder.AddUnSelectableString(params.Theme.Default(" {"))

	if !m.expanded {
		builder.AddUnSelectableString(params.Theme.Preview("..."))
		builder.AddUnSelectableString(params.Theme.Default("}"))
		return builder.String()
	}

	params.IndentRight()
	for _, model := range m.content {
		builder.InsertNewLine()
		params.NextLine()

		builder.WriteString(model.View(params))
	}

	params.IndentLeft()
	builder.InsertNewLine()
	params.NextLine()
	builder.AddString(params.Theme.Default("}"))
	return builder.String()
}
