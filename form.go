package html

type form struct {
	Element
}

func Form(attrs ...func(HtmlElement)) HtmlElement {
	f := &form{}
	f.tag = "form"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *form) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type fieldset struct {
	Element
}

func Fieldset(attrs ...func(HtmlElement)) HtmlElement {
	f := &fieldset{}
	f.tag = "fieldset"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *fieldset) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type input struct {
	Element
}

func Input(attrs ...func(HtmlElement)) HtmlElement {
	f := &input{}
	f.tag = "input"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *input) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type textarea struct {
	Element
}

func TextArea(attrs ...func(HtmlElement)) HtmlElement {
	f := &textarea{}
	f.tag = "textarea"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *textarea) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type button struct {
	Element
}

func Button(attrs ...func(HtmlElement)) HtmlElement {
	f := &button{}
	f.tag = "button"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *button) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type selectelement struct {
	Element
}

func Select(attrs ...func(HtmlElement)) HtmlElement {
	f := &selectelement{}
	f.tag = "select"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *selectelement) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type option struct {
	Element
}

func Option(attrs ...func(HtmlElement)) HtmlElement {
	f := &option{}
	f.tag = "option"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *option) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type optgroup struct {
	Element
}

func OptGroup(attrs ...func(HtmlElement)) HtmlElement {
	f := &optgroup{}
	f.tag = "optgroup"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *optgroup) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type label struct {
	Element
}

func Label(attrs ...func(HtmlElement)) HtmlElement {
	f := &label{}
	f.tag = "label"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *label) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type output struct {
	Element
}

func Output(attrs ...func(HtmlElement)) HtmlElement {
	f := &output{}
	f.tag = "output"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *output) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}

type legend struct {
	Element
}

func Legend(attrs ...func(HtmlElement)) HtmlElement {
	f := &legend{}
	f.tag = "legend"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *legend) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.children = append(f.children, element)
	}
	return f
}
