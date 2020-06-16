package html

import "io"

type form struct {
	Element
	Children []HtmlElement
}

func Form(attrs ...func(HtmlElement)) HtmlElement {
	f := &form{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *form) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "form", f.Children)
}

func (f *form) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *form) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type fieldset struct {
	Element
	Children []HtmlElement
}

func Fieldset(attrs ...func(HtmlElement)) HtmlElement {
	f := &fieldset{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *fieldset) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "fieldset", f.Children)
}

func (f *fieldset) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *fieldset) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}
