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

func (d *docType) addAttribute(key string, val string) {
	d.Element.AddAttribute(key, val)
}
