package tfstate

import (
	encodingJson "encoding/json"
	"fmt"
	"slices"
	"strings"

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

func (m *StateResourceModel) Lines(indent uint8) []*render.ScreenLine {
	var out []*render.ScreenLine

	firstLine := render.ScreenLine{Indentation: indent, PointsTo: m}

	firstLine.AddString(style.Type, m.resourceMode())
	firstLine.AddString(style.Default, " ")
	firstLine.AddString(style.Key, m.resource.Type)
	firstLine.AddString(style.Default, " ")
	firstLine.AddString(style.Key, m.resource.Name)

	if m.resource.Index != nil {
		firstLine.AddString(style.Default, " ")
		firstLine.AddString(style.Key, m.resourceIndex())
	}

	if !m.expanded {
		firstLine.AddUnSelectableString(style.Default, " {")
		firstLine.AddUnSelectableString(style.Preview, "...")
		firstLine.AddUnSelectableString(style.Default, "}")
		out = append(out, &firstLine)
		return out
	}

	firstLine.AddUnSelectableString(style.Default, " {")
	out = append(out, &firstLine)

	// render resource body
	keys := m.Keys()
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.attributes[k]

		line := render.ScreenLine{Indentation: indent + render.INDENT_WIDTH, PointsTo: v}

		line.AddString(style.Key, k)
		line.AddString(style.Default, strings.Repeat(" ", len(longestKey)-len(k)))

		line.AddUnSelectableString(style.Default, " = ")

		nextLines := v.Lines(indent + render.INDENT_WIDTH)
		nextLines[0].RemoveCursor()
		line.MergeWith(nextLines[0])

		out = append(out, &line)
		out = append(out, nextLines[1:]...)
	}

	lastLine := render.ScreenLine{Indentation: indent, PointsTo: m, PointsToModelEnd: true}
	lastLine.AddString(style.Default, "}")
	out = append(out, &lastLine)
	return out
}
