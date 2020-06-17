package html

import "io"

type header struct {
	Element
	Children []HtmlElement
}

func Header(attrs ...func(HtmlElement)) HtmlElement {
	b := &header{}
	b.tag = "header"
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

type footer struct {
	Element
	Children []HtmlElement
}

func Footer(attrs ...func(HtmlElement)) HtmlElement {
	b := &footer{}
	b.tag = "footer"
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

type article struct {
	Element
	Children []HtmlElement
}

func Article(attrs ...func(HtmlElement)) HtmlElement {
	f := &article{}
	f.tag = "article"
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

type aside struct {
	Element
	Children []HtmlElement
}

func Aside(attrs ...func(HtmlElement)) HtmlElement {
	f := &aside{}
	f.tag = "aside"
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

type section struct {
	Element
	Children []HtmlElement
}

func Section(attrs ...func(HtmlElement)) HtmlElement {
	ul := &section{}
	ul.tag = "section"
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

type nav struct {
	Element
	Children []HtmlElement
}

func Nav(attrs ...func(HtmlElement)) HtmlElement {
	ul := &nav{}
	ul.tag = "nav"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *nav) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "nav", s.Children)
}

func (s *nav) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type details struct {
	Element
	Children []HtmlElement
}

func Details(attrs ...func(HtmlElement)) HtmlElement {
	ul := &details{}
	ul.tag = "details"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *details) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "details", s.Children)
}

func (s *details) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type main struct {
	Element
	Children []HtmlElement
}

func Main(attrs ...func(HtmlElement)) HtmlElement {
	ul := &main{}
	ul.tag = "main"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *main) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "main", s.Children)
}

func (s *main) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type mark struct {
	Element
	Children []HtmlElement
}

func Mark(attrs ...func(HtmlElement)) HtmlElement {
	ul := &mark{}
	ul.tag = "mark"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *mark) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "mark", s.Children)
}

func (s *mark) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type summary struct {
	Element
	Children []HtmlElement
}

func Summary(attrs ...func(HtmlElement)) HtmlElement {
	ul := &summary{}
	ul.tag = "summary"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *summary) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "summary", s.Children)
}

func (s *summary) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type time struct {
	Element
	Children []HtmlElement
}

func Time(attrs ...func(HtmlElement)) HtmlElement {
	ul := &time{}
	ul.tag = "time"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *time) Render(w io.Writer) (int, error) {
	return s.Element.Render(w, "time", s.Children)
}

func (s *time) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}
