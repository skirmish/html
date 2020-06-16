package html

import "io"

type div struct {
	Element
	Children []HtmlElement
}

func Div(attrs ...func(HtmlElement)) HtmlElement {
	b := &div{}
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

func (d *div) addAttribute(key string, val string) {
	d.Element.AddAttribute(key, val)
}

type span struct {
	Element
	Children []HtmlElement
}

func Span(attrs ...func(HtmlElement)) HtmlElement {
	b := &span{}
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

func (s *span) addAttribute(key string, val string) {
	s.Element.AddAttribute(key, val)
}
