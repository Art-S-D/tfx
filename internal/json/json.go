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
		out := node.Str(style.String(s).Selectable())
		out.SetAddress(address)
		out.SetSensitive(isSensitive)
		return out, nil
	case float64:
		out := node.Str(style.Number(fmt.Sprintf("%.2f", value)).Selectable())
		out.SetAddress(address)
		out.SetSensitive(isSensitive)
		return out, nil
	case bool:
		out := node.Str(style.Boolean(fmt.Sprintf("%v", value)).Selectable())
		out.SetAddress(address)
		out.SetSensitive(isSensitive)
		return out, nil
	case nil:
		out := node.Str(style.Null("null").Selectable())
		out.SetAddress(address)
		out.SetSensitive(isSensitive)
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
