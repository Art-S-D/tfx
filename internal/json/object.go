package json

import (
	"fmt"
	"slices"

	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
	"golang.org/x/exp/maps"
)

func emptyObject(address string) *node.Node {
	out := &node.Node{}
	s := style.Default("{}").Selectable()
	out.SetCollapsed(s)
	out.SetExpanded(s)
	out.SetAddress(address)
	return out
}

func jsonObjectNode(address string, object map[string]any, sensitiveValues any) (*node.Node, error) {
	if len(object) == 0 {
		return emptyObject(address), nil
	}

	out := &node.Node{}
	out.SetAddress(address)
	out.SetExpanded(style.Default("{").Selectable())
	out.SetCollapsed(
		style.Concat(
			style.Default("{"),
			style.Preview("..."),
			style.Default("}"),
		).Selectable(),
	)

	if _, ok := sensitiveValues.(bool); ok {
		out.SetSensitive(true)
		sensitiveValues = nil
	}
	sensitive, ok := sensitiveValues.(map[string]any)
	if sensitiveValues != nil && !ok {
		return nil, fmt.Errorf("failed to parse sensitive value to object %v", sensitiveValues)
	}

	longestKey := slices.MaxFunc(maps.Keys(object), func(s1, s2 string) int { return len(s1) - len(s2) })
	keys := utils.KeysOrdered(object)

	for _, k := range keys {
		v := object[k]

		addr := fmt.Sprintf("%s.%s", address, k)
		var nextSensitive any
		if sensitiveValues != nil {
			nextSensitive = sensitive[k]
		}
		parsed, err := ParseValue(v, nextSensitive, addr)
		if err != nil {
			return nil, err
		}
		parsed.SetKey(k, uint8(len(longestKey)-len(k)))
		parsed.IncreaseDepth()
		out.AppendChild(parsed)
	}

	lastChild := node.String("}")
	lastChild.SetAddress(address)
	out.AppendChild(lastChild)

	return out, nil
}
