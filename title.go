package html

import "io"

type title struct {
	Element
	Children []HtmlElement
}

func Title(attrs ...func(HtmlElement)) HtmlElement {
	t := &title{}
	t.tag = "title"
	for _, attr := range attrs {
		attr(t)
	}
	return t
}

func (a *title) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "title", a.Children)
}

func (t *title) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		t.Children = append(t.Children, element)
	}
	return t
}
