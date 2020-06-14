package html

type image struct {
	Element
	Children []HtmlElement
}

func Img(attrs ...func(HtmlElement)) HtmlElement {
	i := &image{}
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *image) Render() string {
	output := "<img"
	output += i.Element.RenderAttr()
	//for _,child := range l.Children {
	//	output += child.Render()
	//}
	output += "/>"
	return output
}

func (i *image) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	i.Children = append(i.Children,element)
	//}
	return i
}

func (i *image) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}
