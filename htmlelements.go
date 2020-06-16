package html

import "io"

type HtmlElement interface {
	Render(w io.Writer) (n int, err error)
	AddElements(elements ...HtmlElement) HtmlElement
	addAttribute(key string, val string)
}

// sizer is an interface that returns the size of what the rendered element would be.
type sizer interface {
	GetRenderSize() int
}

type variabler interface {
	IsDynamic() bool
}

type KeyVal struct {
	Key   string
	Value string
}

func (kv *KeyVal) GetRenderSize() int {
	size := 3 // account for =""
	size += len(kv.Key)
	size += len(kv.Value)
	// key="value"
	return size
}

type Element struct {
	attributes []KeyVal
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

func (e *Element) GetRenderSize() int {
	size := 1 // account for :space:
	size += len(e.attributes) - 1
	for _, el := range e.attributes {
		size += el.GetRenderSize()
	}
	return size
}
