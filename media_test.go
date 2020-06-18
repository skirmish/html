package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mediaTestCases = []HTMLTestCase{
	{
		name:    "canvas",
		output:  "<canvas/>",
		element: Canvas(),
	},
	{
		name:    "canvas with class",
		output:  "<canvas class=\"name\"/>",
		element: Canvas(Class("name")),
	},
	{
		name:    "canvas with content",
		output:  "<canvas>Not Supported!</canvas>",
		element: Canvas().AddElements(Content("Not Supported!")),
	},
	{
		name:    "image",
		output:  "<img>",
		element: Img(),
	},
	{
		name:    "image with src",
		output:  "<img src=\"asource\">",
		element: Img(Src("asource")),
	},
	{
		name:   "figure with bits",
		output: "<figure class=\"fig\"><img src=\"another\"><figcaption class=\"fc1\">Some caption</figcaption></figure>",
		element: Figure(Class("fig")).AddElements(
			Img(Src("another")),
			FigCaption(Class("fc1")).AddElements(Content("Some caption")),
		),
	},
	{
		name:   "audio with source",
		output: "<audio class=\"a\"><source src=\"something.ogg\" type=\"audio/ogg\"><source src=\"something.mp3\" type=\"audio/mp3\">audio not supported</audio>",
		element: Audio(Class("a")).AddElements(
			Source(Src("something.ogg"), Type("audio/ogg")),
			Source(Src("something.mp3"), Type("audio/mp3")),
			Content("audio not supported"),
		),
	},
	{
		name:   "video with source",
		output: "<video class=\"v\"><source src=\"something.mp4\" type=\"video/mp4\"><source src=\"something.ogg\" type=\"video/ogg\"><track src=\"subtitles_en.vtt\" kind=\"subtitles\" srclang=\"en\" label=\"English\">video not supported</video>",
		element: Video(Class("v")).AddElements(
			Source(Src("something.mp4"), Type("video/mp4")),
			Source(Src("something.ogg"), Type("video/ogg")),
			Track(Src("subtitles_en.vtt"), Kind("subtitles"), SrcLang("en"), LabelAttr("English")),
			Content("video not supported"),
		),
	},
	{
		name:   "picture",
		output: "<picture class=\"pic1\"><source srcset=\"some.jpg\"></picture>",
		element: Picture(Class("pic1")).AddElements(
			Source(SrcSet("some.jpg")),
		),
	},
	{
		name:    "embed",
		output:  "<embed src=\"asource.html\" type=\"text/html\"/>",
		element: Embed(Src("asource.html"), Type("text/html")),
	},
	{
		name:   "img, area,map",
		output: "<section><img src=\"asource.jpg\" usemap=\"#amap\"><map name=\"amap\"><area shape=\"rect\" coords=\"34,44,270,350\" href=\"one.html\"><area shape=\"circle\" coords=\"100,150,50\" href=\"two.html\"></map></section>",
		element: Section().AddElements(
			Img(Src("asource.jpg"), UseMap("#amap")),
			Map(Name("amap")).AddElements(
				Area(Shape("rect"), Coords("34,44,270,350"), Href("one.html")),
				Area(Shape("circle"), Coords("100,150,50"), Href("two.html")),
			),
		),
	},
	{
		name:   "svg",
		output: "<svg class=\"l\"><circle cx=\"50\" cy=\"50\" r=\"40\" stroke=\"blue\" stroke-width=\"4\" fill=\"yellow\" /></svg>",
		element: Svg(Class("l")).AddElements(
			Content("<circle cx=\"50\" cy=\"50\" r=\"40\" stroke=\"blue\" stroke-width=\"4\" fill=\"yellow\" />"),
		),
	},
	{
		name:    "link",
		output:  "<link rel=\"stylesheet\" href=\"styling.css\">",
		element: Link(Rel("stylesheet"), Href("styling.css")),
	},
}

func Test_MediaGeneration(t *testing.T) {
	for _, testCase := range mediaTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
