package html

func Html(attrs ...func(HtmlElement)) HtmlElement {
	h := &html{}
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

type html struct {
	Element
	Children []HtmlElement
}

func (h *html) Render() string {
	output := "<html"
	h.InitAttributes()
	output += h.Element.RenderAttr()
	output += ">\n"
	for _, child := range h.Children {
		output += child.Render()
	}
	output += "</html>\n"
	return output
}

func (h *html) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}

func (h *html) addAttribute(key string, val string) {
	h.Element.AddAttribute(key, val)
}
