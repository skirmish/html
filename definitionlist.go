package html

import "io"

type dl struct {
	Element
	Children []HtmlElement
}

func Dl(attrs ...func(HtmlElement)) HtmlElement {
	f := &dl{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dl) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "dl", a.Children)
}

func (a *dl) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *dl) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}

type dt struct {
	Element
	Children []HtmlElement
}

func Dt(attrs ...func(HtmlElement)) HtmlElement {
	f := &dt{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dt) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "dt", a.Children)
}

func (a *dt) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *dt) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}

type dd struct {
	Element
	Children []HtmlElement
}

func Dd(attrs ...func(HtmlElement)) HtmlElement {
	f := &dd{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dd) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "dd", a.Children)
}

func (a *dd) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *dd) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}
