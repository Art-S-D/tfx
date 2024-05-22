package style

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	defaultStyle lipgloss.Style
	typeStyle    lipgloss.Style
	key          lipgloss.Style
	stringStyle  lipgloss.Style
	boolean      lipgloss.Style
	number       lipgloss.Style
	null         lipgloss.Style
	preview      lipgloss.Style
	cursor       lipgloss.Style
	selection    lipgloss.Style
}

func (t *Theme) Default(s ...string) lipgloss.Style {
	return t.defaultStyle.SetString(strings.Join(s, ""))
}
func (t *Theme) Type(s ...string) lipgloss.Style {
	return t.typeStyle.SetString(strings.Join(s, ""))
}
func (t *Theme) Key(s ...string) lipgloss.Style {
	return t.key.SetString(strings.Join(s, ""))
}
func (t *Theme) String(s ...string) lipgloss.Style {
	return t.stringStyle.SetString(strings.Join(s, ""))
}
func (t *Theme) Boolean(s ...string) lipgloss.Style {
	return t.boolean.SetString(strings.Join(s, ""))
}
func (t *Theme) Number(s ...string) lipgloss.Style {
	return t.number.SetString(strings.Join(s, ""))
}
func (t *Theme) Null(s ...string) lipgloss.Style {
	return t.null.SetString(strings.Join(s, ""))
}
func (t *Theme) Preview(s ...string) lipgloss.Style {
	return t.preview.SetString(strings.Join(s, ""))
}
func (t *Theme) Cursor(s ...string) lipgloss.Style {
	return t.cursor.SetString(strings.Join(s, ""))
}
func (t *Theme) Selection(s ...string) lipgloss.Style {
	return t.selection.SetString(strings.Join(s, ""))
}

var DefaultTheme *Theme
var NoColor *Theme

func init() {
	DefaultTheme = &Theme{
		defaultStyle: lipgloss.NewStyle(),
		key:          lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("4")),
		stringStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
		boolean:      lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		number:       lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		null:         lipgloss.NewStyle().Foreground(lipgloss.Color("243")),
		typeStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		preview:      lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
		cursor:       lipgloss.NewStyle().Reverse(true),
		selection:    lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("0")),
	}

	NoColor = &Theme{
		key:         lipgloss.NewStyle(),
		stringStyle: lipgloss.NewStyle(),
		boolean:     lipgloss.NewStyle(),
		number:      lipgloss.NewStyle(),
		null:        lipgloss.NewStyle(),
		typeStyle:   lipgloss.NewStyle(),
		preview:     lipgloss.NewStyle(),
		cursor:      lipgloss.NewStyle(),
		selection:   lipgloss.NewStyle(),
	}
}

func (t *Theme) Render(str Str) string {
	switch str.Style {
	case defaultStyle:
		return t.defaultStyle.Render(str.Value)
	case typeStyle:
		return t.typeStyle.Render(str.Value)
	case key:
		return t.key.Render(str.Value)
	case stringStyle:
		return t.stringStyle.Render(str.Value)
	case boolean:
		return t.boolean.Render(str.Value)
	case number:
		return t.number.Render(str.Value)
	case null:
		return t.null.Render(str.Value)
	case preview:
		return t.preview.Render(str.Value)
	case cursor:
		return t.cursor.Render(str.Value)
	case selection:
		return t.selection.Render(str.Value)
	}
	panic(fmt.Errorf("unknown style %v", str))
}

func (t *Theme) RenderCursor(selected bool, str Str) string {
	if selected {
		return t.cursor.Render(str.Value)
	} else {
		return t.Render(str)
	}
}
