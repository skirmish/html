package html

import "io"

type section struct {
	Element
	Children []HtmlElement
}

func Section(attrs ...func(HtmlElement)) HtmlElement {
	ul := &section{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (a *section) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "section", a.Children)
}

func (u *section) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *section) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}
