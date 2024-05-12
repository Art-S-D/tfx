package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func TestString(t *testing.T) {
	s := jsonString{value: "string", address: "here"}
	t.Run("View", func(t *testing.T) {
		view := render.LinesToString(s.Lines(0))
		if !strings.Contains(view, "\"string\"") {
			t.Errorf("expected \"string\", got %s", view)
		}
	})
}
