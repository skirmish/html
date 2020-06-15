package html

import "io"

type unorderedlist struct {
	Element
	Children []HtmlElement
}

func Ul(attrs ...func(HtmlElement)) HtmlElement {
	ul := &unorderedlist{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (u *unorderedlist) Render(w io.Writer) (int, error) {
	return u.Element.Render(w, "ul", u.Children)
}

func (u *unorderedlist) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *unorderedlist) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}
