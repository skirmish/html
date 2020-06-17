package html

type unorderedlist struct {
	Element
}

func Ul(attrs ...func(HtmlElement)) HtmlElement {
	ul := &unorderedlist{}
	ul.tag = "ul"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (u *unorderedlist) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.children = append(u.children, element)
	}
	return u
}

type orderedlist struct {
	Element
}

func Ol(attrs ...func(HtmlElement)) HtmlElement {
	ul := &orderedlist{}
	ul.tag = "ol"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (u *orderedlist) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.children = append(u.children, element)
	}
	return u
}

type listitem struct {
	Element
}

func Li(attrs ...func(HtmlElement)) HtmlElement {
	li := &listitem{}
	li.tag = "li"
	for _, attr := range attrs {
		attr(li)
	}
	return li
}

func (l *listitem) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		l.children = append(l.children, element)
	}
	return l
}

type dl struct {
	Element
}

func Dl(attrs ...func(HtmlElement)) HtmlElement {
	f := &dl{}
	f.tag = "dl"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dl) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}

type dt struct {
	Element
}

func Dt(attrs ...func(HtmlElement)) HtmlElement {
	f := &dt{}
	f.tag = "dt"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dt) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}

type dd struct {
	Element
}

func Dd(attrs ...func(HtmlElement)) HtmlElement {
	f := &dd{}
	f.tag = "dd"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *dd) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}
