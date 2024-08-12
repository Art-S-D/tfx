package preview

import "github.com/Art-S-D/tfx/internal/node"

func (m *PreviewModel) screenEnd() *node.Node {
	end := m.screenStart
	for i := 0; i < m.Ctx.ScreenHeight-2 && end.Next() != nil; i++ {
		end = end.Next()
	}
	return end
}

func (m *PreviewModel) clampScreenStart() {
	end := m.screenStart
	for i := 0; i < m.Ctx.ScreenHeight-2; i++ {
		next := end.Next()
		if next == nil {
			// screenEnd is at the end of the screen => we need to move the.screenStart up

			if m.screenStart == m.node {
				// screenEnd is at the end and.screenStart is at the start
				// => meaning the state is too small to cover the entire screen
				// so we can just stop
				return
			}
			m.screenStart = m.screenStart.Previous()
		} else {
			// otherwise just go down
			end = end.Next()
		}
	}
}

func (m *PreviewModel) moveScreenToCursor() {
	end := m.screenStart
	for i := 0; i < m.Ctx.ScreenHeight-2 && end.Next() != nil; i++ {
		if end == m.cursor {
			// the cursor is already in view => stop here
			return
		}
		end = end.Next()
	}

	// try to scroll down to the cursor
	current := end
	start := m.screenStart
	for current != nil {
		if current == m.cursor {
			m.screenStart = start
			return
		}
		current = current.Next()
		start = start.Next()
	}

	// try to scroll up to the cursor
	current = m.screenStart
	for current != nil {
		if current == m.cursor {
			m.screenStart = current
			return
		}
		current = current.Previous()
	}

	panic("cursor not found")
}
