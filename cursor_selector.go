package main

type CursorSelector interface {
	Address() string
	Toggle()
}
