package html

import (
	"bytes"
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
		output:  "<html/>",
		element: Html(),
	},
	{
		name:    "html lang=en",
		output:  "<html lang=\"en\"/>",
		element: Html(Lang("en")),
	},
	{
		name:    "form",
		output:  "<form/>",
		element: Form(),
	},
	{
		name:    "form with id",
		output:  "<form id=\"formId1\"/>",
		element: Form(Id("formId1")),
	},
	{
		name:    "form with id,name",
		output:  "<form id=\"formId1\" name=\"somename\"/>",
		element: Form(Id("formId1"), Name("somename")),
	},
	{
		name:    "Page with head and body",
		output:  "<html><head/><body/></html>",
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
		output:  "<script/>",
		element: Script(),
	},
}

func Test_HtmlGeneration(t *testing.T) {
	for _, testCase := range htmlTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}

func runBenchmarkCases(b *testing.B, cases []HTMLTestCase) {
	for _, benchCase := range cases {
		b.Run(benchCase.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf := new(bytes.Buffer)
				n, _ := benchCase.element.Render(buf)
				b.ReportMetric(float64(n), "bytes")
			}
		})
	}
}

func BenchmarkHtmlGeneration(b *testing.B) {
	runBenchmarkCases(b, htmlTestCases)
	runBenchmarkCases(b, sectionTestCases)
	runBenchmarkCases(b, headingTestCases)
	runBenchmarkCases(b, meterTestCases)
	runBenchmarkCases(b, articleTestCases)
	runBenchmarkCases(b, asideTestCases)
	//runBenchmarkCases(b, kitchenSinkTestCases)
}
