package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var kitchenSinkTestCases = []HTMLTestCase{
	{
		name: "kitchen",
		output: "<html lang=\"en\">\n<head><meta charset=\"utf-8\"/><title>Page Title</title>" +
			"<style>ul.samples {list-style: none; margin: 0; padding: 0; margin-block-start: 1px;" +
			" margin-block-end: 1px;}\n        ul.samples li {display: inline-block; width: 300px;}\n" +
			"        ul.samples li.good {background-color: #008000;}\n        </style></head>\n" +
			"<body><h1>The Title of the Page</h1><p>This is the first sentence.</p></body>" +
			"<ul class=\"samples\"><li><img src=\"data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcC" +
			"AAAAABXZoBIAAAA7UlEQVR4nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61" +
			"/SV4Zn728imyh4&#43;y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckIN" +
			"TNvDD37/lErhkdXf9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS4" +
			"01VwyAdiBpLganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFT" +
			"kSuQmCC\"/></li><li class=\"good\"><img src=\"data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAA" +
			"ABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGq" +
			"JzBPn4o61/SV4Zn728imyh4&#43;y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JR" +
			"mUPz5ckINTNvDD37/lErhkdXf9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oK" +
			"C96QfppS401VwyAdiBpLganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4M" +
			"AAAAAElFTkSuQmCC\"/></li></ul></html>\n",
		element: Html(Lang("en")).AddElements(
			Head().AddElements(
				Meta(Key("charset", "utf-8")),
				Title().AddElements(Content("Page Title")),
				Style().AddElements(Content(
					"ul.samples {list-style: none; margin: 0; padding: 0;"+
						" margin-block-start: 1px; margin-block-end: 1px;}\n"+
						"ul.samples li {display: inline-block; width: 300px;}\n"+
						"ul.samples li.good {background-color: #008000;}\n")),
			),
			Body().AddElements(
				Heading(Level(1)).AddElements(Content("The Title of the Page")),
				P().AddElements(Content("This is the first sentence."))),
			Ul(Class("samples")).AddElements(
				Li().AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
				Li(Class("good")).AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
			),
		),
	},
}

func Test_KitchenSink(t *testing.T) {
	page := []HtmlElement{
		DocType(),
		Html(Lang("en")).AddElements(
			Head().AddElements(
				Meta(Key("charset", "utf-8")),
				Title().AddElements(Content("Page Title")),
				Style().AddElements(Content(
					"ul.samples {list-style: none; margin: 0; padding: 0;"+
						" margin-block-start: 1px; margin-block-end: 1px;}\n"+
						"ul.samples li {display: inline-block; width: 300px;}\n"+
						"ul.samples li.good {background-color: #008000;}\n")),
			),
			Body().AddElements(
				Heading(Level(1)).AddElements(Content("The Title of the Page")),
				P().AddElements(Content("This is the first sentence."))),
			Ul(Class("samples")).AddElements(
				Li().AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
				Li(Class("good")).AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
			),
		),
	}

	buf := new(bytes.Buffer)
	for _, el := range page {

		_, err := el.Render(buf)
		assert.NoError(t, err, "Rendering")

	}
	t.Log(buf.String())
}
