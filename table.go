package html

func Table(attrs ...func(Element)) Element {
	return (&tag{tag: "table"}).addAttrs(attrs...)
}

func Tr(attrs ...func(Element)) Element {
	return (&tag{tag: "tr"}).addAttrs(attrs...)
}

func Th(attrs ...func(Element)) Element {
	return (&tag{tag: "th"}).addAttrs(attrs...)
}

func Td(attrs ...func(Element)) Element {
	return (&tag{tag: "td"}).addAttrs(attrs...)
}

func Caption(attrs ...func(Element)) Element {
	return (&tag{tag: "caption"}).addAttrs(attrs...)
}

func Colgroup(attrs ...func(Element)) Element {
	return (&tag{tag: "colgroup"}).addAttrs(attrs...)
}

func Col(attrs ...func(Element)) Element {
	return (&tag{tag: "col"}).addAttrs(attrs...)
}

func Thead(attrs ...func(Element)) Element {
	return (&tag{tag: "thead"}).addAttrs(attrs...)
}

func Tfoot(attrs ...func(Element)) Element {
	return (&tag{tag: "tfoot"}).addAttrs(attrs...)
}

func Tbody(attrs ...func(Element)) Element {
	return (&tag{tag: "tbody"}).addAttrs(attrs...)
}
