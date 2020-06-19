package html

import "fmt"

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

func Em(attrs ...func(Element)) Element {
	return (&tag{tag: "em"}).addAttrs(attrs...)
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

func Hr(attrs ...func(Element)) Element {
	return (&tag{tag: "hr", empty: true}).addAttrs(attrs...)
}

func I(attrs ...func(Element)) Element {
	return (&tag{tag: "i"}).addAttrs(attrs...)
}

func Ins(attrs ...func(Element)) Element {
	return (&tag{tag: "ins"}).addAttrs(attrs...)
}

func Kbd(attrs ...func(Element)) Element {
	return (&tag{tag: "kbd"}).addAttrs(attrs...)
}

func Param(attrs ...func(Element)) Element {
	return (&tag{tag: "param", empty: true}).addAttrs(attrs...)
}

func Meter(attrs ...func(Element)) Element {
	return (&tag{tag: "meter"}).addAttrs(attrs...)
}

func NoScript(attrs ...func(Element)) Element {
	return (&tag{tag: "noscript"}).addAttrs(attrs...)
}

func Object(attrs ...func(Element)) Element {
	return (&tag{tag: "object"}).addAttrs(attrs...)
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

func Pre(attrs ...func(Element)) Element {
	return (&tag{tag: "pre"}).addAttrs(attrs...)
}

func Progress(attrs ...func(Element)) Element {
	return (&tag{tag: "progress"}).addAttrs(attrs...)
}

func Q(attrs ...func(Element)) Element {
	return (&tag{tag: "q"}).addAttrs(attrs...)
}

func Rp(attrs ...func(Element)) Element {
	return (&tag{tag: "rp"}).addAttrs(attrs...)
}

func Rt(attrs ...func(Element)) Element {
	return (&tag{tag: "rt"}).addAttrs(attrs...)
}

func Ruby(attrs ...func(Element)) Element {
	return (&tag{tag: "ruby"}).addAttrs(attrs...)
}

func S(attrs ...func(Element)) Element {
	return (&tag{tag: "s"}).addAttrs(attrs...)
}

func Samp(attrs ...func(Element)) Element {
	return (&tag{tag: "samp"}).addAttrs(attrs...)
}

func Small(attrs ...func(Element)) Element {
	return (&tag{tag: "small"}).addAttrs(attrs...)
}

func Strong(attrs ...func(Element)) Element {
	return (&tag{tag: "strong"}).addAttrs(attrs...)
}

func Sub(attrs ...func(Element)) Element {
	return (&tag{tag: "sub"}).addAttrs(attrs...)
}

func Sup(attrs ...func(Element)) Element {
	return (&tag{tag: "sup"}).addAttrs(attrs...)
}

func Template(attrs ...func(Element)) Element {
	return (&tag{tag: "template"}).addAttrs(attrs...)
}

func U(attrs ...func(Element)) Element {
	return (&tag{tag: "u"}).addAttrs(attrs...)
}

func Var(attrs ...func(Element)) Element {
	return (&tag{tag: "var"}).addAttrs(attrs...)
}

func Wbr(attrs ...func(Element)) Element {
	return (&tag{tag: "wbr", empty: true}).addAttrs(attrs...)
}
