package html

type docType struct {
	Element
}

func DocType() HtmlElement {
	return &docType{}
}

func (d *docType) Render() string {
	output := "<!DOCTYPE html>\n"
	return output
}

func (d *docType) AddElements(elements ...HtmlElement) HtmlElement {
	return d
}

func (d *docType) addAttribute(key string, val string) {
	d.Element.AddAttribute(key, val)
}
