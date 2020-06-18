package v2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ElementTestCase struct {
	name    string
	output  string
	element Element
}

var tagTestCases = []ElementTestCase{
	{
		name:    "doctype",
		output:  "<!DOCTYPE html>\n",
		element: DocType(),
	},
	{
		name:   "page",
		output: "<html><head><meta charset=\"utf-8\"></head><body><h1>The Title</h1>Some data.</body></html>",
		element: Html().AddElements(
			Head().AddElements(
				Meta(Charset("utf-8")),
			),
			Body().AddElements(
				H(Level(1)).AddElements(Content("The Title")),
				Content("Some data."),
			),
		),
	},
}

func Test_TagGeneration(t *testing.T) {
	for _, testCase := range tagTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
