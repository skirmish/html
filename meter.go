package html

import "io"

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

func (a *meter) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "meter", a.Children)
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
