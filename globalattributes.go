package html

//HTML Global Attributes
//The global attributes are attributes that can be used with all HTML elements.
//
//Attribute	Description
//accesskey	Specifies a shortcut key to activate/focus an element
//class	Specifies one or more classnames for an element (refers to a class in a style sheet)
//contenteditable	Specifies whether the content of an element is editable or not
//data-*	Used to store custom data private to the page or application
//dir	Specifies the text direction for the content in an element
//draggable	Specifies whether an element is draggable or not
//hidden	Specifies that an element is not yet, or is no longer, relevant
//id	Specifies a unique id for an element
//lang	Specifies the language of the element's content
//spellcheck	Specifies whether the element is to have its spelling and grammar checked or not
//style	Specifies an inline CSS style for an element
//tabindex	Specifies the tabbing order of an element
//title	Specifies extra information about an element
//translate	Specifies whether the content of an element should be translated or not

func AccessKey(accesskey string) func(e HtmlElement) {
	return func(e HtmlElement) {
		e.addAttribute("accesskey", accesskey)
	}
}

func Class(class string) func(e HtmlElement) {
	return func(e HtmlElement) {
		e.addAttribute("class", class)
	}
}

func Hidden(hidden string) func(e HtmlElement) {
	return func(e HtmlElement) {
		e.addAttribute("hidden", hidden)
	}
}

func Id(id string) func(element HtmlElement) {
	return func(element HtmlElement) {
		element.addAttribute("id", id)
	}
}

func Lang(lang string) func(element HtmlElement) {
	return func(element HtmlElement) {
		element.addAttribute("lang", lang)
	}
}

func TitleAttr(title string) func(element HtmlElement) {
	return func(element HtmlElement) {
		element.addAttribute("title", title)
	}
}
