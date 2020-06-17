package html

import "io"

type unorderedlist struct {
	Element
	Children []HtmlElement
}

func Ul(attrs ...func(HtmlElement)) HtmlElement {
	ul := &unorderedlist{}
	ul.tag = "ul"
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

type orderedlist struct {
	Element
	Children []HtmlElement
}

func Ol(attrs ...func(HtmlElement)) HtmlElement {
	ul := &orderedlist{}
	ul.tag = "ol"
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

type listitem struct {
	Element
	Children []HtmlElement
}

func Li(attrs ...func(HtmlElement)) HtmlElement {
	li := &listitem{}
	li.tag = "li"
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

type dl struct {
	Element
	Children []HtmlElement
}

func Dl(attrs ...func(HtmlElement)) HtmlElement {
	f := &dl{}
	f.tag = "dl"
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

type dt struct {
	Element
	Children []HtmlElement
}

func Dt(attrs ...func(HtmlElement)) HtmlElement {
	f := &dt{}
	f.tag = "dt"
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

type dd struct {
	Element
	Children []HtmlElement
}

func Dd(attrs ...func(HtmlElement)) HtmlElement {
	f := &dd{}
	f.tag = "dd"
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
