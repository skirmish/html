package html

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

func (s *script) Render() string {
	output := "<script>"
	for _, child := range s.Children {
		output += child.Render()
	}
	output += "</script>\n"
	return output
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
