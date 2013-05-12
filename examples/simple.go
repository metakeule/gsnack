package main

import (
	"fmt"
	. "github.com/metakeule/snack"
	"github.com/metakeule/snack/examples/clicker"
)

var jquery = Jquery(1, 8, 3)
var libs = []*JsLib{jquery}

func tester() {
	vals := map[string][]byte{
		`id`:      []byte(`tester`),
		`color`:   []byte(`red`),
		`content`: []byte(`hiho`),
		`height`:  []byte(`700px`),
	}
	fmt.Println(clicker.New.Test(vals, ""))
}

type MyLayout struct {
	css  string
	js   string
	html string
}

func (ø *MyLayout) AddCss(css []byte)   { ø.css = string(css) }
func (ø *MyLayout) AddJs(js []byte)     { ø.js = string(js) }
func (ø *MyLayout) AddHtml(html []byte) { ø.html = string(html) }
func (ø *MyLayout) Html() string {
	return `<!DOCTYPE html>
	<head>
		<meta charset="utf-8" />
		<title>Self defined template</title>
		<style>` + ø.css + `</style>
		` + jquery.Include() + `
		<script>` + ø.js + `</script>
	</head>
	<body><div class="main">` + ø.html + `</div></body>
</html>`
}

func layouter() {
	l := &MyLayout{}
	c := clicker.New.WithClass("special")
	vals := map[string][]byte{
		`color`:   []byte(`green`),
		`content`: []byte(`hello world`),
		`width`:   []byte(`70px`),
	}
	c.Plug(l, vals)
	fmt.Println(l.Html())
}

func testserve() {
	c := clicker.New.WithClass("y")
	vals := map[string][]byte{
		`color`:   []byte(`yellow`),
		`content`: []byte(`hello in a yellow world`),
	}
	c.TestServe(vals, "")
}

func main() {
	tester()
	layouter()
	testserve()
}
