package html

func DocType() Element {
	return &rawData{content: "<!DOCTYPE html>\n"}
}

func Html(attrs ...func(Element)) Element {
	return (&tag{tag: "html"}).addAttrs(attrs...)
}

func Head(attrs ...func(Element)) Element {
	return (&tag{tag: "head"}).addAttrs(attrs...)
}

func Meta(attrs ...func(Element)) Element {
	return (&tag{tag: "meta", empty: true}).addAttrs(attrs...)
}

func Base(attrs ...func(Element)) Element {
	return (&tag{tag: "base", empty: true}).addAttrs(attrs...)
}

func Title(attrs ...func(Element)) Element {
	return (&tag{tag: "title"}).addAttrs(attrs...)
}

func Body(attrs ...func(Element)) Element {
	return (&tag{tag: "body"}).addAttrs(attrs...)
}

// Content adds raw content to the stream.
func Content(data string) Element {
	return &rawData{content: data}
}
