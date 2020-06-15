package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var blockquoteTestCases = []HTMLTestCase{
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

func Test_BlockquoteGeneration(t *testing.T) {
	for _, testCase := range blockquoteTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
