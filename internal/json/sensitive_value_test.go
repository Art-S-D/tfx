package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func linesToString(lines []render.Line) string {
	var out strings.Builder
	for _, l := range lines {
		out.WriteString(l.String())
		out.WriteRune('\n')
	}
	return out.String()
}

func TestSensitiveValue(t *testing.T) {
	t.Run("Reveal", func(t *testing.T) {
		value := SensitiveValue{value: &jsonNull{}}

		reprBeforeReveal := linesToString(value.View())
		value.Reveal()
		reprAfterReveal := linesToString(value.View())

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
