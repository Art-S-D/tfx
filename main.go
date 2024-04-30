package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	tfjson "github.com/hashicorp/terraform-json"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	offset                    int // should always be between [0, rootModule.Height() - screenHeight)
	rootModule                StateModuleModel
	rootModuleHeight          int
}

func (m *stateModel) clampOffset() {
	if m.offset < 0 {
		m.offset = 0
	}
	if m.offset >= m.rootModuleHeight-m.screenHeight {
		m.offset = m.rootModuleHeight - m.screenHeight - 1
	}
}

func (m *stateModel) screenBottom() int {
	// should be -1 but we have -2 instead to account for the preview
	return m.screenHeight + m.offset - 2
}

func (m *stateModel) cursorUp() {
	m.cursor -= 1
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor <= m.offset+3 {
		m.offset -= 1
	}
	m.clampOffset()
	m.clampCursor()
}

func (m *stateModel) cursorDown() {
	m.cursor += 1
	screenBottom := m.screenBottom()
	if m.cursor >= screenBottom {
		m.cursor = screenBottom
	}
	if m.cursor >= screenBottom-3 {
		m.offset += 1
	}
	m.clampOffset()
	m.clampCursor()
}

func (m *stateModel) clampCursor() {
	if m.cursor < 0 {
		m.cursor = 0
	}
	screenBottom := m.screenBottom()
	if m.cursor >= screenBottom {
		m.cursor = screenBottom
	}
}

func (m *stateModel) Init() tea.Cmd {
	m.rootModuleHeight = m.rootModule.RenderingHeight()
	return nil
}
func (m *stateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
			m.cursorUp()
			return m, nil
		case "down":
			m.cursorDown()
			return m, nil
		case "enter":
			selected, selectedLine := m.rootModule.Selected(m.cursor)
			selected.Toggle()
			m.cursor -= selectedLine
			m.rootModuleHeight = m.rootModule.RenderingHeight()
			m.clampOffset()
			m.clampCursor()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampOffset()
	}
	return m, nil
}

func (m *stateModel) View() string {
	view, _ := m.rootModule.View(m.cursor)
	lines := strings.Split(view, "\n")
	linesInView := lines[m.offset : m.offset+m.screenHeight]
	if len(linesInView) > 1 {
		selection, _ := m.rootModule.Selected(m.cursor)
		selectionName := selection.Address()

		helpText := "[?]help [q]quit "
		previewLine := fmt.Sprintf(
			"%s%s%s",
			selectionName,
			strings.Repeat(" ", m.screenWidth-len(selectionName)-len(helpText)),
			helpText,
		)

		linesInView[len(linesInView)-1] = style.Selection.Copy().
			PaddingRight(m.screenWidth - len(selectionName)).
			Render(previewLine)
	}
	return strings.Join(linesInView, "\n")
}

func main() {
	args := parseArgs()

	jsonState, err := io.ReadAll(args.src)
	if err != nil {
		panic(err.Error())
	}
	var plan tfjson.State
	err = json.Unmarshal(jsonState, &plan)
	if err != nil {
		panic(fmt.Errorf("failed to read json state %w", err))
	}

	terraformState := stateModel{rootModule: StateModuleModelFromJson(*plan.Values.RootModule)}
	terraformState.rootModule.expanded = true
	p := tea.NewProgram(&terraformState)
	if _, err := p.Run(); err != nil {
		panic(err.Error())
	}
}
