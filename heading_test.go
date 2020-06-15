package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var headingTestCases = []HTMLTestCase{
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
}

func Test_HeadingGeneration(t *testing.T) {
	for _, testCase := range headingTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
