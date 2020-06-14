package html

type listitem struct {
	Element
	Children []HtmlElement
}

func Li(attrs ...func(HtmlElement)) HtmlElement {
	li := &listitem{}
	for _, attr := range attrs {
		attr(li)
	}
	return li
}
func (l *listitem) Render() string {
	output := "<li"
	l.InitAttributes()
	output += l.Element.RenderAttr()
	output += ">"
	for _, child := range l.Children {
		output += child.Render()
	}
	output += "</li>"
	return output
}

func (l *listitem) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		l.Children = append(l.Children, element)
	}
	return l
}

func (l *listitem) addAttribute(key string, val string) {
	l.Element.AddAttribute(key, val)
}
