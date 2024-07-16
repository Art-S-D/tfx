package json

import (
	"fmt"
	"reflect"

	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
)

func ParseValue(jsonValue any, sensitiveValues any, address string) (*node.Node, error) {
	isSensitive := false
	if b, ok := sensitiveValues.(bool); ok && b {
		isSensitive = true
	}

	switch value := jsonValue.(type) {
	case string:
		s := fmt.Sprintf("\"%s\"", value)
		out := node.StrNode(style.String(s).Selectable())
		out.Address = address
		out.Sensitive = isSensitive
		return out, nil
	case float64:
		out := node.StrNode(style.Number(fmt.Sprintf("%.2f", value)).Selectable())
		out.Address = address
		out.Sensitive = isSensitive
		return out, nil
	case bool:
		out := node.StrNode(style.Boolean(fmt.Sprintf("%v", value)).Selectable())
		out.Address = address
		out.Sensitive = isSensitive
		return out, nil
	case nil:
		out := node.StrNode(style.Null("null").Selectable())
		out.Address = address
		out.Sensitive = isSensitive
		return out, nil
	case []any:
		return jsonArrayNode(address, value, sensitiveValues)
	case map[string]any:
		return jsonObjectNode(address, value, sensitiveValues)
	default:
		jsonType := reflect.TypeOf(jsonValue)
		return nil, fmt.Errorf("unknown json value %v of type %v", jsonValue, jsonType)
	}
}
