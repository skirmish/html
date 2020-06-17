package html

type title struct {
	Element
}

func Title(attrs ...func(HtmlElement)) HtmlElement {
	t := &title{}
	t.tag = "title"
	for _, attr := range attrs {
		attr(t)
	}
	return t
}

func (t *title) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		t.children = append(t.children, element)
	}
	return t
}
