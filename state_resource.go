package main

import (
	"slices"
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	"github.com/Art-S-D/tfview/internal/utils"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateResourceModel struct {
	resource tfjson.StateResource
	expanded bool
}

func (m *StateResourceModel) RenderingHeight() int {
	if !m.expanded {
		return 1
	} else {
		return 2 + len(m.resource.AttributeValues)
	}
}

func (m *StateResourceModel) Address() string {
	return m.resource.Address
}
func (m *StateResourceModel) Toggle() {
	m.expanded = !m.expanded
}

func (m *StateResourceModel) resourceIndex() string {
	return resourceIndexToStr(m.resource.Index)
}
func (m *StateResourceModel) resourceMode() string {
	resourceMode := "resource"
	if m.resource.Mode == tfjson.DataResourceMode {
		resourceMode = "data"
	}
	return resourceMode
}

func (m *StateResourceModel) View(cursor int) (out string, newCursor int) {
	var sb strings.Builder

	// render first line (without the final brace)
	sb.WriteString(style.RenderStyleOrCursor(cursor, style.Type, m.resourceMode()))
	sb.WriteString(style.RenderStyleOrCursor(cursor, style.Default, " "))
	sb.WriteString(style.RenderStyleOrCursor(cursor, style.Key, m.resource.Type))
	sb.WriteString(style.RenderStyleOrCursor(cursor, style.Key, m.resource.Name))
	if m.resource.Index != nil {
		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Default, " "))
		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Key, m.resourceIndex()))
	}
	cursor -= 1

	// render braces
	if !m.expanded {
		sb.WriteString(" {")
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String(), cursor
	} else {
		sb.WriteString(" {\n")
	}

	// render resource body
	keys := utils.KeysOrdered(m.resource.AttributeValues)
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.resource.AttributeValues[k]
		sb.WriteString("  ")

		sb.WriteString(style.RenderStyleOrCursor(cursor, style.Key, k))

		for range len(longestKey) - len(k) {
			sb.WriteRune(' ')
		}

		sb.WriteString(" = ")
		sb.WriteString(style.Render(v))
		sb.WriteString("\n")
		cursor -= 1
	}

	// render braces
	sb.WriteString(style.RenderStyleOrCursor(cursor, style.Default, "}"))

	cursor -= 1
	return sb.String(), cursor
}
