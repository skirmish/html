package html

import (
	"fmt"
	"io"
)

type heading struct {
	Element
	Level    int
	Children []HtmlElement
}

func Heading(attrs ...func(*heading)) HtmlElement {
	h := &heading{}
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func Level(level int) func(*heading) {
	return func(h *heading) {
		h.Level = level
	}
}

func (a *heading) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, fmt.Sprintf("h%d", a.Level), a.Children)
}

func (h *heading) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}

func (h *heading) addAttribute(key string, val string) {
	h.Element.AddAttribute(key, val)
}
