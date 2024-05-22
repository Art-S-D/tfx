package style

import "strings"

type Style uint8

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
)

func Default(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: defaultStyle}
}
func Type(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: typeStyle}
}
func Key(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: key}
}
func String(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: stringStyle}
}
func Boolean(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: boolean}
}
func Number(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: number}
}
func Null(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: null}
}
func Preview(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: preview}
}
func Cursor(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: cursor}
}
func Selection(s ...string) Str {
	return Str{Value: strings.Join(s, ""), Style: selection}
}

type Str struct {
	Value string
	Style Style
}
