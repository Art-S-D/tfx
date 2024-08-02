package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
)

func emptyArray(address string) *node.Node {
	out := &node.Node{}
	s := style.Default("[]").Selectable()
	out.SetCollapsed(s)
	out.SetExpanded(s)
	out.SetAddress(address)
	return out
}
func jsonArrayNode(address string, array []any, sensitiveValues any) (*node.Node, error) {
	if len(array) == 0 {
		return emptyArray(address), nil
	}

	out := &node.Node{}
	out.SetAddress(address)
	out.SetExpanded(style.Default("[").Selectable())
	out.SetCollapsed(
		style.Concat(
			style.Default("["),
			style.Preview("..."),
			style.Default("]"),
		).Selectable(),
	)

	if _, ok := sensitiveValues.(bool); ok {
		out.SetSensitive(true)
		sensitiveValues = nil
	}
	sensitive, ok := sensitiveValues.([]any)
	if sensitiveValues != nil && !ok {
		return nil, fmt.Errorf("failed to parse sensitive value to array sensitiveValues:%v for array:%v", sensitiveValues, array)
	}

	for i, v := range array {
		addr := fmt.Sprintf("%s[%d]", address, i)
		var nextSensitive any
		if sensitiveValues != nil {
			nextSensitive = sensitive[i]
		}
		parsed, err := ParseValue(v, nextSensitive, addr)
		if err != nil {
			return nil, err
		}
		parsed.IncreaseDepth()
		parsed.AddEndingColon()
		out.AppendChild(parsed)
	}

	out.AppendEndNode("]")
	return out, nil
}
