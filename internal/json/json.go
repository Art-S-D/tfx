package json

import (
	"fmt"
	"reflect"

	"github.com/Art-S-D/tfx/internal/render"
)

func ParseValue(jsonValue any, address string) render.Model {
	switch value := jsonValue.(type) {
	case string:
		return &jsonString{value, address}
	case float64:
		return &jsonNumber{value, address}
	case bool:
		return &jsonBool{value, address}
	case nil:
		return &jsonNull{address}
	case []any:
		out := jsonArray{nil, address, false}
		for i, v := range value {
			addr := fmt.Sprintf("%s[%d]", address, i)
			parsed := ParseValue(v, addr)
			out.value = append(out.value, parsed)
		}
		return &out
	case map[string]any:
		out := jsonObject{make(map[string]render.Model), address, false}
		for k, v := range value {
			addr := fmt.Sprintf("%s.%s", address, k)
			parsed := ParseValue(v, addr)
			out.value[k] = parsed
		}
		return &out
	default:
		jsonType := reflect.TypeOf(jsonValue)
		panic(fmt.Sprintf("unknown json value %v of type %v", jsonValue, jsonType))
	}
}
