package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func TestString(t *testing.T) {
	s := jsonString{render.BaseModel{Addr: "here"}, "string"}
	t.Run("View", func(t *testing.T) {
		params := render.ViewParams{ScreenWidth: 100, ScreenHeight: 100}
		view := s.View(params)
		if len(view) != 1 {
			t.Errorf("expected a view of length 1, got %d", len(view))
		}
		if !strings.Contains(view[0].String(), "\"string\"") {
			t.Errorf("expected \"string\", got %s", view)
		}
	})
}
