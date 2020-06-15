package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var canvasTestCases = []HTMLTestCase{
	{
		name:    "canvas",
		output:  "<canvas/>",
		element: Canvas(),
	},
	{
		name:    "blockquote with class",
		output:  "<canvas class=\"name\"/>",
		element: Canvas(Class("name")),
	},
}

func Test_CanvasGeneration(t *testing.T) {
	for _, testCase := range canvasTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
