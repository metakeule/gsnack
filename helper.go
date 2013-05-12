package gsnack

import (
	// ŧ "fmt"
	f "github.com/metakeule/fastreplace"
	// ħ "net/http"
)

type testContainer struct {
	js   []byte
	css  []byte
	html []byte
}

func (ø *testContainer) AddJs(js []byte) {
	ø.js = js
}

func (ø *testContainer) AddCss(css []byte) {
	ø.css = css
}

func (ø *testContainer) AddHtml(html []byte) {
	ø.html = html
}

var testCTemplate = []byte(`<!DOCTYPE html>
	<head>
		<meta charset="utf-8" />
		<style>
  			@@style@@
		</style>
		@@libs@@
		<script>
			@@script@@
		</script>
		<title>@@title@@</title>
	</head>
	<body>
		@@body@@
	</body>
</html>
`)

func (ø *testContainer) Html(libs []*JsLib, template string) string {
	var fr *f.FReplace
	if template == "" {
		fr, _ = f.New(delimiter, testCTemplate)
	} else {
		fr, _ = f.New(delimiter, []byte(template))
	}
	i := fr.Instance()
	i.Assign(`style`, ø.css)
	i.Assign(`body`, ø.html)
	i.Assign(`script`, ø.js)
	i.AssignString(`title`, `Snack Test`)
	l := jslibs(libs)
	i.AssignString(`libs`, l.Include())
	return i.String()
}

func B(s string) []byte { return []byte(s) }
