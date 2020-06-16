package html

import "io"

type header struct {
	Element
	Children []HtmlElement
}

func Header(attrs ...func(HtmlElement)) HtmlElement {
	b := &header{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (h *header) Render(w io.Writer) (int, error) {
	return h.Element.Render(w, "header", h.Children)
}

func (h *header) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.Children = append(h.Children, element)
	}
	return h
}

func (h *header) addAttribute(key string, val string) {
	h.Element.AddAttribute(key, val)
}

type footer struct {
	Element
	Children []HtmlElement
}

func Footer(attrs ...func(HtmlElement)) HtmlElement {
	b := &footer{}
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (f *footer) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "footer", f.Children)
}

func (f *footer) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *footer) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

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

type aside struct {
	Element
	Children []HtmlElement
}

func Aside(attrs ...func(HtmlElement)) HtmlElement {
	f := &aside{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *aside) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "aside", a.Children)
}

func (a *aside) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

func (a *aside) addAttribute(key string, val string) {
	a.Element.AddAttribute(key, val)
}

type section struct {
	Element
	Children []HtmlElement
}

func Section(attrs ...func(HtmlElement)) HtmlElement {
	ul := &section{}
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *section) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "section", s.Children)
}

func (s *section) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

func (s *section) addAttribute(key string, val string) {
	s.Element.AddAttribute(key, val)
}
