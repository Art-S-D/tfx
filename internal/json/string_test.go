package json

import (
	"strings"
	"testing"

	"github.com/Art-S-D/tfx/internal/render"
)

func TestString(t *testing.T) {
	s := jsonString{value: "string", address: "here"}
	t.Run("View", func(t *testing.T) {
		params := &render.ViewParams{ScreenWidth: 100, ScreenHeight: 100}
		view := s.View(params)
		if !strings.Contains(view, "\"string\"") {
			t.Errorf("expected \"string\", got %s", view)
		}
	})
}
