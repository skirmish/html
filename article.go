package html

import "io"

type article struct {
	Element
	Children []HtmlElement
}

func Article(attrs ...func(HtmlElement)) HtmlElement {
	f := &article{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *article) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "article", a.Children)
}

func (a *article) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *article) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
