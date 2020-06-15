package html

import "io"

type br struct {
	Element
	Children []HtmlElement
}

func Br(attrs ...func(HtmlElement)) HtmlElement {
	f := &br{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *br) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "br", a.Children)
}

func (a *br) AddElements(elements ...HtmlElement) HtmlElement {
	//	for _, element := range elements {
	//		a.Children = append(a.Children, element)
	//	}
	return a
}

func (a *br) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
