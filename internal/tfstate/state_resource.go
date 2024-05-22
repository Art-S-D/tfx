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
	render.BaseCollapser
	resource   *tfjson.StateResource
	attributes map[string]render.Model
}

func NewStateResourceModel(resource *tfjson.StateResource) *StateResourceModel {
	out := StateResourceModel{resource: resource, attributes: make(map[string]render.Model)}
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

func (m *StateResourceModel) Address() string {
	return m.resource.Address
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

func (m *StateResourceModel) View(params render.ViewParams) []render.Line {
	firstLine := render.Line{Indentation: params.Indentation, PointsTo: m}
	firstLine.AddSelectable(
		style.Type(m.resourceMode()),
		style.Default(" "),
		style.Key(m.resource.Type),
		style.Default(" "),
		style.Key(m.resource.Name),
	)

	if m.resource.Index != nil {
		firstLine.AddSelectable(
			style.Default(" "),
			style.Key(m.resourceIndex()),
		)
	}

	firstLine.AddUnselectable(style.Default(" {"))

	if !m.Expanded {
		firstLine.AddUnselectable(
			style.Preview("..."),
			style.Default("}"),
		)
		return []render.Line{firstLine}
	}

	var out []render.Line
	out = append(out, firstLine)

	// render resource body
	keys := m.Keys()
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.attributes[k]

		kv := json.KeyVal{Key: k, Value: v, KeyPadding: uint8(len(longestKey))}
		lines := kv.View(params.IndentedRight())
		out = append(out, lines...)
	}

	lastLine := render.Line{Indentation: params.Indentation, PointsTo: m, PointsToEnd: true}
	lastLine.AddSelectable(style.Default("}"))
	out = append(out, lastLine)
	return out
}
