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
		name:    "body",
		output:  "<body/>",
		element: Body(),
	},
	{
		name:    "body with class",
		output:  "<body class=\"name\"/>",
		element: Body(Class("name")),
	},
	{
		name:    "body with class and content",
		output:  "<body class=\"name\">content</body>",
		element: Body(Class("name")).AddElements(Content("content")),
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
		name:    "fieldset",
		output:  "<fieldset/>",
		element: Fieldset(),
	},
	{
		name:    "fieldset with formid",
		output:  "<fieldset form=\"someform\"/>",
		element: Fieldset(FormId("someform")),
	},
	{
		name:   "form with fieldset with formid",
		output: "<form><fieldset form=\"someform\"/></form>",
		element: Form().AddElements(
			Fieldset(FormId("someform")),
		),
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
	{
		name:    "script with type and content",
		output:  "<script type=\"text/javascript\">alert(1);\n</script>",
		element: Script(Type("text/javascript")).AddElements(Content("alert(1);\n")),
	},
	{
		name:    "p",
		output:  "<p/>",
		element: P(),
	},
	{
		name:    "p with content",
		output:  "<p>content</p>",
		element: P().AddElements(Content("content")),
	},
	{
		name:    "p with class and content",
		output:  "<p class=\"name\">content</p>",
		element: P(Class("name")).AddElements(Content("content")),
	},
	{
		name:    "style with content",
		output:  "<style>content</style>",
		element: Style().AddElements(Content("content")),
	},
	{
		name:    "style with type content",
		output:  "<style type=\"text/css\">content</style>",
		element: Style(Type("text/css")).AddElements(Content("content")),
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
				b.SetBytes(int64(n))
			}
		})
	}
}

func BenchmarkHtmlGeneration(b *testing.B) {
	runBenchmarkCases(b, htmlTestCases)
	runBenchmarkCases(b, headingTestCases)
	runBenchmarkCases(b, meterTestCases)
	runBenchmarkCases(b, headerfooterTestCases)
	runBenchmarkCases(b, divspanTestCases)
	runBenchmarkCases(b, brTestCases)
	runBenchmarkCases(b, canvasTestCases)
	runBenchmarkCases(b, listTestCases)
	runBenchmarkCases(b, tableTestCases)
	runBenchmarkCases(b, kitchenSinkTestCases)
}
