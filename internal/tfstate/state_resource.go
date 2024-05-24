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
	attributes map[string]*render.Node
}

func NewStateResourceModel(jsonResource *tfjson.StateResource, parent *render.Node) *render.Node {
	resource := &StateResourceModel{resource: jsonResource, attributes: make(map[string]*render.Node)}
	out := &render.Node{Liner: resource, Parent: parent, Depth: parent.Depth + 1, Address: jsonResource.Address}
	keys := utils.KeysOrdered(jsonResource.AttributeValues)
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for k, v := range jsonResource.AttributeValues {
		addr := fmt.Sprintf("%s.%s", jsonResource.Address, k)

		var sensitive map[string]any
		err := encodingJson.Unmarshal(jsonResource.SensitiveValues, &sensitive)
		if err != nil {
			panic(fmt.Sprintf("failed to parse sensitive value %v", sensitive))
		}

		value, err := json.ParseValue(v, sensitive[k], out, addr)
		if err != nil {
			panic(fmt.Errorf("failed to create state resource %w", err))
		}
		kv := &json.KeyVal{Key: k, Value: value, KeyPadding: uint8(len(longestKey))}
		resource.attributes[k] = &render.Node{Address: addr, Parent: out, Depth: out.Depth + 1, Liner: kv}
	}
	return out
}

func (m *StateResourceModel) Keys() []string {
	return utils.KeysOrdered(m.attributes)
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

func (m *StateResourceModel) Children() []*render.Node {
	var resource []*render.Node
	keys := m.Keys()
	for _, k := range keys {
		resource = append(resource, m.attributes[k])
	}
	return resource
}

func (m *StateResourceModel) GenerateLines(node *render.Node) []render.Line {
	firstLine := render.Line{Indentation: node.Depth, PointsTo: node}
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

	if !node.Expanded {
		firstLine.AddUnselectable(
			style.Preview("..."),
			style.Default("}"),
		)
		return []render.Line{firstLine}
	}

	var resource []render.Line
	resource = append(resource, firstLine)

	// render resource body
	keys := m.Keys()
	for _, k := range keys {
		v := m.attributes[k]

		lines := v.Lines()
		resource = append(resource, lines...)
	}

	lastLine := render.Line{Indentation: node.Depth, PointsTo: node, PointsToEnd: true}
	lastLine.AddSelectable(style.Default("}"))
	resource = append(resource, lastLine)
	return resource
}
