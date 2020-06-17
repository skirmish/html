package html

import "io"

type canvas struct {
	Element
	Children []HtmlElement
}

func Canvas(attrs ...func(HtmlElement)) HtmlElement {
	i := &canvas{}
	i.tag = "canvas"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *canvas) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "canvas", a.Children)
}

func (i *canvas) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type image struct {
	Element
	Children []HtmlElement
}

func Img(attrs ...func(HtmlElement)) HtmlElement {
	i := &image{}
	i.tag = "img"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *image) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "img", a.Children)
}

func (i *image) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type figure struct {
	Element
	Children []HtmlElement
}

func Figure(attrs ...func(HtmlElement)) HtmlElement {
	i := &figure{}
	i.tag = "figure"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *figure) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "figure", a.Children)
}

func (i *figure) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type figcaption struct {
	Element
	Children []HtmlElement
}

func FigCaption(attrs ...func(HtmlElement)) HtmlElement {
	i := &figcaption{}
	i.tag = "figcaption"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *figcaption) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "figcaption", a.Children)
}

func (i *figcaption) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type audio struct {
	Element
	Children []HtmlElement
}

func Audio(attrs ...func(HtmlElement)) HtmlElement {
	i := &audio{}
	i.tag = "audio"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *audio) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "audio", a.Children)
}

func (i *audio) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type embed struct {
	Element
	Children []HtmlElement
}

func Embed(attrs ...func(HtmlElement)) HtmlElement {
	i := &embed{}
	i.tag = "embed"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *embed) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "embed", a.Children)
}

func (i *embed) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type video struct {
	Element
	Children []HtmlElement
}

func Video(attrs ...func(HtmlElement)) HtmlElement {
	i := &video{}
	i.tag = "video"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *video) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "video", a.Children)
}

func (i *video) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type picture struct {
	Element
	Children []HtmlElement
}

func Picture(attrs ...func(HtmlElement)) HtmlElement {
	i := &picture{}
	i.tag = "picture"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *picture) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "picture", a.Children)
}

func (i *picture) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type svg struct {
	Element
	Children []HtmlElement
}

func Svg(attrs ...func(HtmlElement)) HtmlElement {
	i := &svg{}
	i.tag = "svg"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *svg) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "svg", a.Children)
}

func (i *svg) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type area struct {
	Element
	Children []HtmlElement
}

func Area(attrs ...func(HtmlElement)) HtmlElement {
	i := &area{}
	i.tag = "area"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *area) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "area", a.Children)
}

func (i *area) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}

type mapElement struct {
	Element
	Children []HtmlElement
}

func Map(attrs ...func(HtmlElement)) HtmlElement {
	i := &mapElement{}
	i.tag = "map"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *mapElement) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "map", a.Children)
}

func (i *mapElement) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		i.Children = append(i.Children, element)
	}
	return i
}

type source struct {
	Element
	Children []HtmlElement
}

func Source(attrs ...func(HtmlElement)) HtmlElement {
	i := &source{}
	i.tag = "source"
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *source) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "source", a.Children)
}

func (i *source) AddElements(elements ...HtmlElement) HtmlElement {
	return i
}
