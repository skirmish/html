package html

import "io"

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

func (a *image) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "img", a.Children)
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