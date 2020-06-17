package html

type header struct {
	Element
}

func Header(attrs ...func(HtmlElement)) HtmlElement {
	b := &header{}
	b.tag = "header"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (h *header) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}

type footer struct {
	Element
}

func Footer(attrs ...func(HtmlElement)) HtmlElement {
	b := &footer{}
	b.tag = "footer"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (f *footer) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type article struct {
	Element
}

func Article(attrs ...func(HtmlElement)) HtmlElement {
	f := &article{}
	f.tag = "article"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *article) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}

type aside struct {
	Element
}

func Aside(attrs ...func(HtmlElement)) HtmlElement {
	f := &aside{}
	f.tag = "aside"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *aside) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}

type section struct {
	Element
}

func Section(attrs ...func(HtmlElement)) HtmlElement {
	ul := &section{}
	ul.tag = "section"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *section) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type nav struct {
	Element
}

func Nav(attrs ...func(HtmlElement)) HtmlElement {
	ul := &nav{}
	ul.tag = "nav"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *nav) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type details struct {
	Element
}

func Details(attrs ...func(HtmlElement)) HtmlElement {
	ul := &details{}
	ul.tag = "details"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *details) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type main struct {
	Element
}

func Main(attrs ...func(HtmlElement)) HtmlElement {
	ul := &main{}
	ul.tag = "main"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *main) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type mark struct {
	Element
}

func Mark(attrs ...func(HtmlElement)) HtmlElement {
	ul := &mark{}
	ul.tag = "mark"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *mark) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type summary struct {
	Element
}

func Summary(attrs ...func(HtmlElement)) HtmlElement {
	ul := &summary{}
	ul.tag = "summary"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *summary) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type time struct {
	Element
}

func Time(attrs ...func(HtmlElement)) HtmlElement {
	ul := &time{}
	ul.tag = "time"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (s *time) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}
