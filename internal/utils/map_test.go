package utils

import (
	"slices"
	"testing"
)

func TestMaps(t *testing.T) {
	testMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	keys := KeysOrdered(testMap)
	if !slices.Equal(keys, []string{"a", "b", "c"}) {
		t.Errorf("keys should be [a,b,c]")
	}

	testMap2 := map[string]int{
		"b": 2,
		"a": 1,
		"_": 3,
	}
	keys2 := KeysOrdered(testMap2)
	if !slices.Equal(keys2, []string{"_", "a", "b"}) {
		t.Errorf("keys should be [_,a,b]")
	}
}
