package html

type head struct {
	Element
	Children []HtmlElement
}

func Head(attrs ...func(HtmlElement)) HtmlElement {
	h := &head{}
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func (h *head) Render() string {
	output := "<head>"
	for _, child := range h.Children {
		output += child.Render()
	}
	output += "</head>\n"
	return output
}

func (h *head) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}

func (h *head) addAttribute(key string, val string) {
	h.Element.AddAttribute(key, val)
}
