package json

import (
	"fmt"
	"reflect"

	"github.com/Art-S-D/tfx/internal/render"
)

func ParseValue(jsonValue any, sensitiveValues any, parent *render.Node, address string) (*render.Node, error) {
	out := &render.Node{
		Address: address,
		Parent:  parent,
		Depth:   parent.Depth + 1,
	}
	if b, ok := sensitiveValues.(bool); ok && b {
		result, err := ParseValue(jsonValue, nil, out, address)
		out.Liner = &SensitiveValue{value: result}
		return out, err
	}

	switch value := jsonValue.(type) {
	case string:
		out.Liner = &jsonString{value}
		return out, nil
	case float64:
		out.Liner = &jsonNumber{value}
		return out, nil
	case bool:
		out.Liner = &jsonBool{value}
		return out, nil
	case nil:
		out.Liner = &jsonNull{}
		return out, nil
	case []any:
		array := &jsonArray{nil}

		sensitive, ok := sensitiveValues.([]any)
		if sensitiveValues != nil && !ok {
			return nil, fmt.Errorf("failed to parse sensitive value to array %v for json %v", sensitiveValues, value)
		}

		for i, v := range value {
			addr := fmt.Sprintf("%s[%d]", address, i)
			var nextSensitive any
			if sensitiveValues != nil {
				nextSensitive = sensitive[i]
			}
			parsed, err := ParseValue(v, nextSensitive, out, addr)
			if err != nil {
				return nil, err
			}
			array.value = append(array.value, parsed)
		}
		out.Liner = array
		return out, nil
	case map[string]any:
		object := &jsonObject{make(map[string]*render.Node)}

		sensitive, ok := sensitiveValues.(map[string]any)
		if sensitiveValues != nil && !ok {
			return nil, fmt.Errorf("failed to parse sensitive value to object %v", sensitiveValues)
		}

		for k, v := range value {
			addr := fmt.Sprintf("%s.%s", address, k)
			var nextSensitive any
			if sensitiveValues != nil {
				nextSensitive = sensitive[k]
			}
			parsed, err := ParseValue(v, nextSensitive, out, addr)
			if err != nil {
				return nil, err
			}
			object.value[k] = parsed
		}
		out.Liner = object
		return out, nil
	default:
		jsonType := reflect.TypeOf(jsonValue)
		return nil, fmt.Errorf("unknown json value %v of type %v", jsonValue, jsonType)
	}
}
