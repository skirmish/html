package html

import "io"

type docType struct {
	Element
}

func DocType() HtmlElement {
	return &docType{}
}

func (d *docType) Render(w io.Writer) (int, error) {
	return w.Write([]byte("<!DOCTYPE html>\n"))
}

func (d *docType) AddElements(elements ...HtmlElement) HtmlElement {
	return d
}

type html struct {
	Element
}

func Html(attrs ...func(HtmlElement)) HtmlElement {
	h := &html{}
	h.tag = "html"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func (h *html) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}

type head struct {
	Element
}

func Head(attrs ...func(HtmlElement)) HtmlElement {
	h := &head{}
	h.tag = "head"
	for _, attr := range attrs {
		attr(h)
	}
	return h
}

func (h *head) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		h.children = append(h.children, element)
	}
	return h
}

type meta struct {
	Element
}

func Meta(attrs ...func(HtmlElement)) HtmlElement {
	m := &meta{}
	m.tag = "meta"
	for _, attr := range attrs {
		attr(m)
	}
	return m
}

func (m *meta) AddElements(elements ...HtmlElement) HtmlElement {
	return m
}

type title struct {
	Element
}

func Title(attrs ...func(HtmlElement)) HtmlElement {
	t := &title{}
	t.tag = "title"
	for _, attr := range attrs {
		attr(t)
	}
	return t
}

func (t *title) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		t.children = append(t.children, element)
	}
	return t
}

type body struct {
	Element
}

func Body(attrs ...func(HtmlElement)) HtmlElement {
	b := &body{}
	b.Element.tag = "body"
	for _, attr := range attrs {
		attr(b)
	}
	return b
}

func (b *body) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		b.children = append(b.children, element)
	}
	return b
}

type content struct {
	Element
	Content string
}

// Content adds raw content to the stream.
func Content(data string) HtmlElement {
	c := &content{Content: data}
	c.Element.tag = "body"
	return c
}

func (c *content) Render(w io.Writer) (int, error) {
	return w.Write([]byte(c.Content))
}

func (c *content) AddElements(elements ...HtmlElement) HtmlElement {
	return c
}
