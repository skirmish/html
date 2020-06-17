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
		output:  "<img/>",
		element: Img(),
	},
	{
		name:    "image with src",
		output:  "<img src=\"asource\"/>",
		element: Img(Src("asource")),
	},
	{
		name:   "figure with bits",
		output: "<figure class=\"fig\"><img src=\"another\"/><figcaption class=\"fc1\">Some caption</figcaption></figure>",
		element: Figure(Class("fig")).AddElements(
			Img(Src("another")),
			FigCaption(Class("fc1")).AddElements(Content("Some caption")),
		),
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
