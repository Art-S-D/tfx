package node

import (
	"testing"

	"github.com/Art-S-D/tfx/internal/style"
)

func makeNode() *Node {
	out := &Node{}
	out.address = "root"
	out.collapsed = style.Default("{...}")
	out.expanded = style.Default("{")
	out.isExpanded = true

	child1 := String("value")
	child1.key = "key1"
	out.AppendChild(child1)

	child2 := String("12")
	child2.key = "key1"
	out.AppendChild(child2)

	lastChild := String("}")
	out.AppendChild(lastChild)
	return out
}

func TestNode(t *testing.T) {
	t.Run("Expand does not reveal a sensitive value", func(t *testing.T) {
		node := makeNode()
		node.sensitive = true
		node.Expand()
		if !node.sensitive {
			t.Errorf("the node should be sensitive")
		}
	})

	t.Run("IncreaseDepth", func(t *testing.T) {
		t.Run("Does not crash if there are no children", func(t *testing.T) {
			node := String("testnode")
			node.IncreaseDepth()
		})
		t.Run("Increase the depth of the children", func(t *testing.T) {
			node := makeNode()
			if node.depth != 0 || node.children[0].depth != 0 {
				t.Errorf("wrong struct initialization")
			}
			node.IncreaseDepth()
			if node.depth != 1 || node.children[0].depth != 1 {
				t.Errorf("depth should increase")
			}
		})
	})
}
