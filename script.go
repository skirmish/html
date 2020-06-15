package html

import "io"

type script struct {
	Element
	Children []HtmlElement
}

func Script(attrs ...func(HtmlElement)) HtmlElement {
	s := &script{}
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (a *script) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "script", a.Children)
}

func (s *script) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

func (s *script) addAttribute(key string, val string) {
	s.Element.AddAttribute(key, val)
}
