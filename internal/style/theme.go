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
	sensitive    lipgloss.Style
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
func (t *Theme) Sensitive(s ...string) lipgloss.Style {
	return t.sensitive.SetString(strings.Join(s, ""))
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
		sensitive:    lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
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
		sensitive:   lipgloss.NewStyle(),
	}
}

func (t *Theme) render(str str) string {
	switch str.style {
	case defaultStyle:
		return t.defaultStyle.Render(str.value)
	case typeStyle:
		return t.typeStyle.Render(str.value)
	case key:
		return t.key.Render(str.value)
	case stringStyle:
		return t.stringStyle.Render(str.value)
	case boolean:
		return t.boolean.Render(str.value)
	case number:
		return t.number.Render(str.value)
	case null:
		return t.null.Render(str.value)
	case preview:
		return t.preview.Render(str.value)
	case cursor:
		return t.cursor.Render(str.value)
	case selection:
		return t.selection.Render(str.value)
	case sensitive:
		return t.sensitive.Render(str.value)
	}
	panic(fmt.Errorf("unknown style %v", str))
}

func (t *Theme) Render(str Str) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(t.render(s))
	}
	return sb.String()
}

func (t *Theme) renderCursor(s str) string {
	if s.selectable {
		return t.cursor.Render(s.value)
	} else {
		return t.render(s)
	}
}

func (t *Theme) RenderWithCursor(str Str) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(t.renderCursor(s))
	}
	return sb.String()
}
