package html

type body struct {
	Element
	Children []HtmlElement
}

func Body(attrs ...func(HtmlElement)) HtmlElement {
	b := &body{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (b *body) Render() string {
	output := "<body>"
	for _, el := range b.Children {
		output += el.Render()
	}
	output += "</body>"
	return output
}

func (b *body) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Children = append(b.Children, element)
	}
	return b
}

func (b *body) addAttribute(key string, val string) {
	b.Element.AddAttribute(key, val)
}
