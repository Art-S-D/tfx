package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSiblings(t *testing.T) {
	assert := assert.New(t)
	t.Run("test after", func(t *testing.T) {
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			after := node.after(node.children[0])
			assert.Same(after, node.children[1])
		})
		t.Run("last child", func(t *testing.T) {
			node := makeNode()
			node.key = "aa"
			parent := String("test")
			parent.AppendChild(node)
			secondChild := String("test2")
			parent.AppendChild(secondChild)
			after := node.after(node.children[len(node.children)-1])

			assert.Same(after, secondChild)
		})

		t.Run("no parent", func(t *testing.T) {
			node := makeNode()
			after := node.after(node.children[len(node.children)-1])
			assert.Nil(after)
		})
	})

	t.Run("test Next", func(t *testing.T) {
		t.Run("test collapsed", func(t *testing.T) {
			node := makeNode()
			child := node.children[0]
			child.AppendChild(String("aaa"))
			child.Collapse()
			nextNext := child.Next()
			assert.Same(nextNext, node.children[1])
		})
		t.Run("test children", func(t *testing.T) {
			node := makeNode()
			next := node.Next()
			assert.Same(next, node.children[0])
		})
		t.Run("test no child", func(t *testing.T) {
			node := makeNode()
			child := node.children[0]
			next := child.Next()
			assert.Same(next, node.children[1])
		})
		t.Run("test last child", func(t *testing.T) {
			node := makeNode()
			lastChild := node.children[len(node.children)-1]
			next := lastChild.Next()
			assert.Nil(next)
		})
	})

	t.Run("LastChild", func(t *testing.T) {
		t.Run("no child", func(t *testing.T) {
			node := String("aaa")
			last := node.LastChild()
			assert.Same(last, node)
		})
		t.Run("collapsed", func(t *testing.T) {
			node := makeNode()
			node.Collapse()
			last := node.LastChild()
			assert.Same(last, node)
		})
		t.Run("with children", func(t *testing.T) {
			node := makeNode()
			last := node.LastChild()
			assert.Same(last, node.children[len(node.children)-1])
		})

		t.Run("depth of two", func(t *testing.T) {
			node := makeNode()
			last := node.LastChild()
			parent := String("aaa")
			parent.Expand()
			parent.AppendChild(node)
			assert.Same(parent.LastChild(), last)
		})
	})

	t.Run("before", func(t *testing.T) {
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			before := node.before(node.children[0])
			assert.Same(before, node)
		})

		t.Run("any child", func(t *testing.T) {
			node := makeNode()
			before := node.before(node.children[1])
			assert.Same(before, node.children[0])
		})

		t.Run("depth of two", func(t *testing.T) {
			node1 := makeNode()
			node2 := makeNode()
			parent := String("aaa")
			parent.AppendChild(node1)
			parent.AppendChild(node2)

			before := parent.before(node2)
			assert.Same(before, node1.children[len(node1.children)-1])
		})
	})

	t.Run("Previous", func(t *testing.T) {
		t.Run("no parent", func(t *testing.T) {
			node := makeNode()
			parent := node.Previous()
			assert.Nil(parent)
		})
		t.Run("first child", func(t *testing.T) {
			node := makeNode()
			previous := node.children[0].Previous()
			assert.Same(previous, node)
		})
		t.Run("any child", func(t *testing.T) {
			node := makeNode()
			previous := node.children[1].Previous()
			assert.Same(previous, node.children[0])
		})
	})
}
