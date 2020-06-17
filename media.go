package html

type canvas struct {
	Element
}

func Canvas(attrs ...func(HtmlElement)) HtmlElement {
	i := &canvas{}
	i.tag = "canvas"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *canvas) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type image struct {
	Element
}

func Img(attrs ...func(HtmlElement)) HtmlElement {
	i := &image{}
	i.tag = "img"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *image) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type figure struct {
	Element
}

func Figure(attrs ...func(HtmlElement)) HtmlElement {
	i := &figure{}
	i.tag = "figure"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *figure) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type figcaption struct {
	Element
}

func FigCaption(attrs ...func(HtmlElement)) HtmlElement {
	i := &figcaption{}
	i.tag = "figcaption"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *figcaption) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type audio struct {
	Element
}

func Audio(attrs ...func(HtmlElement)) HtmlElement {
	i := &audio{}
	i.tag = "audio"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *audio) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type embed struct {
	Element
}

func Embed(attrs ...func(HtmlElement)) HtmlElement {
	i := &embed{}
	i.tag = "embed"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *embed) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type video struct {
	Element
}

func Video(attrs ...func(HtmlElement)) HtmlElement {
	i := &video{}
	i.tag = "video"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *video) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type picture struct {
	Element
}

func Picture(attrs ...func(HtmlElement)) HtmlElement {
	i := &picture{}
	i.tag = "picture"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *picture) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type svg struct {
	Element
}

func Svg(attrs ...func(HtmlElement)) HtmlElement {
	i := &svg{}
	i.tag = "svg"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *svg) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type area struct {
	Element
}

func Area(attrs ...func(HtmlElement)) HtmlElement {
	i := &area{}
	i.tag = "area"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *area) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type mapElement struct {
	Element
}

func Map(attrs ...func(HtmlElement)) HtmlElement {
	i := &mapElement{}
	i.tag = "map"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *mapElement) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.children = append(i.children, element)
	}
	return i
}

type source struct {
	Element
}

func Source(attrs ...func(HtmlElement)) HtmlElement {
	i := &source{}
	i.tag = "source"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (i *source) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}
