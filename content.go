package html

type content struct {
	Element
	Content string
}

func Content(data string) HtmlElement {
	c := &content{Content: data}
	return c
}

func (c *content) Render() string {
	output := c.Content
	return output
}

func (c *content) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	c.Children = append(c.Children,element)
	//}
	return c
}

func (c *content) addAttribute(key string, val string) {
	c.Element.AddAttribute(key, val)
}
