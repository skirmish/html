package html

import "io"

type meta struct {
	Element
}

func Meta(attrs ...func(HtmlElement)) HtmlElement {
	m := &meta{}
	m.tag = "meta"
	for _, attr := range attrs {
		attr(m)
	}
	return m
}

func (a *meta) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "meta", nil)
}

func (m *meta) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	p.Children = append(p.Children,element)
	//}
	return m
}
