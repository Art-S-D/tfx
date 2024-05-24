package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

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
        1-9                 collapse to nth level
        J, shift+down       next sibling
        K, shift+up         previous sibling
        z                   toggle string wrap
        y                   yank/copy
        /                   search regexp
        n                   next search result
        N                   prev search result
        p                   preview
        P                   print
        .                   dig
`

func (m *stateModel) helpScreen() string {
	help := lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, helpText)

	lastLine := ": press q or ? to close help"
	lastLine = lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Left, lastLine)
	lastLine = m.theme.Selection(lastLine).String()

	return fmt.Sprintf("%s\n%s", help, lastLine)
}
