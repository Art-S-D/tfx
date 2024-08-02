package tfstate

import (
	encodingJson "encoding/json"
	"fmt"
	"slices"

	json "github.com/Art-S-D/tfx/internal/json"
	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
	tfjson "github.com/hashicorp/terraform-json"
)

func sensitiveValues(json *tfjson.StateResource) map[string]any {
	var sensitive map[string]any
	err := encodingJson.Unmarshal(json.SensitiveValues, &sensitive)
	if err != nil {
		panic(fmt.Sprintf("failed to parse sensitive value %v", sensitive))
	}
	return sensitive
}
func StateResourceNode(jsonResource *tfjson.StateResource) *node.Node {
	resource := &node.Node{}
	resource.SetAddress(jsonResource.Address)

	keys := utils.KeysOrdered(jsonResource.AttributeValues)
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })

	sensitive := sensitiveValues(jsonResource)

	for _, k := range keys {
		v := jsonResource.AttributeValues[k]

		addr := fmt.Sprintf("%s.%s", jsonResource.Address, k)

		value, err := json.ParseValue(v, sensitive[k], addr)
		if err != nil {
			panic(fmt.Errorf("failed to create state resource: %w", err))
		}

		value.SetKey(k, uint8(len(longestKey)-len(k)))
		value.IncreaseDepth()
		resource.AppendChild(value)
	}

	resource.SetExpanded(resourceExpanded(jsonResource))
	resource.SetCollapsed(resourceCollapsed(jsonResource))

	resource.AppendEndNode("}")
	return resource
}

func mode(resource *tfjson.StateResource) string {
	resourceMode := "resource"
	if resource.Mode == tfjson.DataResourceMode {
		resourceMode = "data"
	}
	return resourceMode
}
func resourceExpanded(resource *tfjson.StateResource) style.Str {
	s := style.Concat(
		style.Type(mode(resource)).Selectable(),
		style.Space().Selectable(),
		style.Key(resource.Type).Selectable(),
		style.Space().Selectable(),
		style.Key(resource.Name).Selectable(),
	)
	if resource.Index != nil {
		s = style.Concat(
			s,
			style.Space().Selectable(),
			style.Key(resourceIndexToStr(resource.Index)).Selectable(),
		)
	}
	return style.Concat(
		s,
		style.Space().UnSelectable(),
		style.Default("{").UnSelectable(),
	)
}

func resourceCollapsed(resource *tfjson.StateResource) style.Str {
	return style.Concat(
		resourceExpanded(resource),
		style.Preview("...").UnSelectable(),
		style.Default("}").UnSelectable(),
	)
}
