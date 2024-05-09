package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func TestSensitiveValue(t *testing.T) {
	t.Run("Expand", func(t *testing.T) {
		value := SensitiveValue{value: &jsonNull{}}
		value.Expand()
		if value.shown {
			t.Errorf("sensitive value should not expand")
		}
	})

	t.Run("Reveal", func(t *testing.T) {
		value := SensitiveValue{value: &jsonNull{}}
		params := render.ViewParams{ScreenWidth: 100, ScreenHeight: 100}

		reprBeforeReveal := value.View(&params)

		value.Reveal()

		reprAfterReveal := value.View(&params)

		if !strings.Contains(reprBeforeReveal, "(sensitive)") {
			t.Errorf(
				"expected sensitive value to be marked as sensitive, go %s",
				reprBeforeReveal,
			)
		}
		if reprAfterReveal == reprBeforeReveal {
			t.Errorf(
				"expected representation before and after reveal to be different but got %s %s",
				reprBeforeReveal,
				reprAfterReveal,
			)
		}
	})
}
