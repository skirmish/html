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
			"<ul class=\"samples\"><li><img src=\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABKQAAASkCAYAAABdI9LQAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAgAElEQVR4nOzasYqcZRiG4RlZxIQIEcTao7ARtojY2KS19zA8AAtrS60DYmEhCIIIYmlrlT4gCzuYf50lRcbOQoRgivv9Jl7XETzV7vff8+52AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwKtvPz0AgPPw6TeXp+kNK3j+6O70BBby2dffe0sBALyE16YHAAAAAPD/IkgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACC1nx4AsLrPf/7oNL1hBYfft+kJS3j+6O70BBbyx5+30xOW8MV3P3lTAgD/iQspAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUhfTAwA4D/ffuTc9YQmHj7fpCUs4fOU3LQAAXp7XJAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAKn99AAAzsZpegDr+OTD96cnsJAvf/hlegJr8Y0BwAu5kAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACC1nx4ALO00PYB1nB4+mJ7AQvbf/jg9gYVsH7w3PWEJV4dtesIS3v31t+kJrMU3J/CvXEgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQ2k8PgEWdpges4PTwwfQEWM7Ntk1PABZ1dfD3Ybfb7a6Px+kJS3j67HZ6whIuHz+ZnrAK397wDy6kAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASO2nB7Cc0/SAFZwePpiewEJutm16AgBn4Org/8Vut9tdH4/TE5bw9Nnt9AQWcvn4yfSEVWgQ/M2FFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAKmL6QHAum62bXoCAMBZevP1N6YnLOHps9vpCcCiXEgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQupgeAAAAvDreunNnesISro/H6QkAS3MhBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQOpiegCs6GbbpicAAGfm7fv3pics4ergHQXAi7mQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAIHUxPQAAAF4FV4dtegIAnA0XUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQupgfAiq4O2/QEAICzdH08Tk8A4Ay4kAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFJ/tWvHpg1EURQFWXATylWHilcdyhU6cCDjZJ0uQuDEnPcFMxXc6LN7eIIUAAAAAClBCgAAAICUIAUAAABASpACAAAAIPUxPQBW9Pl4TE8AAOCNff18T08AWJoLKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFLb9ABY1D49YAXX82l6AgAAb+xyu09PWIV/b3jiQgoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAIDUNj0AWNo+PYB1XM+n6Qks5HK7T09gId4HjrwPPPHPCbzkQgoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAIDUNj0AgLexTw9gKb4hOPI+cOR9AOBPLqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAABICVIAAAAApAQpAAAAAFKCFAAAAAApQQoAAACAlCAFAAAAQEqQAgAAACAlSAEAAACQEqQAAAAASAlSAAAAAKQEKQAAAABSghQAAAAAKUEKAAAAgJQgBQAAAEBKkAIAAAAgJUgBAAAAkBKkAAAAAEgJUgAAAACkBCkAAAAAUoIUAAAAAClBCgAAAICUIAUAAABASpACAAAAICVIAQAAAJASpAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALpo2esAAADMSURBVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+Ge/YHx2PVYVb+8AAAAASUVORK5CYII=\"/></li></ul></html>\n",
		element: Html(Lang("en")).AddElements(
			Head().AddElements(
				Meta(K("charset", "utf-8")),
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
					Img(Src("data:image/png;base64,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")),
				),
				Li(Class("good")).AddElements(
					Img(Src("data:image/png;base64,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")),
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
