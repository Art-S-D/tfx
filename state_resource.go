package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	"github.com/Art-S-D/tfview/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	tfjson "github.com/hashicorp/terraform-json"
)

type StateResourceModel struct {
	resource tfjson.StateResource
	expanded bool
}

func (m *StateResourceModel) Height() int {
	if !m.expanded {
		return 1
	} else {
		return 2 + len(m.resource.AttributeValues)
	}
}

func (m *StateResourceModel) Address() string {
	return m.resource.Address
}
func (m *StateResourceModel) Toggle() {
	m.expanded = !m.expanded
}

func (m *StateResourceModel) Init() tea.Cmd {
	return nil
}

func (m *StateResourceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			m.expanded = !m.expanded
			return m, nil
		}

	}
	return m, nil
}

func (m *StateResourceModel) ViewCursor(cursor int) (out string, newCursor int) {
	var sb strings.Builder

	resourceMode := "resource"
	if m.resource.Mode == tfjson.DataResourceMode {
		resourceMode = "data"
	}

	if cursor == 0 {
		sb.WriteString(
			style.Cursor.Render(
				resourceMode,
				fmt.Sprintf("\"%s\"", m.resource.Type),
				fmt.Sprintf("\"%s\"", m.resource.Name),
			),
		)
	} else {
		sb.WriteString(
			fmt.Sprintf(
				"%s %s %s",
				style.Type.Render(resourceMode),
				style.Key.Render(fmt.Sprintf("\"%s\"", m.resource.Type)),
				style.Key.Render(fmt.Sprintf("\"%s\"", m.resource.Name)),
			),
		)
	}
	cursor -= 1

	if !m.expanded {
		sb.WriteString(" {")
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String(), cursor
	} else {
		sb.WriteString(" {\n")
	}

	keys := utils.KeysOrdered(m.resource.AttributeValues)
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.resource.AttributeValues[k]
		sb.WriteString("  ")

		usedStyle := style.Key
		if cursor == 0 {
			usedStyle = style.Cursor
		}
		renderedKey := usedStyle.Copy().MarginRight(len(longestKey) - len(k)).Render(k)

		sb.WriteString(renderedKey)

		sb.WriteString(" = ")
		sb.WriteString(style.Render(v))
		sb.WriteString("\n")
		cursor -= 1
	}

	if cursor == 0 {
		sb.WriteString(style.Cursor.Render("}"))
	} else {
		sb.WriteRune('}')
	}
	cursor -= 1
	return sb.String(), cursor
}

func (m *StateResourceModel) View() string {
	var sb strings.Builder

	sb.WriteString(
		fmt.Sprintf(
			"%s %s %s",
			style.Type.Render("resource"),
			style.Key.Render(fmt.Sprintf("\"%s\"", m.resource.Type)),
			style.Key.Render(fmt.Sprintf("\"%s\"", m.resource.Name)),
		),
	)

	if !m.expanded {
		sb.WriteString(" {")
		sb.WriteString(style.Preview.Render("..."))
		sb.WriteString("}")
		return sb.String()
	} else {
		sb.WriteString(" {\n")
	}

	keys := utils.KeysOrdered(m.resource.AttributeValues)
	longestKey := slices.MaxFunc(keys, func(s1, s2 string) int { return len(s1) - len(s2) })
	for _, k := range keys {
		v := m.resource.AttributeValues[k]
		renderedKey := style.Key.Copy().MarginRight(len(longestKey) - len(k)).Render(k)
		sb.WriteString("  ")
		sb.WriteString(renderedKey)
		sb.WriteString(" = ")
		sb.WriteString(style.Render(v))
		sb.WriteString("\n")
	}
	sb.WriteRune('}')
	return sb.String()
}
