package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var articleTestCases = []HTMLTestCase{
	{
		name:    "article",
		output:  "<article/>",
		element: Article(),
	},
	{
		name:    "article with content and class",
		output:  "<article class=\"name\">something</article>",
		element: Article(Class("name")).AddElements(Content("something")),
	},
}

func Test_ArticleGeneration(t *testing.T) {
	for _, testCase := range articleTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
