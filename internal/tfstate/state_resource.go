package tfstate

import (
	encodingJson "encoding/json"
	"fmt"
	"slices"

	json "github.com/Art-S-D/tfx/internal/json"
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateResourceModel struct {
	resource   *tfjson.StateResource
	attributes map[string]render.Model
	expanded   bool
}

func NewStateResourceModel(resource *tfjson.StateResource) *StateResourceModel {
	out := StateResourceModel{resource, make(map[string]render.Model), false}
	for k, v := range resource.AttributeValues {
		addr := fmt.Sprintf("%s.%s", resource.Address, k)

		var sensitive map[string]any
		err := encodingJson.Unmarshal(resource.SensitiveValues, &sensitive)
		if err != nil {
			panic(fmt.Sprintf("failed to parse sensitive value %v", sensitive))
		}

		out.attributes[k], err = json.ParseValue(v, sensitive[k], addr)
		if err != nil {
			panic(fmt.Errorf("failed to create state resource %w", err))
		}
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
	cursor -= 1
	for _, k := range m.Keys() {
		v := m.attributes[k]
		height := v.ViewHeight()
		if cursor < height {
			return v.Selected(cursor)
		} else {
			cursor -= height
		}
	}
	if cursor == 0 {
		return m, m.ViewHeight() - 1
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

func (m *StateResourceModel) Children() []render.Model {
	var out []render.Model
	keys := m.Keys()
	for _, k := range keys {
		out = append(out, m.attributes[k])
	}
	return out
}

func (m *StateResourceModel) View(r *render.Renderer) {

	// render first line (without the final brace)
	r.CursorWrite(style.Type, m.resourceMode())
	r.CursorWrite(style.Default, " ")
	r.CursorWrite(style.Key, m.resource.Type)
	r.CursorWrite(style.Default, " ")
	r.CursorWrite(style.Key, m.resource.Name)

	if m.resource.Index != nil {
		r.CursorWrite(style.Default, " ")
		r.CursorWrite(style.Key, m.resourceIndex())
	}

	// render braces
	if !m.expanded {
		r.Write(" {")
		r.Write(style.Preview.Render("..."))
		r.Write("}")
		return
	}

	r.Write(" {")
	r.IndentRight()

	// render resource body
	keys := m.Keys()
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.attributes[k]

		r.NewLine()
		r.CursorWrite(style.Key, k)

		for range len(longestKey) - len(k) {
			r.Write(" ")
		}

		r.EndCursorForCurrentLine()
		r.Write(" = ")

		// this makes it so that only the key is selected instead of the key and the value
		// FIXME
		// vParams := params
		// if params.Cursor == 0 {
		// 	vParams.Cursor -= 1
		// }

		v.View(r)
	}

	r.IndentLeft()
	r.NewLine()
	r.CursorWrite(style.Default, "}")
}
