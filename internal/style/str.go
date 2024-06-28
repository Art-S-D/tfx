package style

import (
	"slices"
	"strings"
)

type Style uint8

type str struct {
	value      string
	style      Style
	selectable bool
}
type Str []str

const (
	defaultStyle = Style(iota)
	typeStyle
	key
	stringStyle
	boolean
	number
	null
	preview
	cursor
	selection
	sensitive
)

func (s str) String() string {
	return s.value
}
func (s Str) String() string {
	if s == nil {
		return ""
	}
	var sb strings.Builder
	for _, str := range s {
		sb.WriteString(str.String())
	}
	return sb.String()
}
func (s str) Selectable() str {
	s.selectable = true
	return s
}
func (s Str) setSelectable(selectable bool) Str {
	if s == nil {
		return nil
	}
	out := slices.Clone(s)
	for i := range out {
		out[i].selectable = selectable
	}
	return out
}
func (s Str) Selectable() Str {
	return s.setSelectable(true)
}
func (s Str) UnSelectable() Str {
	return s.setSelectable(false)
}
func (s Str) Concat(other Str) Str {
	if s == nil {
		return other
	}
	if other == nil {
		return s
	}
	return slices.Concat(s, other)
}

func Concat(s ...Str) Str {
	return slices.Concat(s...)
}

func makeStrFromStyle(style Style, strs ...string) Str {
	s := str{value: strings.Join(strs, ""), style: style, selectable: false}
	return Str([]str{s})
}

func Default(s ...string) Str {
	return makeStrFromStyle(defaultStyle, s...)
}
func Type(s ...string) Str {
	return makeStrFromStyle(typeStyle, s...)
}
func Key(s ...string) Str {
	return makeStrFromStyle(key, s...)
}
func String(s ...string) Str {
	return makeStrFromStyle(stringStyle, s...)
}
func Boolean(s ...string) Str {
	return makeStrFromStyle(boolean, s...)
}
func Number(s ...string) Str {
	return makeStrFromStyle(number, s...)
}
func Null(s ...string) Str {
	return makeStrFromStyle(null, s...)
}
func Preview(s ...string) Str {
	return makeStrFromStyle(preview, s...)
}
func Cursor(s ...string) Str {
	return makeStrFromStyle(cursor, s...)
}
func Selection(s ...string) Str {
	return makeStrFromStyle(selection, s...)
}
func Sensitive(s ...string) Str {
	return makeStrFromStyle(sensitive, s...)
}

func Space() Str {
	return Default(" ")
}
func Spaces(count int) Str {
	return Default(strings.Repeat(" ", count))
}
