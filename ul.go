package html

type unorderedlist struct {
	Element
	Children []HtmlElement
}

func Ul(attrs ...func(HtmlElement)) HtmlElement {
	ul := &unorderedlist{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}
func (u *unorderedlist) Render() string {
	output := "<ul"
	u.InitAttributes()
	output += u.Element.RenderAttr()
	output += ">"
	for _, child := range u.Children {
		output += child.Render()
	}
	output += "</ul>"
	return output
}

func (u *unorderedlist) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *unorderedlist) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}
