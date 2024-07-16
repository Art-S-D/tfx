package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
)

func jsonArrayNode(address string, array []any, sensitiveValues any) (*node.Node, error) {
	if len(array) == 0 {
		out := node.StringNode("[]")
		out.Address = address
		return out, nil
	}

	out := &node.Node{
		Address:  address,
		Childer:  &node.DefaultChilder{},
		Expander: &node.DefaultExpander{},
		Renderer: &node.LineRenderer{
			Expanded: style.Default("[").Selectable(),
			Collapsed: style.Concat(
				style.Default("["),
				style.Preview("..."),
				style.Default("]"),
			).Selectable(),
		},
	}

	if _, ok := sensitiveValues.(bool); ok {
		out.Sensitive = true
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
		parsed.Indent()
		// parsed.AddEndingColon() // TODO
		out.AppendChild(parsed)
	}

	lastChild := node.StringNode("]")
	lastChild.Address = address
	out.AppendChild(lastChild)

	return out, nil
}
