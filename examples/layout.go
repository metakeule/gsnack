package main

import (
	"github.com/metakeule/snack"
	"github.com/metakeule/snack/examples/clicker"
	"github.com/metakeule/snack/examples/hover"
)

var templ = b(`<!DOCTYPE html>
	<head>
		<meta charset="utf-8" />
		<title>Self defined template</title>
		<style>@@css@@</style>
		@@libs@@
		<script>@@script@@</script>
	</head>
	<body>
		<div class="container">@@g@@</div>
		<div class="container">@@r@@</div>
		<div class="container">@@y@@</div>
	</body>
</html>`)

func b(s string) []byte { return []byte(s) }

var vals = map[string]map[string][]byte{
	"g": map[string][]byte{`color`: b(`green`), `content`: b(`hello green`)},
	"r": map[string][]byte{`color`: b(`red`), `content`: b(`hello red`)},
	"y": map[string][]byte{`color-active`: b(`yellow`), `content`: b(`hello yellow`)},
}

var snacks = map[string]snack.Plugger{
	"g": clicker.New.WithId("green"),
	"r": clicker.New.WithId("red"),
	"y": hover.New.WithId("yellow"),
}

func main() {
	libs := []*snack.JsLib{snack.Jquery(1, 8, 3)}
	snack.Layout(templ).Serve(libs, snacks, vals)
}
