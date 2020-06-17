package html

type div struct {
	Element
}

func Div(attrs ...func(HtmlElement)) HtmlElement {
	b := &div{}
	b.tag = "div"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (d *div) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		d.children = append(d.children, element)
	}
	return d
}

type span struct {
	Element
}

func Span(attrs ...func(HtmlElement)) HtmlElement {
	b := &span{}
	b.tag = "span"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (s *span) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}
