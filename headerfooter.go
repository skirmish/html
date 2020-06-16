package html

import "io"

type header struct {
	Element
	Children []HtmlElement
}

func Header(attrs ...func(HtmlElement)) HtmlElement {
	b := &header{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (d *header) Render(w io.Writer) (int, error) {
	return d.Element.Render(w, "header", d.Children)
}

func (d *header) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		d.Children = append(d.Children, element)
	}
	return d
}

func (d *header) addAttribute(key string, val string) {
	d.Element.AddAttribute(key, val)
}

type footer struct {
	Element
	Children []HtmlElement
}

func Footer(attrs ...func(HtmlElement)) HtmlElement {
	b := &footer{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (s *footer) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "footer", s.Children)
}

func (s *footer) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

func (s *footer) addAttribute(key string, val string) {
	s.Element.AddAttribute(key, val)
}
