package html

import (
	"bytes"
	"io"
)

type HtmlElement interface {
	Render(w io.Writer) (n int, err error)
	AddElements(elements ...HtmlElement) HtmlElement
	AddAttribute(key string, val string)
}

// sizer is an interface that returns the size of what the rendered element would be.
type sizer interface {
	getRenderSize() int // If it is a static thing, then we can get a size.
}

type variabler interface {
	isDynamic() bool // If the thing is dynamic
}

type KeyVal struct {
	Key   string
	Value string
}

func (kv *KeyVal) getRenderSize() int {
	size := 3 // account for =""
	size += len(kv.Key)
	size += len(kv.Value)
	// key="value"
	return size
}

type Element struct {
	tag        string
	attributes []KeyVal
	childData  []byte
	children   []HtmlElement
}

func (e *Element) InitAttributes() {
	//e.attributes = make([]KeyVal,0)
}

func (e *Element) AddAttribute(key string, val string) {
	e.attributes = append(e.attributes, KeyVal{
		Key:   key,
		Value: val,
	})
}

func (e Element) RenderAttr(w io.Writer) (int, error) {
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

func (a *Element) Render(w io.Writer, tag string, children []HtmlElement) (int, error) {
	n, err := w.Write([]byte("<"))
	if err != nil {
		return n, err
	}
	o, err := w.Write([]byte(tag))
	if err != nil {
		return n, err
	}
	n += o
	o, err = a.RenderAttr(w) // Attributes
	if err != nil {
		return n, err
	}
	n += o
	if children == nil {
		o, err = w.Write([]byte("/>"))
		if err != nil {
			return n, err
		}
		n += o
		return n, err
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

	// Render end
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

// FastRender - Do NOT use
func (a *Element) FastRender(w io.Writer, tag string, children []HtmlElement) (int, error) {
	if len(a.childData) == 0 {
		// Record the rendering into childata.
		buf := new(bytes.Buffer)
		_, err := a.Render(buf, tag, children)
		if err != nil {
			return 0, err
		}
		a.childData = buf.Bytes()
	}
	n, err := w.Write(a.childData)
	return n, err
}

func (e *Element) getRenderSize() int {
	size := 1 // account for :space:
	size += len(e.attributes) - 1
	for _, el := range e.attributes {
		size += el.getRenderSize()
	}
	return size
}
