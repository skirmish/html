package html

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

func (m *meta) AddElements(elements ...HtmlElement) HtmlElement {
	return m
}
