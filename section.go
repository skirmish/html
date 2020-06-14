package html

type section struct {
	Element
	Children []HtmlElement
}

func Section(attrs ...func(HtmlElement)) HtmlElement {
	ul := &section{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}
func (u *section) Render() string {
	output := "<section"
	u.InitAttributes()
	output += u.Element.RenderAttr()
	output += ">"
	for _, child := range u.Children {
		output += child.Render()
	}
	output += "</section>"
	return output
}

func (u *section) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *section) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}
