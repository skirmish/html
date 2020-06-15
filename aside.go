package html

import "io"

type aside struct {
	Element
	Children []HtmlElement
}

func Aside(attrs ...func(HtmlElement)) HtmlElement {
	f := &aside{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *aside) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "aside", a.Children)
}

func (a *aside) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *aside) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
