package html

func Ul(attrs ...func(Element)) Element {
	ul := &tag{}
	ul.tag = "ul"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func Ol(attrs ...func(Element)) Element {
	ul := &tag{}
	ul.tag = "ol"
	for _, attr := range attrs {
		attr(ul)
	}
	return ul
}

func Li(attrs ...func(Element)) Element {
	li := &tag{}
	li.tag = "li"
	for _, attr := range attrs {
		attr(li)
	}
	return li
}

func Dl(attrs ...func(Element)) Element {
	f := &tag{}
	f.tag = "dl"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func Dt(attrs ...func(Element)) Element {
	f := &tag{}
	f.tag = "dt"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func Dd(attrs ...func(Element)) Element {
	f := &tag{}
	f.tag = "dd"
	for _, attr := range attrs {
		attr(f)
	}
	return f
}
