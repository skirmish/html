package html

import "io"

type canvas struct {
	Element
	Children []HtmlElement
}

func Canvas(attrs ...func(HtmlElement)) HtmlElement {
	i := &canvas{}
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

func (i *canvas) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}

type image struct {
	Element
	Children []HtmlElement
}

func Img(attrs ...func(HtmlElement)) HtmlElement {
	i := &image{}
	for _, attr := range attrs {
		attr(i)
	}
	return i
}

func (a *image) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "img", a.Children)
}

func (i *image) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	i.Children = append(i.Children,element)
	//}
	return i
}

func (i *image) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}

type figure struct {
	Element
	Children []HtmlElement
}

func Figure(attrs ...func(HtmlElement)) HtmlElement {
	i := &figure{}
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

func (i *figure) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}

type figcaption struct {
	Element
	Children []HtmlElement
}

func FigCaption(attrs ...func(HtmlElement)) HtmlElement {
	i := &figcaption{}
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

func (i *figcaption) addAttribute(key string, val string) {
	i.Element.AddAttribute(key, val)
}
