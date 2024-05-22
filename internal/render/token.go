package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
)

type TokenType int

const (
	NewLine = TokenType(iota)
	Selectable
	Unselectable
	Indentation
)

type Token struct {
	Type  TokenType
	Value style.Str
}

func NewIndentationToken(indent int) Token {
	return Token{
		Type:  Indentation,
		Value: style.Default(strings.Repeat(" ", indent)),
	}
}

type TokenSlice struct {
	Start, End *TokenList
}
type TokenList struct {
	Prev, Next *TokenList
	Value      Token
}

func (t *TokenList) Append(tokens ...Token) {
	tail := t
	for _, tok := range tokens {
		elem := &TokenList{Next: tail.Next, Prev: tail, Value: tok}
		if tail.Next != nil {
			tail.Next.Prev = elem
			tail.Next = elem
		}
		tail = elem
	}
}
func (t *TokenList) AppendSelectable(str ...style.Str) {
	for _, s := range str {
		t.Append(Token{Type: Selectable, Value: s})
	}
}
func (t *TokenList) AppendUnselectable(str ...style.Str) {
	for _, s := range str {
		t.Append(Token{Type: Unselectable, Value: s})
	}
}
func (t *TokenList) AppendNewLine() {
	t.Append(Token{Type: NewLine})
}

func (t *TokenList) NextLine() *TokenList {
	current := t
	for current != nil && current.Value.Type != NewLine {
		current = current.Next
	}
	return current
}
func (t *TokenList) PreviousLine() *TokenList {
	current := t
	for current != nil && current.Value.Type != NewLine {
		current = current.Prev
	}
	return current
}
