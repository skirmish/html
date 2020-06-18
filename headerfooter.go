package html

func Header(attrs ...func(Element)) Element {
	return (&tag{tag: "header"}).addAttrs(attrs...)
}

func Footer(attrs ...func(Element)) Element {
	return (&tag{tag: "footer"}).addAttrs(attrs...)
}

func Article(attrs ...func(Element)) Element {
	return (&tag{tag: "article"}).addAttrs(attrs...)
}

func Aside(attrs ...func(Element)) Element {
	return (&tag{tag: "aside"}).addAttrs(attrs...)
}

func Section(attrs ...func(Element)) Element {
	return (&tag{tag: "section"}).addAttrs(attrs...)
}

func Nav(attrs ...func(Element)) Element {
	return (&tag{tag: "nav"}).addAttrs(attrs...)
}

func Details(attrs ...func(Element)) Element {
	return (&tag{tag: "details"}).addAttrs(attrs...)
}

func Main(attrs ...func(Element)) Element {
	return (&tag{tag: "main"}).addAttrs(attrs...)
}

func Mark(attrs ...func(Element)) Element {
	return (&tag{tag: "mark"}).addAttrs(attrs...)
}

func Summary(attrs ...func(Element)) Element {
	return (&tag{tag: "summary"}).addAttrs(attrs...)
}

func Time(attrs ...func(Element)) Element {
	return (&tag{tag: "time"}).addAttrs(attrs...)
}
