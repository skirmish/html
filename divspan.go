package html

func Div(attrs ...func(Element)) Element {
	return (&tag{tag: "div"}).addAttrs(attrs...)
}

func Span(attrs ...func(Element)) Element {
	return (&tag{tag: "span"}).addAttrs(attrs...)
}
