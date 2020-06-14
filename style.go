package html

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

func (s *style) Render() string {
	output := "<style>"
	for _, child := range s.Children {
		output += child.Render()
	}
	output += "</style>"
	return output
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
