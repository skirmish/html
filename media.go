package html

func Canvas(attrs ...func(Element)) Element {
	return (&tag{tag: "canvas"}).addAttrs(attrs...)
}

func Img(attrs ...func(Element)) Element {
	return (&tag{tag: "img", empty: true}).addAttrs(attrs...)
}

func Figure(attrs ...func(Element)) Element {
	return (&tag{tag: "figure"}).addAttrs(attrs...)
}

func FigCaption(attrs ...func(Element)) Element {
	return (&tag{tag: "figcaption"}).addAttrs(attrs...)
}

func Audio(attrs ...func(Element)) Element {
	return (&tag{tag: "audio"}).addAttrs(attrs...)
}

func Embed(attrs ...func(Element)) Element {
	return (&tag{tag: "embed"}).addAttrs(attrs...)
}

func Video(attrs ...func(Element)) Element {
	return (&tag{tag: "video"}).addAttrs(attrs...)
}

func Picture(attrs ...func(Element)) Element {
	return (&tag{tag: "picture"}).addAttrs(attrs...)
}

func Svg(attrs ...func(Element)) Element {
	return (&tag{tag: "svg"}).addAttrs(attrs...)
}

func Area(attrs ...func(Element)) Element {
	return (&tag{tag: "area", empty: true}).addAttrs(attrs...)
}

func Map(attrs ...func(Element)) Element {
	return (&tag{tag: "map"}).addAttrs(attrs...)
}

func Source(attrs ...func(Element)) Element {
	return (&tag{tag: "source", empty: true}).addAttrs(attrs...)
}

func Link(attrs ...func(Element)) Element {
	return (&tag{tag: "link", empty: true}).addAttrs(attrs...)
}
