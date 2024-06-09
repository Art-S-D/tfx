package render

import (
	"strings"
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

func TestIndentation(t *testing.T) {
	line := &Line{}
	line.AddSelectable(style.Default("abc"))
	line.Indentation = 2
	if strings.HasPrefix(line.String(), "  ") {
		t.Errorf("wrong indentation: got <%s> should be indented ny 2", line.String())
	}
}

func TestMerge(t *testing.T) {
	makeLines := func() (line *Line, other *Line) {
		other = &Line{}
		other.AddSelectable(style.Default("456"))
		line = &Line{}
		line.AddSelectable(style.Default("123-"))
		return line, other
	}
	t.Run("concatenation", func(t *testing.T) {
		line, other := makeLines()

		result := line.MergeWith(other)
		if result.String() != "123-456" {
			t.Errorf("wrong line merge %s", result.String())
		}
	})
}
