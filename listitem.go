package html

import "io"

type listitem struct {
	Element
	Children []HtmlElement
}

func Li(attrs ...func(HtmlElement)) HtmlElement {
	li := &listitem{}
	for _, attr := range attrs {
		attr(li)
	}
	return li
}
func (l *listitem) Render(w io.Writer) (int, error) {
	return l.Element.Render(w, "li", l.Children)
}

func (l *listitem) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		l.Children = append(l.Children, element)
	}
	return l
}

func (l *listitem) addAttribute(key string, val string) {
	l.Element.AddAttribute(key, val)
}
