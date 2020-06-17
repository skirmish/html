package html

type head struct {
	Element
}

func Head(attrs ...func(HtmlElement)) HtmlElement {
	h := &head{}
	h.tag = "head"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func (h *head) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}
