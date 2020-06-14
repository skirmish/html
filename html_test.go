package html

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type HTMLTestCase struct {
	name    string
	output  string
	element HtmlElement
}

var htmlTestCases = []HTMLTestCase{
	{
		name:    "doctype",
		output:  "<!DOCTYPE html>\n",
		element: DocType(),
	},
	{
		name:    "html",
		output:  "<html>\n</html>\n",
		element: Html(),
	},
	{
		name:    "html lang=en",
		output:  "<html lang=\"en\">\n</html>\n",
		element: Html(Lang("en")),
	},
	{
		name:    "form",
		output:  "<form></form>",
		element: Form(),
	},
	{
		name:    "form with id",
		output:  "<form id=\"formId1\"></form>",
		element: Form(Id("formId1")),
	},
	{
		name:    "form with id,name",
		output:  "<form id=\"formId1\" name=\"somename\"></form>",
		element: Form(Id("formId1"), Name("somename")),
	},
	{
		name:    "Page with head and body",
		output:  "<html>\n<head></head>\n<body></body></html>\n",
		element: Html().AddElements(Head(), Body()),
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
		name:    "script",
		output:  "<script></script>\n",
		element: Script(),
	},
}

func Test_HtmlGeneration(t *testing.T) {
	for _, testCase := range htmlTestCases {
		t.Logf("Test %s", testCase.name)

		output := testCase.element.Render()
		assert.Equal(t, testCase.output, output)
	}
}

func runBenchmarkCases(b *testing.B, cases []HTMLTestCase) {
	for _, benchCase := range cases {
		b.Run(benchCase.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				output := benchCase.element.Render()
				b.ReportMetric(float64(len(output)), "bytes")
			}
		})
	}
}
func BenchmarkHtmlGeneration(b *testing.B) {
	runBenchmarkCases(b, htmlTestCases)
	runBenchmarkCases(b, sectionTestCases)
	runBenchmarkCases(b, headingTestCases)
	runBenchmarkCases(b, meterTestCases)
	//runBenchmarkCases(b,kitchenSinkTestCases)
}
