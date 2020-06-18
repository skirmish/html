package v2

import (
	"fmt"
	"io"
)

type Element interface {
	Render(w io.Writer) (n int, err error)
	AddElements(elements ...Element) Element
	AddAttribute(key string, val string)
}

type keyVal struct {
	Key   string
	Value string
}

type tag struct {
	tag        string    // the tag name
	empty      bool      // Is it an empty tag e.g. <br>
	level      int       // h1-h6
	attributes []keyVal  // Attributes
	children   []Element // Child elements
}

func (e *tag) addAttrs(attrs ...func(Element)) Element {
	for _, attr := range attrs {
		attr(e)
	}
	return e
}

func (e *tag) AddAttribute(key string, val string) {
	e.attributes = append(e.attributes, keyVal{
		Key:   key,
		Value: val,
	})
}

func (e *tag) AddElements(elements ...Element) Element {
	for _, element := range elements {
		e.children = append(e.children, element)
	}
	return e
}

func (e *tag) renderAttr(w io.Writer) (int, error) {
	if len(e.attributes) == 0 {
		return 0, nil
	}
	n, err := w.Write([]byte(" "))
	if err != nil {
		return n, err
	}
	for index, attr := range e.attributes {
		o, err := w.Write([]byte(attr.Key))
		if err != nil {
			return n, err
		}
		n += o
		o, err = w.Write([]byte("=\""))
		if err != nil {
			return n, err
		}
		n += o
		o, err = w.Write([]byte(attr.Value))
		if err != nil {
			return n, err
		}
		n += o
		o, err = w.Write([]byte("\""))
		if err != nil {
			return n, err
		}
		n += o
		if index+1 < len(e.attributes) {
			o, err = w.Write([]byte(" "))
			if err != nil {
				return n, err
			}
			n += o
		}
	}
	return n, err
}

func (e *tag) Render(w io.Writer) (int, error) {
	return e.renderElement(w, e.tag, e.children)
}

func (e *tag) renderElement(w io.Writer, tag string, children []Element) (int, error) {
	n, err := w.Write([]byte("<"))
	if err != nil {
		return n, err
	}
	o, err := w.Write([]byte(tag))
	if err != nil {
		return n, err
	}
	n += o
	o, err = e.renderAttr(w) // Attributes
	if err != nil {
		return n, err
	}
	n += o
	if children == nil {
		if e.empty {
			o, err = w.Write([]byte(">"))
			if err != nil {
				return n, err
			}
			n += o
			return n, err
		} else {
			o, err = w.Write([]byte("/>"))
			if err != nil {
				return n, err
			}
			n += o
			return n, err
		}
	}
	o, err = w.Write([]byte(">"))
	if err != nil {
		return n, err
	}
	n += o

	for _, child := range children {
		o, err = child.Render(w)
		if err != nil {
			return n, err
		}
		n += o
	}

	// RenderElement end
	o, err = w.Write([]byte("</"))
	if err != nil {
		return n, err
	}
	n += o

	o, err = w.Write([]byte(tag))
	if err != nil {
		return n, err
	}
	n += o
	o, err = w.Write([]byte(">"))
	if err != nil {
		return n, err
	}
	n += o
	return n, err
}

func Html(attrs ...func(Element)) Element {
	return (&tag{
		tag: "html",
	}).addAttrs(attrs...)
}

func Head(attrs ...func(Element)) Element {
	return (&tag{
		tag: "head",
	}).addAttrs(attrs...)
}

func Meta(attrs ...func(Element)) Element {
	return (&tag{
		tag:   "meta",
		empty: true,
	}).addAttrs(attrs...)
}

// Charset - for meta
func Charset(charset string) func(Element) {
	return func(m Element) {
		m.AddAttribute("charset", charset)
	}
}

func Body(attrs ...func(Element)) Element {
	return (&tag{
		tag: "body",
	}).addAttrs(attrs...)
}

func H(attrs ...func(*tag)) Element {
	t := &tag{}
	for _, attr := range attrs {
		attr(t)
	}
	t.tag = fmt.Sprintf("h%d", t.level)
	return t
}

func Level(level int) func(*tag) {
	return func(h *tag) {
		h.level = level
	}
}

type rawData struct {
	content string
}

func (e *rawData) AddElements(elements ...Element) Element {
	return e
}

func (e *rawData) AddAttribute(key string, val string) {
}

func (e *rawData) Render(w io.Writer) (int, error) {
	return w.Write([]byte(e.content))
}

func Content(data string) Element {
	return &rawData{
		content: data,
	}
}

func DocType() Element {
	return &rawData{
		content: "<!DOCTYPE html>\n",
	}
}
