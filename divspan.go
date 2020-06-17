package html

import "io"

type div struct {
	Element
	Children []HtmlElement
}

func Div(attrs ...func(HtmlElement)) HtmlElement {
	b := &div{}
	b.tag = "div"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (d *div) Render(w io.Writer) (int, error) {
	return d.Element.Render(w, "div", d.Children)
}

func (d *div) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		d.Children = append(d.Children, element)
	}
	return d
}

type span struct {
	Element
	Children []HtmlElement
}

func Span(attrs ...func(HtmlElement)) HtmlElement {
	b := &span{}
	b.tag = "span"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (s *span) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "span", s.Children)
}

func (s *span) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}
