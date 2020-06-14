package html

type title struct {
	Element
	Children []HtmlElement
}

func Title(attrs ...func(HtmlElement)) HtmlElement {
	t := &title{}
	for _, attr := range attrs {
		attr(t)
	}
	return t
}

func (t *title) Render() string {
	output := "<title>"
	for _, el := range t.Children {
		output += el.Render()
	}
	output += "</title>"
	return output
}

func (t *title) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		t.Children = append(t.Children, element)
	}
	return t
}

func (t *title) addAttribute(key string, val string) {
	t.Element.AddAttribute(key, val)
}
