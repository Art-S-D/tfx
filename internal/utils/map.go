package utils

import (
	"slices"

	"golang.org/x/exp/maps"
)

func KeysOrdered[T any](m map[string]T) []string {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}
