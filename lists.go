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

type orderedlist struct {
	Element
	Children []HtmlElement
}

func Ol(attrs ...func(HtmlElement)) HtmlElement {
	ul := &orderedlist{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (u *orderedlist) Render(w io.Writer) (int, error) {
	return u.Element.Render(w, "ol", u.Children)
}

func (u *orderedlist) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

func (u *orderedlist) addAttribute(key string, val string) {
	u.Element.AddAttribute(key, val)
}

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
