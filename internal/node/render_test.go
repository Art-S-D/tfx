package node

import (
	"testing"

	"github.com/Art-S-D/tfx/internal/style"
)

func TestView(t *testing.T) {
	t.Run("AddEndingColon", func(t *testing.T) {
		node := makeNode()
		node.collapsed = style.Default("{...}")
		node.expanded = style.Default("{")
		node.AddEndingColon()
		if node.collapsed.String() != "{...}," {
			t.Errorf("wrong collapsed view")
		}
		if node.expanded.String() != "{" {
			t.Errorf("wrong collapsed view")
		}
		if node.children[len(node.children)-1].expanded.String() != "}," {
			t.Errorf("wrong collapsed view")
		}
	})

	t.Run("View", func(t *testing.T) {
		t.Run("collapsed", func(t *testing.T) {
			node := makeNode()
			node.isExpanded = false
			view := node.View().String()
			if view != "{...}" {
				t.Errorf("wrong view %s", view)
			}
		})
		t.Run("expanded", func(t *testing.T) {
			node := makeNode()
			node.isExpanded = true
			view := node.View().String()
			if view != "{" {
				t.Errorf("wrong view %s", view)
			}
		})
		t.Run("sensitive", func(t *testing.T) {
			node := makeNode()
			node.sensitive = true
			view := node.View().String()
			if view != "(sensitive)" {
				t.Errorf("wrong view %s", view)
			}
		})
	})
}
