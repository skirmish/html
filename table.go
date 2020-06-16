package html

import "io"

type table struct {
	Element
	Children []HtmlElement
}

func Table(attrs ...func(HtmlElement)) HtmlElement {
	f := &table{}
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

func (f *table) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type tr struct {
	Element
	Children []HtmlElement
}

func Tr(attrs ...func(HtmlElement)) HtmlElement {
	f := &tr{}
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

func (f *tr) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type th struct {
	Element
	Children []HtmlElement
}

func Th(attrs ...func(HtmlElement)) HtmlElement {
	f := &th{}
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

func (f *th) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type td struct {
	Element
	Children []HtmlElement
}

func Td(attrs ...func(HtmlElement)) HtmlElement {
	f := &td{}
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

func (f *td) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type caption struct {
	Element
	Children []HtmlElement
}

func Caption(attrs ...func(HtmlElement)) HtmlElement {
	f := &caption{}
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

func (f *caption) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type colgroup struct {
	Element
	Children []HtmlElement
}

func Colgroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &colgroup{}
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

func (f *colgroup) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type col struct {
	Element
	Children []HtmlElement
}

func Col(attrs ...func(HtmlElement)) HtmlElement {
	f := &col{}
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

func (f *col) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type thead struct {
	Element
	Children []HtmlElement
}

func Thead(attrs ...func(HtmlElement)) HtmlElement {
	f := &thead{}
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

func (f *thead) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type tfoot struct {
	Element
	Children []HtmlElement
}

func Tfoot(attrs ...func(HtmlElement)) HtmlElement {
	f := &tfoot{}
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

func (f *tfoot) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type tbody struct {
	Element
	Children []HtmlElement
}

func Tbody(attrs ...func(HtmlElement)) HtmlElement {
	f := &tbody{}
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

func (f *tbody) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}
