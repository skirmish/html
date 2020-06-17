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

type input struct {
	Element
	Children []HtmlElement
}

func Input(attrs ...func(HtmlElement)) HtmlElement {
	f := &input{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *input) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "input", f.Children)
}

func (f *input) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *input) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type textarea struct {
	Element
	Children []HtmlElement
}

func TextArea(attrs ...func(HtmlElement)) HtmlElement {
	f := &textarea{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *textarea) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "textarea", f.Children)
}

func (f *textarea) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *textarea) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type button struct {
	Element
	Children []HtmlElement
}

func Button(attrs ...func(HtmlElement)) HtmlElement {
	f := &button{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *button) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "button", f.Children)
}

func (f *button) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *button) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type selectelement struct {
	Element
	Children []HtmlElement
}

func Select(attrs ...func(HtmlElement)) HtmlElement {
	f := &selectelement{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *selectelement) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "select", f.Children)
}

func (f *selectelement) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *selectelement) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type option struct {
	Element
	Children []HtmlElement
}

func Option(attrs ...func(HtmlElement)) HtmlElement {
	f := &option{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *option) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "option", f.Children)
}

func (f *option) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *option) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type optgroup struct {
	Element
	Children []HtmlElement
}

func OptGroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &optgroup{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *optgroup) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "optgroup", f.Children)
}

func (f *optgroup) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *optgroup) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type label struct {
	Element
	Children []HtmlElement
}

func Label(attrs ...func(HtmlElement)) HtmlElement {
	f := &label{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *label) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "label", f.Children)
}

func (f *label) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *label) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}

type output struct {
	Element
	Children []HtmlElement
}

func Output(attrs ...func(HtmlElement)) HtmlElement {
	f := &output{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *output) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "output", f.Children)
}

func (f *output) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *output) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}
