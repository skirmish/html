package html

type table struct {
	Element
}

func Table(attrs ...func(HtmlElement)) HtmlElement {
	f := &table{}
	f.tag = "table"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *table) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type tr struct {
	Element
}

func Tr(attrs ...func(HtmlElement)) HtmlElement {
	f := &tr{}
	f.tag = "tr"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tr) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type th struct {
	Element
}

func Th(attrs ...func(HtmlElement)) HtmlElement {
	f := &th{}
	f.tag = "th"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *th) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type td struct {
	Element
}

func Td(attrs ...func(HtmlElement)) HtmlElement {
	f := &td{}
	f.tag = "td"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *td) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type caption struct {
	Element
}

func Caption(attrs ...func(HtmlElement)) HtmlElement {
	f := &caption{}
	f.tag = "caption"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *caption) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type colgroup struct {
	Element
}

func Colgroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &colgroup{}
	f.tag = "colgroup"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *colgroup) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type col struct {
	Element
}

func Col(attrs ...func(HtmlElement)) HtmlElement {
	f := &col{}
	f.tag = "col"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *col) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type thead struct {
	Element
}

func Thead(attrs ...func(HtmlElement)) HtmlElement {
	f := &thead{}
	f.tag = "thead"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *thead) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type tfoot struct {
	Element
}

func Tfoot(attrs ...func(HtmlElement)) HtmlElement {
	f := &tfoot{}
	f.tag = "tfoot"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tfoot) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type tbody struct {
	Element
}

func Tbody(attrs ...func(HtmlElement)) HtmlElement {
	f := &tbody{}
	f.tag = "tbody"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tbody) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}
