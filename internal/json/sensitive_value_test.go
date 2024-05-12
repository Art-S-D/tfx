package json

import (
	"testing"
)

func TestSensitiveValue(t *testing.T) {
	t.Run("Expand", func(t *testing.T) {
		value := SensitiveValue{value: &jsonNull{}}
		value.Expand()
		if value.shown {
			t.Errorf("sensitive value should not expand")
		}
	})
}
