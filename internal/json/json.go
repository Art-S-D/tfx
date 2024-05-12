package json

import (
	"fmt"
	"reflect"

	"github.com/Art-S-D/tfx/internal/render"
)

func ParseValue(jsonValue any, sensitiveValues any, address string) (render.Model, error) {
	if b, ok := sensitiveValues.(bool); ok && b {
		result, err := ParseValue(jsonValue, nil, address)
		return &SensitiveValue{value: result}, err
	}

	switch value := jsonValue.(type) {
	case string:
		return &jsonString{value, address}, nil
	case float64:
		return &jsonNumber{value, address}, nil
	case bool:
		return &jsonBool{value, address}, nil
	case nil:
		return &jsonNull{address}, nil
	case []any:
		out := jsonArray{nil, address, false}

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
			parsed, err := ParseValue(v, nextSensitive, addr)
			if err != nil {
				return nil, err
			}
			out.value = append(out.value, parsed)
		}
		return &out, nil
	case map[string]any:
		out := jsonObject{make(map[string]render.Model), address, false}

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
			parsed, err := ParseValue(v, nextSensitive, addr)
			if err != nil {
				return nil, err
			}
			out.value[k] = parsed
		}
		return &out, nil
	default:
		jsonType := reflect.TypeOf(jsonValue)
		return nil, fmt.Errorf("unknown json value %v of type %v", jsonValue, jsonType)
	}
}
