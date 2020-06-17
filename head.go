package html

import "io"

type head struct {
	Element
	Children []HtmlElement
}

func Head(attrs ...func(HtmlElement)) HtmlElement {
	h := &head{}
	h.tag = "head"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func (a *head) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "head", a.Children)
}

func (h *head) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}
