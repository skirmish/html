package html

import "io"

func Html(attrs ...func(HtmlElement)) HtmlElement {
	h := &html{}
	h.tag = "html"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

type html struct {
	Element
	Children []HtmlElement
}

func (a *html) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "html", a.Children)
}

func (h *html) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}
