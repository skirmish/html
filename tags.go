package html

import "io"

// base, bdi, bdo,
// cite, code, data, datalist, del, ins, dfn, dialog,
// em, hr, i, kbd, link, noscript, object, param, pre, progress,
// q, rp, rt, ruby, s, samp, small, strong, sub, sup, template, track, tt, u, var, wbr

type abbr struct {
	Element
	Children []HtmlElement
}

func Abbr(attrs ...func(HtmlElement)) HtmlElement {
	f := &abbr{}
	f.tag = "abbr"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *abbr) Render(w io.Writer) (int, error) {
	return b.Element.Render(w, "abbr", b.Children)
}

func (b *abbr) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Children = append(b.Children, element)
	}
	return b
}

type address struct {
	Element
	Children []HtmlElement
}

func Address(attrs ...func(HtmlElement)) HtmlElement {
	f := &address{}
	f.tag = "address"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *address) Render(w io.Writer) (int, error) {
	return b.Element.Render(w, "address", b.Children)
}

func (b *address) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Children = append(b.Children, element)
	}
	return b
}

type b struct {
	Element
	Children []HtmlElement
}

func B(attrs ...func(HtmlElement)) HtmlElement {
	f := &b{}
	f.tag = "b"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (b *b) Render(w io.Writer) (int, error) {
	return b.Element.Render(w, "b", b.Children)
}

func (b *b) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.Children = append(b.Children, element)
	}
	return b
}

type blockquote struct {
	Element
	Children []HtmlElement
}

func Blockquote(attrs ...func(HtmlElement)) HtmlElement {
	f := &blockquote{}
	f.tag = "blockquote"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *blockquote) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "blockquote", a.Children)
}

func (a *blockquote) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		a.Children = append(a.Children, element)
	}
	return a
}

type br struct {
	Element
	Children []HtmlElement
}

func Br(attrs ...func(HtmlElement)) HtmlElement {
	f := &br{}
	f.tag = "br"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (a *br) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "br", a.Children)
}

func (a *br) AddElements(elements ...HtmlElement) HtmlElement {
	//	for _, element := range elements {
	//		a.Children = append(a.Children, element)
	//	}
	return a
}

type meter struct {
	Element
	Children []HtmlElement
}

func Meter(attrs ...func(HtmlElement)) HtmlElement {
	ul := &meter{}
	ul.tag = "meter"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func (a *meter) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "meter", a.Children)
}

func (u *meter) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		u.Children = append(u.Children, element)
	}
	return u
}

type script struct {
	Element
	Children []HtmlElement
}

func Script(attrs ...func(HtmlElement)) HtmlElement {
	s := &script{}
	s.tag = "script"
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (a *script) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "script", a.Children)
}

func (s *script) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type style struct {
	Element
	Children []HtmlElement
}

func Style(attrs ...func(HtmlElement)) HtmlElement {
	s := &style{}
	s.tag = "style"
	for _, attr := range attrs {
		attr(s)
	}
	return s
}

func (a *style) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "style", a.Children)
}

func (s *style) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		s.Children = append(s.Children, element)
	}
	return s
}

type paragraph struct {
	Element
	Children []HtmlElement
}

func P(attrs ...func(HtmlElement)) HtmlElement {
	p := &paragraph{}
	p.tag = "p"
	for _, attr := range attrs {
		attr(p)
	}
	return p
}

func (a *paragraph) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "p", a.Children)
}

func (p *paragraph) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		p.Children = append(p.Children, element)
	}
	return p
}
