package html

import (
	"bytes"
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

func (kv *keyVal) getRenderSize() int {
	size := 3 + len(kv.Key) + len(kv.Value) // account for =""
	return size
}

type tag struct {
	tag        string    // the tag name
	empty      bool      // Is it an empty tag e.g. <br>, <wbr>, <media>, <source>
	level      int       // h1-h6
	attributes []keyVal  // Attributes
	children   []Element // Child elements
	childData  []byte    // cached children renderings
}

func (e *tag) addAttrs(attrs ...func(Element)) Element {
	for _, attr := range attrs {
		attr(e)
	}
	return e
}

// AddAttribute adds an attribute tot he node
func (e *tag) AddAttribute(key string, val string) {
	e.attributes = append(e.attributes, keyVal{Key: key, Value: val})
}

// AddElements adds child elements to this element
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

// below is for optimizing the render---snip

func (e *tag) fastRender(w io.Writer, tag string, children []Element) (int, error) {
	if len(e.childData) == 0 {
		// Record the rendering into child data.
		buf := new(bytes.Buffer)
		_, err := e.renderElement(buf, tag, children)
		if err != nil {
			return 0, err
		}
		e.childData = buf.Bytes()
	}
	n, err := w.Write(e.childData)
	return n, err
}

func (e *tag) getRenderSize() int {
	size := 1 // account for :space:
	size += len(e.attributes) - 1
	for _, el := range e.attributes {
		size += el.getRenderSize()
	}
	return size
}

//---endsnip---

// rawData type used for elements that are plain e.g. text.
type rawData struct {
	content string
}

// AddElements adds child elements to this element
func (e *rawData) AddElements(elements ...Element) Element { return e }

// AddAttribute adds an attribute tot he node
func (e *rawData) AddAttribute(key string, val string) {}

// Render outputs to an io.Writer the rendered version
func (e *rawData) Render(w io.Writer) (int, error) { return w.Write([]byte(e.content)) }

/*

// sizer is an interface that returns the size of what the rendered element would be.
type sizer interface {
	getRenderSize() int // If it is a static thing, then we can get a size.
}

type variabler interface {
	isDynamic() bool // If the thing is dynamic
}

func (e *Element) getRenderSize() int {
	size := 1 // account for :space:
	size += len(e.attributes) - 1
	for _, el := range e.attributes {
		size += el.getRenderSize()
	}
	return size
}
*/
