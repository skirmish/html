package html

import "io"

type body struct {
	Element
	Children []HtmlElement
}

func Body(attrs ...func(HtmlElement)) HtmlElement {
	b := &body{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (a *body) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "body", a.Children)
}

func (b *body) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Children = append(b.Children, element)
	}
	return b
}

func (b *body) addAttribute(key string, val string) {
	b.Element.AddAttribute(key, val)
}
