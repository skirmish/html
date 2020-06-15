package html

import "io"

type b struct {
	Element
	Children []HtmlElement
}

func B(attrs ...func(HtmlElement)) HtmlElement {
	f := &b{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *b) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "b", a.Children)
}

func (a *b) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *b) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
