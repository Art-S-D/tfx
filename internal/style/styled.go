package style

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Art-S-D/tfview/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

var Default lipgloss.Style
var Type lipgloss.Style
var Key lipgloss.Style
var String lipgloss.Style
var Boolean lipgloss.Style
var Number lipgloss.Style
var Null lipgloss.Style
var Preview lipgloss.Style
var Indented lipgloss.Style
var Cursor lipgloss.Style
var Selection lipgloss.Style

func init() {
	Default = lipgloss.NewStyle()
	Key = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("4"))
	String = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	Boolean = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	Number = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	Null = lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
	Type = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	Preview = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	Indented = lipgloss.NewStyle().MarginLeft(2)
	Cursor = lipgloss.NewStyle().Reverse(true)
	Selection = lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("0"))
}

func Render(json any) string {
	if v, ok := json.(string); ok {
		return String.Render(fmt.Sprintf("\"%v\"", v))
	} else if v, ok := json.(bool); ok {
		return Boolean.Render(fmt.Sprintf("%v", v))
	} else if v, ok := json.(float64); ok {
		return Number.Render(fmt.Sprintf("%.2f", v))
	} else if json == nil {
		return Null.Render("null")
	} else if v, ok := json.(map[string]any); ok {
		var sb strings.Builder
		sb.WriteRune('{')
		keys := utils.KeysOrdered(v)
		for i, k := range keys {
			sb.WriteString(Key.Render(fmt.Sprintf("\"%v\"", k)))
			sb.WriteString("=")
			sb.WriteString(Render(v[k]))
			if i < len(keys)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune('}')
		return sb.String()
	} else if v, ok := json.([]any); ok {
		var sb strings.Builder
		sb.WriteRune('[')
		for i := range v {
			sb.WriteString(Render(v[i]))
			if i < len(v)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune(']')
		return sb.String()
	} else {
		jsonType := reflect.TypeOf(json)
		panic(fmt.Sprintf("unknown json value %v of type %v", json, jsonType))
	}
}

// renders the string with the style `usedStyle` of with the cursor style if the cursor value is 0
func RenderStyleOrCursor(cursor int, usedStyle lipgloss.Style, str string) string {
	if cursor == 0 {
		return Cursor.Render(str)
	} else {
		return usedStyle.Render(str)
	}
}
