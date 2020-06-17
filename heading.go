package html

import (
	"fmt"
)

type heading struct {
	Element
	Level int
}

func Heading(attrs ...func(*heading)) HtmlElement {
	h := &heading{}
	for _, attr := range attrs {
		attr(h)
	}
	h.tag = fmt.Sprintf("h%d", h.Level)
	return h
}

func Level(level int) func(*heading) {
	return func(h *heading) {
		h.Level = level
	}
}

func (h *heading) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}
