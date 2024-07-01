package node

import "testing"

func TestSiblings(t *testing.T) {
	t.Run("test after", func(t *testing.T) {
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			after := node.after(node.children[0])
			if after != node.children[1] {
				t.Errorf("wrong after child")
			}
		})
		t.Run("last child", func(t *testing.T) {
			node := makeNode()
			node.key = "aa"
			parent := String("test")
			parent.AppendChild(node)
			secondChild := String("test2")
			parent.AppendChild(secondChild)
			after := node.after(node.children[len(node.children)-1])
			if after != secondChild {
				t.Errorf("wrong child")
			}
		})

		t.Run("no parent", func(t *testing.T) {
			node := makeNode()
			after := node.after(node.children[len(node.children)-1])
			if after != nil {
				t.Errorf("wrong child")
			}
		})
	})

	t.Run("test Next", func(t *testing.T) {
		t.Run("test collapsed", func(t *testing.T) {
			node := makeNode()
			child := node.children[0]
			child.AppendChild(String("aaa"))
			child.Collapse()
			nextNext := child.Next()
			if nextNext != node.children[1] {
				t.Errorf("wrong next node %+v", nextNext)
			}
		})
		t.Run("test children", func(t *testing.T) {
			node := makeNode()
			next := node.Next()
			if next != node.children[0] {
				t.Errorf("wrong child")
			}
		})
		t.Run("test no child", func(t *testing.T) {
			node := makeNode()
			child := node.children[0]
			next := child.Next()
			if next != node.children[1] {
				t.Errorf("wrong child")
			}
		})
		t.Run("test last child", func(t *testing.T) {
			node := makeNode()
			lastChild := node.children[len(node.children)-1]
			next := lastChild.Next()
			if next != nil {
				t.Errorf("wrong child")
			}
		})
	})

	t.Run("LastChild", func(t *testing.T) {
		t.Run("no child", func(t *testing.T) {
			node := String("aaa")
			last := node.LastChild()
			if last != node {
				t.Errorf("wrong last child")
			}
		})
		t.Run("collapsed", func(t *testing.T) {
			node := makeNode()
			node.Collapse()
			last := node.LastChild()
			if last != node {
				t.Errorf("wrong last child")
			}
		})
		t.Run("with children", func(t *testing.T) {
			node := makeNode()
			last := node.LastChild()
			if last != node.children[len(node.children)-1] {
				t.Errorf("wrong last child")
			}
		})

		t.Run("depth of two", func(t *testing.T) {
			node := makeNode()
			last := node.LastChild()
			parent := String("aaa")
			parent.Expand()
			parent.AppendChild(node)
			if parent.LastChild() != last {
				t.Errorf("wrong last child")
			}
		})
	})

	t.Run("before", func(t *testing.T) {
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			before := node.before(node.children[0])
			if before != node {
				t.Errorf("wrong child")
			}
		})

		t.Run("any child", func(t *testing.T) {
			node := makeNode()
			before := node.before(node.children[1])
			if before != node.children[0] {
				t.Errorf("wrong child")
			}
		})

		t.Run("depth of two", func(t *testing.T) {
			node1 := makeNode()
			node2 := makeNode()
			parent := String("aaa")
			parent.AppendChild(node1)
			parent.AppendChild(node2)

			before := parent.before(node2)
			if before != node1.children[len(node1.children)-1] {
				t.Errorf("wrong child")
			}
		})
	})

	t.Run("Previous", func(t *testing.T) {
		t.Run("no parent", func(t *testing.T) {
			node := makeNode()
			parent := node.Previous()
			if parent != nil {
				t.Errorf("wrong previous")
			}
		})
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			previous := node.children[0].Previous()
			if previous != node {
				t.Errorf("wrong previous")
			}
		})
		t.Run("any child", func(t *testing.T) {
			node := makeNode()
			previous := node.children[1].Previous()
			if previous != node.children[0] {
				t.Errorf("wrong previous")
			}
		})
	})
}
