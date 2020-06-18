package html

func Form(attrs ...func(Element)) Element {
	return (&tag{tag: "form"}).addAttrs(attrs...)
}

func Fieldset(attrs ...func(Element)) Element {
	return (&tag{tag: "fieldset"}).addAttrs(attrs...)
}

func Input(attrs ...func(Element)) Element {
	return (&tag{tag: "input"}).addAttrs(attrs...)
}

func TextArea(attrs ...func(Element)) Element {
	return (&tag{tag: "textarea"}).addAttrs(attrs...)
}

func Button(attrs ...func(Element)) Element {
	return (&tag{tag: "button"}).addAttrs(attrs...)
}

func Select(attrs ...func(Element)) Element {
	return (&tag{tag: "select"}).addAttrs(attrs...)
}

func Option(attrs ...func(Element)) Element {
	return (&tag{tag: "option"}).addAttrs(attrs...)
}

func OptGroup(attrs ...func(Element)) Element {
	return (&tag{tag: "optgroup"}).addAttrs(attrs...)
}

func Label(attrs ...func(Element)) Element {
	return (&tag{tag: "label"}).addAttrs(attrs...)
}

func Output(attrs ...func(Element)) Element {
	return (&tag{tag: "output"}).addAttrs(attrs...)
}

func Legend(attrs ...func(Element)) Element {
	return (&tag{tag: "legend"}).addAttrs(attrs...)
}
