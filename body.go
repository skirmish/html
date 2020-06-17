package html

import "io"

type body struct {
	Element
}

func Body(attrs ...func(HtmlElement)) HtmlElement {
	b := &body{}
	b.Element.tag = "body"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (a *body) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "body", a.Element.children)
}

func (b *body) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Element.children = append(b.Element.children, element)
	}
	return b
}
