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
		output:  "<b class=\"name\">something</b>",
		element: B(Class("name")).AddElements(Content("something")),
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
		output: "<ul><li><data value=\"1\">one</data></li><li><data value=\"2\">two</data></li></ul>",
		element: Ul().AddElements(
			Li().AddElements(
				Data(Value("1")).AddElements(Content("one")),
			),
			Li().AddElements(
				Data(Value("2")).AddElements(Content("two")),
			),
		),
	},
	{
		name:   "datalist",
		output: "<datalist class=\"a1\"><option value=\"One\"/><option value=\"two\"/></datalist>",
		element: DataList(Class("a1")).AddElements(
			Option(Value("One")),
			Option(Value("two")),
		),
	},
	{
		name:    "del",
		output:  "<del>delete</del>",
		element: Del().AddElements(Content("delete")),
	},
	{
		name:   "dfn",
		output: "<p><dfn>something</dfn> is not nothing</p>",
		element: P().AddElements(
			Dfn().AddElements(Content("something")),
			Content(" is not nothing"),
		),
	},
	{
		name:   "dialog",
		output: "<dialog><p>this is a dialog</p></dialog>",
		element: Dialog().AddElements(
			P().AddElements(Content("this is a dialog")),
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
	{
		name:    "em",
		output:  "<em>content</em>",
		element: Em().AddElements(Content("content")),
	},
	{
		name:    "hr",
		output:  "<hr>",
		element: Hr(),
	},
	{
		name:    "i",
		output:  "<i>content</i>",
		element: I().AddElements(Content("content")),
	},
	{
		name:    "ins",
		output:  "<ins>content</ins>",
		element: Ins().AddElements(Content("content")),
	},
	{
		name:    "kbd",
		output:  "<kbd>content</kbd>",
		element: Kbd().AddElements(Content("content")),
	},
	{
		name:    "object and param",
		output:  "<object><param name=\"autoplay\" value=\"true\"></object>",
		element: Object().AddElements(Param(Name("autoplay"), Value("true"))),
	},
	{
		name:    "noscript",
		output:  "<noscript>script not enabled</noscript>",
		element: NoScript().AddElements(Content("script not enabled")),
	},
	{
		name:    "pre",
		output:  "<pre>preformatted</pre>",
		element: Pre().AddElements(Content("preformatted")),
	},
	{
		name:    "progress",
		output:  "<progress id=\"p1\" value=\"27\" max=\"100\">27/100</progress>",
		element: Progress(Id("p1"), Value("27"), Max("100")).AddElements(Content("27/100")),
	},
	{
		name:   "q",
		output: "<p>content <q>short quotation</q> after quote</p>",
		element: P().AddElements(
			Content("content "),
			Q().AddElements(Content("short quotation")),
			Content(" after quote"),
		),
	},
	{
		name:   "ruby",
		output: "<ruby>漢 <rt><rp>(</rp>ㄏㄢˋ<rp>)</rp></rt></ruby>",
		element: Ruby().AddElements(
			Content("漢 "),
			Rt().AddElements(
				Rp().AddElements(Content("(")),
				Content("ㄏㄢˋ"),
				Rp().AddElements(Content(")")),
			),
		),
	},
	{
		name:    "s",
		output:  "<s>not here</s>",
		element: S().AddElements(Content("not here")),
	},
	{
		name:    "samp",
		output:  "<samp>example output</samp>",
		element: Samp().AddElements(Content("example output")),
	},
	{
		name:    "small",
		output:  "<small>small text</small>",
		element: Small().AddElements(Content("small text")),
	},
	{
		name:    "strong",
		output:  "<strong>i am strong</strong>",
		element: Strong().AddElements(Content("i am strong")),
	},
	{
		name:    "sub",
		output:  "<sub>i am sub</sub>",
		element: Sub().AddElements(Content("i am sub")),
	},
	{
		name:    "sup",
		output:  "<sup>i am sup</sup>",
		element: Sup().AddElements(Content("i am sup")),
	},
	{
		name:    "template",
		output:  "<template><p>templated section</p></template>",
		element: Template().AddElements(P().AddElements(Content("templated section"))),
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
