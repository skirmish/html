package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tagsTestCases = []HTMLTestCase{
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
		output:  "<br/>",
		element: Br(),
	},
	{
		name:    "b with content and class",
		output:  "<br class=\"name\"/>",
		element: Br(Class("name")).AddElements(Content("something")),
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
