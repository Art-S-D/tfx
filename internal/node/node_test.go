package node

import (
	"testing"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/stretchr/testify/assert"
)

func makeNode() *Node {
	out := &Node{}
	out.address = "root"
	out.collapsed = style.Default("{...}")
	out.expanded = style.Default("{")
	out.isExpanded = true

	child1 := String("value")
	child1.key = "key1"
	child1.depth = 1
	out.AppendChild(child1)

	child2 := String("12")
	child2.key = "key1"
	child2.depth = 1
	out.AppendChild(child2)

	out.AppendEndNode("}")
	return out
}

func TestNode(t *testing.T) {
	assert := assert.New(t)

	t.Run("Expand does not reveal a sensitive value", func(t *testing.T) {
		node := makeNode()
		node.sensitive = true
		node.Expand()
		assert.True(node.sensitive)
	})

	t.Run("IncreaseDepth", func(t *testing.T) {
		t.Run("Does not crash if there are no children", func(t *testing.T) {
			assert.NotPanics(func() {
				node := String("testnode")
				node.IncreaseDepth()
			})
		})
		t.Run("Increase the depth of the children", func(t *testing.T) {
			node := makeNode()

			assert.EqualValues(0, node.depth)
			assert.EqualValues(1, node.children[0].depth)

			node.IncreaseDepth()

			assert.EqualValues(1, node.depth)
			assert.EqualValues(2, node.children[0].depth)
		})
	})

	t.Run("Target", func(t *testing.T) {
		node := makeNode()
		lastChild := node.children[len(node.children)-1]

		assert.Same(node, lastChild.Target())

		assert.True(node.isExpanded)
		lastChild.Collapse()
		assert.False(node.isExpanded)
	})
}
