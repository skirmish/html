package html

import "io"

type form struct {
	Element
	Children []HtmlElement
}

func Form(attrs ...func(HtmlElement)) HtmlElement {
	f := &form{}
	f.tag = "form"
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

type fieldset struct {
	Element
	Children []HtmlElement
}

func Fieldset(attrs ...func(HtmlElement)) HtmlElement {
	f := &fieldset{}
	f.tag = "fieldset"
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

type input struct {
	Element
	Children []HtmlElement
}

func Input(attrs ...func(HtmlElement)) HtmlElement {
	f := &input{}
	f.tag = "input"
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

type textarea struct {
	Element
	Children []HtmlElement
}

func TextArea(attrs ...func(HtmlElement)) HtmlElement {
	f := &textarea{}
	f.tag = "textarea"
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

type button struct {
	Element
	Children []HtmlElement
}

func Button(attrs ...func(HtmlElement)) HtmlElement {
	f := &button{}
	f.tag = "button"
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

type selectelement struct {
	Element
	Children []HtmlElement
}

func Select(attrs ...func(HtmlElement)) HtmlElement {
	f := &selectelement{}
	f.tag = "select"
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

type option struct {
	Element
	Children []HtmlElement
}

func Option(attrs ...func(HtmlElement)) HtmlElement {
	f := &option{}
	f.tag = "option"
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

type optgroup struct {
	Element
	Children []HtmlElement
}

func OptGroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &optgroup{}
	f.tag = "optgroup"
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

type label struct {
	Element
	Children []HtmlElement
}

func Label(attrs ...func(HtmlElement)) HtmlElement {
	f := &label{}
	f.tag = "label"
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

type output struct {
	Element
	Children []HtmlElement
}

func Output(attrs ...func(HtmlElement)) HtmlElement {
	f := &output{}
	f.tag = "output"
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

type legend struct {
	Element
	Children []HtmlElement
}

func Legend(attrs ...func(HtmlElement)) HtmlElement {
	f := &legend{}
	f.tag = "legend"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *legend) Render(w io.Writer) (int, error) {
	return f.Element.Render(w, "legend", f.Children)
}

func (f *legend) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}
