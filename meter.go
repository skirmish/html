package html

type meter struct {
	Element
	Children []HtmlElement
}

func Meter(attrs ...func(HtmlElement)) HtmlElement {
	ul := &meter{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}
func (u *meter) Render() string {
	output := "<meter"
	u.InitAttributes()
	output += u.Element.RenderAttr()
	output += ">"
	for _, child := range u.Children {
		output += child.Render()
	}
	output += "</meter>"
	return output
}

func (u *meter) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *meter) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}
