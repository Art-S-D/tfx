package render

import (
	"testing"

	"github.com/Art-S-D/tfx/internal/style"
)

func TestIndent(t *testing.T) {
	line1 := Line{}
	line1.AddSelectable(style.Default("{"))

	line2 := Line{}
	line2.AddSelectable(style.Default("  \"a\": 2"))

	line3 := Line{}
	line3.AddSelectable(style.Default("}"))

	lines := []Line{line1, line2, line3}

	if lines[0].Indentation != 0 {
		t.Errorf("first line should not be indented")
	}

	Indent(lines)

	if lines[0].Indentation != 1 {
		t.Errorf("first line should be indented by 1")
	}
}
