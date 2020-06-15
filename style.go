package html

import "io"

type style struct {
	Element
	Children []HtmlElement
}

func Style(attrs ...func(HtmlElement)) HtmlElement {
	s := &style{}
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (a *style) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "style", a.Children)
}

func (s *style) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

func (s *style) addAttribute(key string, val string) {
	s.Element.AddAttribute(key, val)
}
