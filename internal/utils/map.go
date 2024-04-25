package utils

import (
	"slices"

	"golang.org/x/exp/maps"
)

func KeysOrdered(m map[string]any) []string {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}
