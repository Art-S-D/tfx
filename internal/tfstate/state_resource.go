package tfstate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Art-S-D/tfview/internal/json"
	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
	"github.com/Art-S-D/tfview/internal/utils"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateResourceModel struct {
	resource *tfjson.StateResource
	// TODO: replace with json.jsonObject
	attributes map[string]render.Model
	expanded   bool
}

func NewStateResourceModel(resource *tfjson.StateResource) *StateResourceModel {
	out := StateResourceModel{resource, make(map[string]render.Model), false}
	for k, v := range resource.AttributeValues {
		out.attributes[k] = json.ParseValue(v, resource.Address)
	}
	return &out
}

func (m *StateResourceModel) Keys() []string {
	return utils.KeysOrdered(m.attributes)
}

func (m *StateResourceModel) ViewHeight() int {
	if !m.expanded {
		return 1
	} else {
		// one for each curly brackets
		out := 2
		for _, v := range m.attributes {
			out += v.ViewHeight()
		}
		return out
	}
}

func (m *StateResourceModel) Address() string {
	return m.resource.Address
}
func (m *StateResourceModel) Expand() {
	m.expanded = true
}
func (m *StateResourceModel) Collapse() {
	m.expanded = false
}

func (m *StateResourceModel) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor == 0 {
		return m, 0
	}
	for _, v := range m.attributes {
		height := v.ViewHeight()
		if cursor < height {
			return v, cursor
		} else {
			cursor -= height
		}
	}
	if cursor == 0 {
		return m, m.ViewHeight()
	}
	panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, m, m.ViewHeight()))
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

func (m *StateResourceModel) View(params render.ViewParams) string {
	var sb strings.Builder

	// render first line (without the final brace)
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Type, m.resourceMode()))
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, " "))
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Key, m.resource.Type))
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, " "))
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Key, m.resource.Name))
	if m.resource.Index != nil {
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, " "))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Key, m.resourceIndex()))
	}

	// render braces
	if !m.expanded {
		sb.WriteString(" {")
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String()
	} else {
		sb.WriteString(" {\n")
		params.Cursor -= 1
	}

	// render resource body
	keys := m.Keys()
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.attributes[k]

		sb.WriteString(style.Indented.Render(style.RenderStyleOrCursor(params.Cursor, style.Key, k)))

		for range len(longestKey) - len(k) {
			sb.WriteRune(' ')
		}

		sb.WriteString(" = ")
		sb.WriteString(v.View(params))
		sb.WriteString("\n")
		params.Cursor -= v.ViewHeight()
	}

	// render braces
	sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "}"))
	params.Cursor -= 1

	return sb.String()
}
