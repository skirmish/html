package html

func Html(attrs ...func(HtmlElement)) HtmlElement {
	h := &html{}
	h.tag = "html"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

type html struct {
	Element
}

func (h *html) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}
