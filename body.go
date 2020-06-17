package html

type body struct {
	Element
}

func Body(attrs ...func(HtmlElement)) HtmlElement {
	b := &body{}
	b.Element.tag = "body"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (b *body) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.children = append(b.children, element)
	}
	return b
}
