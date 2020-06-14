package html

type paragraph struct {
	Element
	Children []HtmlElement
}

func P(attrs ...func(HtmlElement)) HtmlElement {
	p := &paragraph{}
	for _, attr := range attrs {
		attr(p)
	}
	return p
}

func (p *paragraph) Render() string {
	output := "<p"
	p.InitAttributes()
	output += p.Element.RenderAttr()
	output += ">"
	for _, child := range p.Children {
		output += child.Render()
	}
	output += "</p>"
	return output
}

func (p *paragraph) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		p.Children = append(p.Children, element)
	}
	return p
}

func (p *paragraph) addAttribute(key string, val string) {
	p.Element.AddAttribute(key, val)
}
