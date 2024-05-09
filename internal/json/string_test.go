package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func TestString(t *testing.T) {
	s := jsonString{value: "string", address: "here"}
	t.Run("View", func(t *testing.T) {
		r := render.NewRenderer(0, 0, 10, 10)
		s.View(r)
		view := r.String()
		if !strings.Contains(view, "\"string\"") {
			t.Errorf("expected \"string\", got %s", view)
		}
	})
}
