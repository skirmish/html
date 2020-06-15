package html

import "io"

type blockquote struct {
	Element
	Children []HtmlElement
}

func Blockquote(attrs ...func(HtmlElement)) HtmlElement {
	f := &blockquote{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *blockquote) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "blockquote", a.Children)
}

func (a *blockquote) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *blockquote) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
