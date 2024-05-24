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
		return &jsonString{render.BaseModel{Addr: address}, value}, nil
	case float64:
		return &jsonNumber{render.BaseModel{Addr: address}, value}, nil
	case bool:
		return &jsonBool{render.BaseModel{Addr: address}, value}, nil
	case nil:
		return &jsonNull{render.BaseModel{Addr: address}}, nil
	case []any:
		array := &jsonArray{render.BaseModel{Addr: address}, render.BaseCollapser{Expanded: false}, nil}

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
			array.value = append(array.value, parsed)
		}
		return array, nil
	case map[string]any:
		object := &jsonObject{render.BaseModel{Addr: address}, render.BaseCollapser{Expanded: false}, make(map[string]render.Model)}

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
			object.value[k] = parsed
		}
		return object, nil
	default:
		jsonType := reflect.TypeOf(jsonValue)
		return nil, fmt.Errorf("unknown json value %v of type %v", jsonValue, jsonType)
	}
}
