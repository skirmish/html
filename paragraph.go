package html

import "io"

type paragraph struct {
	Element
	Children []HtmlElement
}

func P(attrs ...func(HtmlElement)) HtmlElement {
	p := &paragraph{}
	for _, attr := range attrs {
		attr(p)
	}
	return p
}

func (a *paragraph) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "p", a.Children)
}

func (p *paragraph) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		p.Children = append(p.Children, element)
	}
	return p
}

func (p *paragraph) addAttribute(key string, val string) {
	p.Element.AddAttribute(key, val)
}
