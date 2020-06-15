package html

import "io"

type form struct {
	Element
	Children []HtmlElement
}

func Form(attrs ...func(HtmlElement)) HtmlElement {
	f := &form{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *form) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "form", a.Children)
}

func (f *form) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *form) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}
