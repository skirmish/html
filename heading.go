package html

import "fmt"

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

func (h *heading) Render() string {
	output := ""
	output += fmt.Sprintf("<h%d>", h.Level)
	for _, el := range h.Children {
		output += el.Render()
	}
	output += fmt.Sprintf("</h%d>", h.Level)
	return output
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
