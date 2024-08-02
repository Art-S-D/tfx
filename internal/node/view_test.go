package node

import (
	"testing"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	assert := assert.New(t)

	t.Run("AddEndingColon", func(t *testing.T) {
		node := makeNode()
		node.collapsed = style.Default("{...}")
		node.expanded = style.Default("{")
		node.AddEndingColon()

		assert.Equal("{...},", node.collapsed.String())
		assert.Equal("{", node.expanded.String())
		assert.Equal("},", node.children[len(node.children)-1].expanded.String())
	})

	t.Run("View", func(t *testing.T) {
		t.Run("collapsed", func(t *testing.T) {
			node := makeNode()
			node.isExpanded = false
			view := node.View().String()
			assert.Equal("{...}", view)
		})
		t.Run("expanded", func(t *testing.T) {
			node := makeNode()
			node.isExpanded = true
			view := node.View().String()
			assert.Equal("{", view)
		})
		t.Run("sensitive", func(t *testing.T) {
			node := makeNode()
			node.sensitive = true
			view := node.View().String()
			assert.Equal("(sensitive)", view)
		})
	})
}
