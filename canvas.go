package html

import "io"

type canvas struct {
	Element
	Children []HtmlElement
}

func Canvas(attrs ...func(HtmlElement)) HtmlElement {
	i := &canvas{}
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *canvas) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "canvas", a.Children)
}

func (i *canvas) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	i.Children = append(i.Children,element)
	//}
	return i
}

func (i *canvas) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}
