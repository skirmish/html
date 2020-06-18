package html

import "fmt"

// em, hr, i, ins, kbd, link, noscript, object, param, pre, progress,
// q, rp, rt, ruby, s, samp, small, strong, sub, sup, template, track, tt, u, var, wbr

func Abbr(attrs ...func(Element)) Element {
	return (&tag{tag: "abbr"}).addAttrs(attrs...)
}

func Address(attrs ...func(Element)) Element {
	return (&tag{tag: "address"}).addAttrs(attrs...)
}

func B(attrs ...func(Element)) Element {
	return (&tag{tag: "b"}).addAttrs(attrs...)
}

func Bdi(attrs ...func(Element)) Element {
	return (&tag{tag: "bdi"}).addAttrs(attrs...)
}

func Bdo(attrs ...func(Element)) Element {
	return (&tag{tag: "bdo"}).addAttrs(attrs...)
}

func Blockquote(attrs ...func(Element)) Element {
	return (&tag{tag: "blockquote"}).addAttrs(attrs...)
}

func Br(attrs ...func(Element)) Element {
	return (&tag{tag: "br", empty: true}).addAttrs(attrs...)
}

func Cite(attrs ...func(Element)) Element {
	return (&tag{tag: "cite"}).addAttrs(attrs...)
}

func Code(attrs ...func(Element)) Element {
	return (&tag{tag: "code"}).addAttrs(attrs...)
}

func Data(attrs ...func(Element)) Element {
	return (&tag{tag: "data"}).addAttrs(attrs...)
}

func DataList(attrs ...func(Element)) Element {
	return (&tag{tag: "datalist"}).addAttrs(attrs...)
}

func Del(attrs ...func(Element)) Element {
	return (&tag{tag: "del"}).addAttrs(attrs...)
}

func Dfn(attrs ...func(Element)) Element {
	return (&tag{tag: "dfn"}).addAttrs(attrs...)
}

func Dialog(attrs ...func(Element)) Element {
	return (&tag{tag: "dialog"}).addAttrs(attrs...)
}

func Heading(attrs ...func(*tag)) Element {
	h := &tag{}
	for _, attr := range attrs {
		attr(h)
	}
	h.tag = fmt.Sprintf("h%d", h.level)
	return h
}

func Level(level int) func(*tag) {
	return func(h *tag) { h.level = level }
}

func Meter(attrs ...func(Element)) Element {
	return (&tag{tag: "meter"}).addAttrs(attrs...)
}

func Script(attrs ...func(Element)) Element {
	return (&tag{tag: "script"}).addAttrs(attrs...)
}

func Style(attrs ...func(Element)) Element {
	return (&tag{tag: "style"}).addAttrs(attrs...)
}

func P(attrs ...func(Element)) Element {
	return (&tag{tag: "p"}).addAttrs(attrs...)
}
