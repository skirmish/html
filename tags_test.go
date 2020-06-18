package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tagsTestCases = []HTMLTestCase{
	{
		name:   "abbr",
		output: "<p>The <abbr class=\"a\" title=\"abbrevaited text\">some</abbr> what</p>",
		element: P().AddElements(
			Content("The "),
			Abbr(Class("a"), TitleAttr("abbrevaited text")).AddElements(Content("some")),
			Content(" what"),
		),
	},
	{
		name:    "address",
		output:  "<address class=\"a\">The North Pole</address>",
		element: Address(Class("a")).AddElements(Content("The North Pole")),
	},
	{
		name:    "b",
		output:  "<b/>",
		element: B(),
	},
	{
		name:    "b with content and class",
		output:  "<b class=\"name\">something</b>",
		element: B(Class("name")).AddElements(Content("something")),
	},
	{
		name:    "br",
		output:  "<br>",
		element: Br(),
	},
	{
		name:    "b with content and class",
		output:  "<br class=\"name\">",
		element: Br(Class("name")).AddElements(Content("something")),
	},
	{
		name:    "bdi",
		output:  "<bdi class=\"b\">name</bdi>",
		element: Bdi(Class("b")).AddElements(Content("name")),
	},
	{
		name:    "bdo",
		output:  "<bdo dir=\"rtl\">wrong way</bdo>",
		element: Bdo(Dir("rtl")).AddElements(Content("wrong way")),
	},
	{
		name:   "cite",
		output: "<p><cite class=\"c\">data</cite> the wrong way</p>",
		element: P().AddElements(
			Cite(Class("c")).AddElements(Content("data")),
			Content(" the wrong way"),
		),
	},
	{
		name:   "code",
		output: "<p><code class=\"c\">data</code> another way</p>",
		element: P().AddElements(
			Code(Class("c")).AddElements(Content("data")),
			Content(" another way"),
		),
	},
	{
		name:   "data",
		output: "<p><code class=\"c\">data</code> another way</p>",
		element: Ul().AddElements(
			Code(Class("c")).AddElements(Content("data")),
			Content(" another way"),
		),
	},
	{
		name:    "Heading 1",
		output:  "<h1>TestContent</h1>",
		element: Heading(Level(1)).AddElements(Content("TestContent")),
	},
	{
		name:    "Heading 2",
		output:  "<h2>TestContent</h2>",
		element: Heading(Level(2)).AddElements(Content("TestContent")),
	},
	{
		name:    "Heading 3",
		output:  "<h3>TestContent</h3>",
		element: Heading(Level(3)).AddElements(Content("TestContent")),
	},
	{
		name:    "Heading 4",
		output:  "<h4>TestContent</h4>",
		element: Heading(Level(4)).AddElements(Content("TestContent")),
	},
	{
		name:    "Heading 5",
		output:  "<h5>TestContent</h5>",
		element: Heading(Level(5)).AddElements(Content("TestContent")),
	},
	{
		name:    "Heading 6",
		output:  "<h6>TestContent</h6>",
		element: Heading(Level(6)).AddElements(Content("TestContent")),
	},
	{
		name:    "meter",
		output:  "<meter/>",
		element: Meter(),
	},
	{
		name:   "meter with content",
		output: "<meter id=\"disk_c\" value=\"2\" min=\"0\" max=\"10\">2 out of 10</meter>",
		element: Meter(
			Id("disk_c"),
			Value("2"),
			Min("0"),
			Max("10"),
		).AddElements(
			Content("2 out of 10"),
		),
	},
	{
		name:   "meter with all attrs",
		output: "<meter id=\"disk_c\" value=\"2\" min=\"0\" max=\"10\" form=\"test\" low=\"2\" high=\"8\" optimum=\"6\">2 out of 10</meter>",
		element: Meter(
			Id("disk_c"),
			Value("2"),
			Min("0"),
			Max("10"),
			FormId("test"),
			Low("2"),
			High("8"),
			Optimum("6"),
		).AddElements(
			Content("2 out of 10"),
		),
	},
	{
		name:    "blockquote",
		output:  "<blockquote/>",
		element: Blockquote(),
	},
	{
		name:    "blockquote with content and class",
		output:  "<blockquote class=\"name\">something</blockquote>",
		element: Blockquote(Class("name")).AddElements(Content("something")),
	},
}

func Test_TagGeneration(t *testing.T) {
	for _, testCase := range tagsTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}