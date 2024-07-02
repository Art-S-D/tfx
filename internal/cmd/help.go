package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// missing keybindings from fx
// 1-9                 collapse to nth level
// z                   toggle string wrap
// y                   yank/copy

var helpText string = `
    Key Bindings
        q, ctrl+c, exc      exit program
        pgdown, space, f    page down
        pgup, b             page up
        u, ctrl+u           half page up
        d, ctrl+d    	    half page down
        g, home             go to top
        G, end              go to bottom
        down, j             down
        up, k               up
        ?                   show help
        right, l, enter     expand
        left, h, backspace  collapse
        L, shift+right      expand recursively
        H, shift+left       collapse recursively
        e                   expand all
        E                   collapse all
        J, shift+down       next sibling
        K, shift+up         previous sibling
        r                   reveal value
        /                   search regexp
        n                   next search result
        N                   prev search result
        p                   preview
        P                   print
        .                   dig
`

func (m *TfxModel) helpScreen() string {
	help := lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, helpText)

	lastLine := ": press q or ? to close help"
	lastLine = lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Left, lastLine)
	lastLine = m.theme.Selection(lastLine).String()

	return fmt.Sprintf("%s\n%s", help, lastLine)
}

func (m *TfxModel) updateHelpView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "?", "q", "esc":
			m.state = viewState
		}
	}
	return m, nil
}
