package html

// base, bdi, bdo,
// cite, code, data, datalist, del, ins, dfn, dialog,
// em, hr, i, kbd, link, noscript, object, param, pre, progress,
// q, rp, rt, ruby, s, samp, small, strong, sub, sup, template, track, tt, u, var, wbr

type abbr struct {
	Element
}

func Abbr(attrs ...func(HtmlElement)) HtmlElement {
	f := &abbr{}
	f.tag = "abbr"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *abbr) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.children = append(b.children, element)
	}
	return b
}

type address struct {
	Element
}

func Address(attrs ...func(HtmlElement)) HtmlElement {
	f := &address{}
	f.tag = "address"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *address) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.children = append(b.children, element)
	}
	return b
}

type b struct {
	Element
}

func B(attrs ...func(HtmlElement)) HtmlElement {
	f := &b{}
	f.tag = "b"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *b) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.children = append(b.children, element)
	}
	return b
}

type blockquote struct {
	Element
}

func Blockquote(attrs ...func(HtmlElement)) HtmlElement {
	f := &blockquote{}
	f.tag = "blockquote"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *blockquote) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.children = append(a.children, element)
	}
	return a
}

type br struct {
	Element
}

func Br(attrs ...func(HtmlElement)) HtmlElement {
	f := &br{}
	f.tag = "br"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *br) AddElements(elements ...HtmlElement) HtmlElement {
	//	for _, element := range elements {
	//		a.children = append(a.children, element)
	//	}
	return a
}

type meter struct {
	Element
}

func Meter(attrs ...func(HtmlElement)) HtmlElement {
	ul := &meter{}
	ul.tag = "meter"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (u *meter) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.children = append(u.children, element)
	}
	return u
}

type script struct {
	Element
}

func Script(attrs ...func(HtmlElement)) HtmlElement {
	s := &script{}
	s.tag = "script"
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (s *script) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type style struct {
	Element
}

func Style(attrs ...func(HtmlElement)) HtmlElement {
	s := &style{}
	s.tag = "style"
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (s *style) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.children = append(s.children, element)
	}
	return s
}

type paragraph struct {
	Element
}

func P(attrs ...func(HtmlElement)) HtmlElement {
	p := &paragraph{}
	p.tag = "p"
	for _, attr := range attrs {
		attr(p)
	}
	return p
}

func (p *paragraph) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		p.children = append(p.children, element)
	}
	return p
}
