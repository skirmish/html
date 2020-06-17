package html

import "io"

type table struct {
	Element
	Children []HtmlElement
}

func Table(attrs ...func(HtmlElement)) HtmlElement {
	f := &table{}
	f.tag = "table"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *table) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "table", f.Children)
}

func (f *table) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type tr struct {
	Element
	Children []HtmlElement
}

func Tr(attrs ...func(HtmlElement)) HtmlElement {
	f := &tr{}
	f.tag = "tr"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tr) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "tr", f.Children)
}

func (f *tr) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type th struct {
	Element
	Children []HtmlElement
}

func Th(attrs ...func(HtmlElement)) HtmlElement {
	f := &th{}
	f.tag = "th"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *th) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "th", f.Children)
}

func (f *th) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type td struct {
	Element
	Children []HtmlElement
}

func Td(attrs ...func(HtmlElement)) HtmlElement {
	f := &td{}
	f.tag = "td"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *td) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "td", f.Children)
}

func (f *td) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type caption struct {
	Element
	Children []HtmlElement
}

func Caption(attrs ...func(HtmlElement)) HtmlElement {
	f := &caption{}
	f.tag = "caption"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *caption) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "caption", f.Children)
}

func (f *caption) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type colgroup struct {
	Element
	Children []HtmlElement
}

func Colgroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &colgroup{}
	f.tag = "colgroup"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *colgroup) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "colgroup", f.Children)
}

func (f *colgroup) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type col struct {
	Element
	Children []HtmlElement
}

func Col(attrs ...func(HtmlElement)) HtmlElement {
	f := &col{}
	f.tag = "col"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *col) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "col", f.Children)
}

func (f *col) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type thead struct {
	Element
	Children []HtmlElement
}

func Thead(attrs ...func(HtmlElement)) HtmlElement {
	f := &thead{}
	f.tag = "thead"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *thead) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "thead", f.Children)
}

func (f *thead) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type tfoot struct {
	Element
	Children []HtmlElement
}

func Tfoot(attrs ...func(HtmlElement)) HtmlElement {
	f := &tfoot{}
	f.tag = "tfoot"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tfoot) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "tfoot", f.Children)
}

func (f *tfoot) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

type tbody struct {
	Element
	Children []HtmlElement
}

func Tbody(attrs ...func(HtmlElement)) HtmlElement {
	f := &tbody{}
	f.tag = "tbody"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *tbody) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "tbody", f.Children)
}

func (f *tbody) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}
